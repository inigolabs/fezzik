package basic

import (
	"context"

	"github.com/inigolabs/fezzik/client"
)

var OneAllTypesOperation string = `
query OneAllTypes {
   one {
      oneInt
      oneIntMust
      oneIntList
      oneIntMustList
      oneIntMustListMust
   }
}
`

type OneAllTypesResponse struct {
	One *struct {
		OneInt             *int
		OneIntMust         int
		OneIntList         *[]*int
		OneIntMustList     *[]int
		OneIntMustListMust []int
	}
}

func (c *gqlclient) OneAllTypes(ctx context.Context) (*OneAllTypesResponse, error) {

	gqlreq := &client.GQLRequest{
		OperationName: "OneAllTypes",
		Query:         OneAllTypesOperation,
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

var OneWithSubSelectionsOperation string = `
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
}
`

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

func (c *gqlclient) OneWithSubSelections(ctx context.Context) (*OneWithSubSelectionsResponse, error) {

	gqlreq := &client.GQLRequest{
		OperationName: "OneWithSubSelections",
		Query:         OneWithSubSelectionsOperation,
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

var QueryWithInputsOperation string = `
query QueryWithInputs($input_one : String, $input_two : String) {
   one(input: $input_one) {
      oneInt
   }
   two(input: $input_two) {
      twoInt
   }
}
`

type QueryWithInputsInputArgs struct {
	InputOne *string `json:"input_one"`
	InputTwo *string `json:"input_two"`
}

type QueryWithInputsResponse struct {
	One *struct {
		OneInt *int
	}
	Two *struct {
		TwoInt *int
	}
}

func (c *gqlclient) QueryWithInputs(ctx context.Context, input *QueryWithInputsInputArgs) (
	*QueryWithInputsResponse, error) {

	gqlreq := &client.GQLRequest{
		OperationName: "QueryWithInputs",
		Query:         QueryWithInputsOperation,
		Variables: map[string]interface{}{
			"input_one": input.InputOne,
			"input_two": input.InputTwo,
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

var OneAddOperation string = `
mutation OneAdd($input : OneInput) {
   oneAdd(input: $input)
}
`

type OneAddInputArgs struct {
	Input *OneInput `json:"input"`
}

type OneAddResponse struct {
	OneAdd *string
}

func (c *gqlclient) OneAdd(ctx context.Context, input *OneAddInputArgs) (
	*OneAddResponse, error) {

	gqlreq := &client.GQLRequest{
		OperationName: "OneAdd",
		Query:         OneAddOperation,
		Variables: map[string]interface{}{
			"input": input.Input,
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
