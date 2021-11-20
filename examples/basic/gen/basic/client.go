package basic

import (
	"context"
	"net/http"

	"github.com/inigolabs/fezzik"
)

type Client interface {
	OneAllTypes(ctx context.Context) (*OneAllTypesResponse, error)
	OneWithSubSelections(ctx context.Context) (*OneWithSubSelectionsResponse, error)
	QueryWithInputs(ctx context.Context, input *QueryWithInputsInputArgs) (*QueryWithInputsResponse, error)
	OneAdd(ctx context.Context, input *OneAddInputArgs) (*OneAddResponse, error)
}

func NewClient(url string, httpclient *http.Client) Client {
	gqlclient := fezzik.NewGQLClient(url, httpclient)
	return &client{
		gql: gqlclient,
	}
}

type client struct {
	gql *fezzik.GQLClient
}
