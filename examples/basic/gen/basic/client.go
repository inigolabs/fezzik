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

	// UnionSearchResultInLine from examples/basic/operations/operations.graphql:27
	UnionSearchResultInLine(ctx context.Context) (*UnionSearchResultInLineResponse, error)

	// UnionSearchResult from examples/basic/operations/operations.graphql:39
	UnionSearchResult(ctx context.Context) (*UnionSearchResultResponse, error)

	// UnionSearchResultNoTypename from examples/basic/operations/operations.graphql:49
	UnionSearchResultNoTypename(ctx context.Context) (*UnionSearchResultNoTypenameResponse, error)

	// InterfaceGetCreated from examples/basic/operations/operations.graphql:57
	InterfaceGetCreated(ctx context.Context) (*InterfaceGetCreatedResponse, error)

	// InterfaceGetCreatedNoTypename from examples/basic/operations/operations.graphql:67
	InterfaceGetCreatedNoTypename(ctx context.Context) (*InterfaceGetCreatedNoTypenameResponse, error)

	// OneWithAlias from examples/basic/operations/operations.graphql:91
	OneWithAlias(ctx context.Context) (*OneWithAliasResponse, error)

	// QueryWithInputs from examples/basic/operations/operations.graphql:102
	QueryWithInputs(ctx context.Context,
		inputOne *string,
		inputTwo *string,
	) (*QueryWithInputsResponse, error)

	// OneAdd from examples/basic/operations/operations.graphql:113
	OneAdd(ctx context.Context,
		input *OneInput,
	) (*OneAddResponse, error)

	// TwoAdd from examples/basic/operations/operations.graphql:118
	TwoAdd(ctx context.Context,
		input *types.TwoInput,
	) (*TwoAddResponse, error)
}

func NewClient(url string, httpclient *http.Client) Client {
	return &gqlclient{
		gql: client.NewGQLClient(url, httpclient),
	}
}

func NewDebugClient(url string, httpclient *http.Client) Client {
	return &gqlclient{
		gql: client.NewGQLClient(url, httpclient, client.WithDebug()),
	}
}

type gqlclient struct {
	gql *client.GQLClient
}
