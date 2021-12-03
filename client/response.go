package client

type GQLResponse struct {
	Data   interface{}
	Errors *GQLErrors
}
