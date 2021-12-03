package client

type GQLResponse struct {
	Data   interface{} `json:"data"`
	Errors *GQLErrors  `json:"errors,omitempty"`
}
