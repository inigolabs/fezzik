package client

import (
	"bytes"
	"context"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type GQLClient struct {
	endpoint   string
	httpClient *http.Client
}

func NewGQLClient(endpoint string, httpClient *http.Client) *GQLClient {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &GQLClient{
		endpoint:   endpoint,
		httpClient: httpClient,
	}
}

func (c *GQLClient) SetHttpClient(httpClient *http.Client) {
	c.httpClient = httpClient
}

func (c *GQLClient) Query(ctx context.Context, req *GQLRequest, resp interface{}, errors *GQLErrors) error {
	var body bytes.Buffer
	encoder := json.NewEncoder(&body)
	err := encoder.Encode(req)
	if err != nil {
		return err
	}
	request, err := http.NewRequest(http.MethodPost, c.endpoint, &body)
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	request.Header.Set("Accept", "application/json; charset=utf-8")
	request = request.WithContext(ctx)
	request.Close = true
	response, err := c.httpClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	gqlresponse := &GQLResponse{
		Data:   resp,
		Errors: errors,
	}
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(gqlresponse)
	if err != nil {
		return err
	}
	return nil
}
