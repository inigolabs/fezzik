package basic

import (
	"context"
	"net/http"

	"github.com/machinebox/graphql"
)

type ClientInterface interface {
	OneAllTypes(ctx context.Context) (*OneAllTypesResponse, error)
	OneWithSubSelections(ctx context.Context) (*OneWithSubSelectionsResponse, error)
	QueryWithInputs(ctx context.Context, input *QueryWithInputsInputArgs) (*QueryWithInputsResponse, error)
	OneAdd(ctx context.Context, input *OneAddInputArgs) (*OneAddResponse, error)
}

func NewClient(url string, httpclient *http.Client) ClientInterface {
	if httpclient != nil {
		return &client{
			gql: graphql.NewClient(url, graphql.WithHTTPClient(httpclient)),
		}
	} else {
		return &client{
			gql: graphql.NewClient(url),
		}
	}
}

type client struct {
	gql *graphql.Client
}
