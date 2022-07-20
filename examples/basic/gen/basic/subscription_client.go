// Code generated by github.com/inigolabs/fezzik, DO NOT EDIT.

package basic

import (
	subscription "github.com/hasura/go-graphql-client"
)

type SubscriptionClient interface {
	// Updated from examples/basic/operations/operations.graphql:53
	Updated(fn func(out *UpdatedResponse, err error) error) (string, error)

	// Changed from examples/basic/operations/operations.graphql:59
	Changed(
		input *string,
		fn func(out *ChangedResponse, err error) error) (string, error)

	Close() (err error)
}

func NewSubscriptionClient(url string) SubscriptionClient {
	gql := subscription.NewSubscriptionClient(url)

	go gql.Run()

	return &gqlSubscriptionClient{gql: gql}
}

type gqlSubscriptionClient struct {
	gql *subscription.SubscriptionClient
}

func (c *gqlSubscriptionClient) Close() (err error) {
	return c.gql.Close()
}
