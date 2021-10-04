package basic

import "github.com/machinebox/graphql"

func NewClient(url string) *Client {
	return &Client{
		gql: graphql.NewClient(url),
	}
}

type Client struct {
	gql *graphql.Client
}
