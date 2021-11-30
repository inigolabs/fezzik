package render

var ClientTemplate string = `
package {{ .PackageName }}

import (
	"context"
	"net/http"

	"github.com/inigolabs/fezzik/client"
)

type Client interface {
{{- range $operation := .Operations }}
{{- if len $operation.Inputs }}
	{{ $operation.Name }}(ctx context.Context, input *{{ $operation.Name }}InputArgs) (*{{ $operation.Name }}Response, error)
{{- else }}
	{{ $operation.Name }}(ctx context.Context) (*{{ $operation.Name }}Response, error)
{{- end }}	
{{- end }}
}

func NewClient(url string, httpclient *http.Client) Client {
	return &gqlclient{
		gql: client.NewGQLClient(url, httpclient),
	}
}

type gqlclient struct {
	gql *client.GQLClient
}
`
