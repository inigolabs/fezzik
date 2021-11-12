package basic

import (
	"context"
	"net/http"

	"github.com/machinebox/graphql"
	"github.com/rs/zerolog/log"
)

type Client interface {
	OneAllTypes(ctx context.Context) (*OneAllTypesResponse, error)
	OneWithSubSelections(ctx context.Context) (*OneWithSubSelectionsResponse, error)
	QueryWithInputs(ctx context.Context, input *QueryWithInputsInputArgs) (*QueryWithInputsResponse, error)
	OneAdd(ctx context.Context, input *OneAddInputArgs) (*OneAddResponse, error)
}

func NewClient(url string, httpclient *http.Client) Client {
	var gqlclient *graphql.Client

	if httpclient != nil {
		gqlclient = graphql.NewClient(url, graphql.WithHTTPClient(httpclient))
	} else {
		gqlclient = graphql.NewClient(url)
	}

	gqlclient.Log = func(s string) { log.Debug().Msg(s) }

	return &client{
		gql: gqlclient,
	}
}

type client struct {
	gql *graphql.Client
}
