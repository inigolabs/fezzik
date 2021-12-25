package basic

import (
	"context"
	"net/http"

	"github.com/inigolabs/fezzik/client"
)

type Client interface {
	OneAllTypes(ctx context.Context) (*OneAllTypesResponse, error)

	OneWithSubSelections(ctx context.Context) (*OneWithSubSelectionsResponse, error)

	QueryWithInputs(ctx context.Context,
		inputOne *string,
		inputTwo *string,
	) (*QueryWithInputsResponse, error)

	OneAdd(ctx context.Context,
		input *OneInput,
	) (*OneAddResponse, error)
}

func NewClient(url string, httpclient *http.Client) Client {
	return &gqlclient{
		gql: client.NewGQLClient(url, httpclient),
	}
}

type gqlclient struct {
	gql *client.GQLClient
}
