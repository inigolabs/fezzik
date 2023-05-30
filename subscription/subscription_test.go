package subscription

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/require"
	"net"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

const ack = `{"type":"connection_ack"}`

type message struct {
	ID      string      `json:"id"`
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

var upgrader = websocket.Upgrader{}

func handle(responses ...map[string]interface{}) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}

		defer c.Close()

		for {
			mt, msg, err := c.ReadMessage()
			if err != nil {
				break
			}

			var incoming message
			if err = json.Unmarshal(msg, &incoming); err != nil {
				break
			}

			switch incoming.Type {
			case "connection_init":
				if err = c.WriteMessage(mt, []byte(ack)); err != nil {
					break
				}
			case "start":
				for i := range responses {
					resp := message{
						ID:      incoming.ID,
						Type:    "data",
						Payload: map[string]interface{}{"data": responses[i]},
					}

					bytes, err := json.Marshal(resp)
					if err != nil {
						break
					}

					if err = c.WriteMessage(mt, bytes); err != nil {
						break
					}
				}
			}
		}
	}
}

func startServer(t *testing.T, l net.Listener, fn http.HandlerFunc) func() {
	t.Helper()

	svr := &httptest.Server{
		Listener: l,
		Config:   &http.Server{Handler: fn},
	}

	svr.Start()

	return svr.Close
}

func TestNewSubscriptionClient_Reconnect(t *testing.T) {
	const reconnections int32 = 3

	l, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)

	responses := []map[string]interface{}{
		{"updated": 1},
		{"updated": 2},
		{"updated": 3},
		{"updated": 4},
	}

	stop := startServer(t, l, handle(responses...))

	client := NewSubscriptionClient("http://" + l.Addr().String())
	client.WithTimeout(time.Second)

	var connectedCounter int32
	client.OnConnected(func() {
		atomic.AddInt32(&connectedCounter, 1)
	})

	var disconnectedCounter int32
	client.OnDisconnected(func() {
		atomic.AddInt32(&disconnectedCounter, 1)
	})

	go client.Run()

	var done = make(chan struct{})

	var counter int32

	_, err = client.Exec("subscription Updated { updated { id }}", nil, func(message []byte, err error) error {
		var msg map[string]int
		require.NoError(t, json.Unmarshal(message, &msg))

		counterVal := atomic.AddInt32(&counter, 1)
		if counterVal == int32(len(responses))*reconnections {
			require.NoError(t, client.Close())

			close(done)

			return nil
		}

		if counterVal%int32(len(responses)) == 0 {
			stop()

			l, err = net.Listen("tcp", l.Addr().String())
			require.NoError(t, err)
			stop = startServer(t, l, handle(responses...))
		}

		return nil
	})
	require.NoError(t, err)

	select {
	case <-time.After(time.Second * 10):
		require.Fail(t, "timeout")
	case <-done:
	}

	require.Equal(t, reconnections*int32(len(responses)), counter)
	require.Equal(t, reconnections, connectedCounter)
	require.Equal(t, reconnections, disconnectedCounter)
}
