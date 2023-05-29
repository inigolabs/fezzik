package subscription

import "github.com/hasura/go-graphql-client"

type SubscriptionClient = graphql.SubscriptionClient

func NewSubscriptionClient(url string) *SubscriptionClient {
	return graphql.NewSubscriptionClient(url)
}
