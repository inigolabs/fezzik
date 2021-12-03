package client

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type GQLErrors []*GQLError

type GQLError struct {
	Message    string                 `json:"message"`
	Path       []interface{}          `json:"path,omitempty"`
	Locations  []Location             `json:"locations,omitempty"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`
}

type Location struct {
	Line   int `json:"line,omitempty"`
	Column int `json:"column,omitempty"`
}

func (e *GQLErrors) Error() string {
	var out strings.Builder
	for _, err := range *e {
		out.WriteString(err.Error())
		out.WriteString("\n")
	}
	return out.String()
}

func (e *GQLError) Error() string {
	var out strings.Builder
	if e == nil {
		return ""
	}
	filename, _ := e.Extensions["file"].(string)
	if filename == "" {
		filename = "input"
	}
	out.WriteString(filename)

	if len(e.Locations) > 0 {
		out.WriteByte(':')
		out.WriteString(strconv.Itoa(e.Locations[0].Line))
	}

	out.WriteString(": ")
	if ps := e.pathString(); ps != "" {
		out.WriteString(ps)
		out.WriteByte(' ')
	}

	out.WriteString(e.Message)

	return out.String()
}

func (e GQLError) pathString() string {
	var str bytes.Buffer
	for i, v := range e.Path {

		switch v := v.(type) {
		case int, int64:
			str.WriteString(fmt.Sprintf("[%d]", v))
		default:
			if i != 0 {
				str.WriteByte('.')
			}
			str.WriteString(fmt.Sprint(v))
		}
	}
	return str.String()
}
