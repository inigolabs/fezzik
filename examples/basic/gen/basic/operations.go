package basic

import (
	"context"

	"github.com/inigolabs/fezzik/client"
)

// OneAllTypesResponse response type for OneAllTypes
type OneAllTypesResponse struct {
	One *struct {
		OneInt             *int
		OneIntMust         int
		OneIntList         *[]*int
		OneIntMustList     *[]int
		OneIntMustListMust []int
	}
}

// OneAllTypes from examples/basic/operations/operations.graphql:2
func (c *gqlclient) OneAllTypes(ctx context.Context) (*OneAllTypesResponse, error) {

	var oneAllTypesOperation string = `
	query OneAllTypes {
		one {
			oneInt
			oneIntMust
			oneIntList
			oneIntMustList
			oneIntMustListMust
		}
	}`

	gqlreq := &client.GQLRequest{
		OperationName: "OneAllTypes",
		Query:         oneAllTypesOperation,
		Variables:     map[string]interface{}{},
	}

	var gqldata OneAllTypesResponse
	var gqlerrs client.GQLErrors
	err := c.gql.Query(ctx, gqlreq, &gqldata, &gqlerrs)
	if err != nil {
		return nil, err
	}
	if len(gqlerrs) == 0 {
		return &gqldata, nil
	}
	return &gqldata, &gqlerrs
}

// OneWithSubSelectionsResponse response type for OneWithSubSelections
type OneWithSubSelectionsResponse struct {
	One *struct {
		Two *struct {
			TwoInt *int
			TwoStr *string
			Three  *struct {
				ThreeInt *int
			}
		}
	}
}

// OneWithSubSelections from examples/basic/operations/operations.graphql:13
func (c *gqlclient) OneWithSubSelections(ctx context.Context) (*OneWithSubSelectionsResponse, error) {

	var oneWithSubSelectionsOperation string = `
	query OneWithSubSelections {
		one {
			two {
				twoInt
				twoStr
				three {
					threeInt
				}
			}
		}
	}`

	gqlreq := &client.GQLRequest{
		OperationName: "OneWithSubSelections",
		Query:         oneWithSubSelectionsOperation,
		Variables:     map[string]interface{}{},
	}

	var gqldata OneWithSubSelectionsResponse
	var gqlerrs client.GQLErrors
	err := c.gql.Query(ctx, gqlreq, &gqldata, &gqlerrs)
	if err != nil {
		return nil, err
	}
	if len(gqlerrs) == 0 {
		return &gqldata, nil
	}
	return &gqldata, &gqlerrs
}

// QueryWithInputsResponse response type for QueryWithInputs
type QueryWithInputsResponse struct {
	One *struct {
		OneInt *int
	}
	Two *struct {
		TwoInt *int
	}
}

// QueryWithInputs from examples/basic/operations/operations.graphql:27
func (c *gqlclient) QueryWithInputs(ctx context.Context,
	inputOne *string,
	inputTwo *string,
) (*QueryWithInputsResponse, error) {

	var queryWithInputsOperation string = `
	query QueryWithInputs($input_one : String, $input_two : String) {
		one(input: $input_one) {
			oneInt
		}
		two(input: $input_two) {
			twoInt
		}
	}`

	gqlreq := &client.GQLRequest{
		OperationName: "QueryWithInputs",
		Query:         queryWithInputsOperation,
		Variables: map[string]interface{}{
			"input_one": inputOne,
			"input_two": inputTwo,
		},
	}

	var gqldata QueryWithInputsResponse
	var gqlerrs client.GQLErrors
	err := c.gql.Query(ctx, gqlreq, &gqldata, &gqlerrs)
	if err != nil {
		return nil, err
	}
	if len(gqlerrs) == 0 {
		return &gqldata, nil
	}
	return &gqldata, &gqlerrs
}

// OneAddResponse response type for OneAdd
type OneAddResponse struct {
	OneAdd *string
}

// OneAdd from examples/basic/operations/operations.graphql:38
func (c *gqlclient) OneAdd(ctx context.Context,
	input *OneInput,
) (*OneAddResponse, error) {

	var oneAddOperation string = `
	mutation OneAdd($input : OneInput) {
		oneAdd(input: $input)
	}`

	gqlreq := &client.GQLRequest{
		OperationName: "OneAdd",
		Query:         oneAddOperation,
		Variables: map[string]interface{}{
			"input": input,
		},
	}

	var gqldata OneAddResponse
	var gqlerrs client.GQLErrors
	err := c.gql.Query(ctx, gqlreq, &gqldata, &gqlerrs)
	if err != nil {
		return nil, err
	}
	if len(gqlerrs) == 0 {
		return &gqldata, nil
	}
	return &gqldata, &gqlerrs
}
