// Code generated by github.com/inigolabs/fezzik, DO NOT EDIT.

package basic

import (
	"context"
	"net/http"

	"github.com/inigolabs/fezzik/client"
	"github.com/inigolabs/fezzik/examples/basic/types"
)

type Client interface {
	// OneAllTypes from examples/basic/operations/operations.graphql:2
	OneAllTypes(ctx context.Context) (*OneAllTypesResponse, error)

	// OneWithSubSelections from examples/basic/operations/operations.graphql:13
	OneWithSubSelections(ctx context.Context) (*OneWithSubSelectionsResponse, error)

	// OneWithAlias from examples/basic/operations/operations.graphql:27
	OneWithAlias(ctx context.Context) (*OneWithAliasResponse, error)

	// QueryWithInputs from examples/basic/operations/operations.graphql:38
	QueryWithInputs(ctx context.Context,
		inputOne *string,
		inputTwo *string,
	) (*QueryWithInputsResponse, error)

	// OneAdd from examples/basic/operations/operations.graphql:49
	OneAdd(ctx context.Context,
		input *OneInput,
	) (*OneAddResponse, error)

	// TwoAdd from examples/basic/operations/operations.graphql:54
	TwoAdd(ctx context.Context,
		input *types.TwoInput,
	) (*TwoAddResponse, error)
}

func NewClient(url string, httpclient *http.Client) Client {
	return &gqlclient{
		gql: client.NewGQLClient(url, httpclient),
	}
}

type gqlclient struct {
	gql *client.GQLClient
}
