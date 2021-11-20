package basic

import (
	"context"

	"github.com/inigolabs/fezzik"
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

func (c *client) OneAllTypes(ctx context.Context) (*OneAllTypesResponse, fezzik.GQLErrors, error) {

	gqlreq := &fezzik.GQLRequest{
		Operation: "OneAllTypes",
		Query:     OneAllTypesOperation,
		Variables: map[string]interface{}{},
	}

	gqlresp := fezzik.GQLResponse{
		Data:   &OneAllTypesResponse{},
		Errors: fezzik.GQLErrors{},
	}
	err := c.gql.Query(ctx, gqlreq, gqlresp.Data, &gqlresp.Errors)
	if err != nil {
		return nil, nil, err
	}
	return gqlresp.Data.(*OneAllTypesResponse), gqlresp.Errors, nil
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

func (c *client) OneWithSubSelections(ctx context.Context) (*OneWithSubSelectionsResponse, fezzik.GQLErrors, error) {

	gqlreq := &fezzik.GQLRequest{
		Operation: "OneWithSubSelections",
		Query:     OneWithSubSelectionsOperation,
		Variables: map[string]interface{}{},
	}

	gqlresp := fezzik.GQLResponse{
		Data:   &OneWithSubSelectionsResponse{},
		Errors: fezzik.GQLErrors{},
	}
	err := c.gql.Query(ctx, gqlreq, gqlresp.Data, &gqlresp.Errors)
	if err != nil {
		return nil, nil, err
	}
	return gqlresp.Data.(*OneWithSubSelectionsResponse), gqlresp.Errors, nil
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

func (c *client) QueryWithInputs(ctx context.Context, input *QueryWithInputsInputArgs) (
	*QueryWithInputsResponse, fezzik.GQLErrors, error) {

	gqlreq := &fezzik.GQLRequest{
		Operation: "QueryWithInputs",
		Query:     QueryWithInputsOperation,
		Variables: map[string]interface{}{
			"input_one": input.InputOne,
			"input_two": input.InputTwo,
		},
	}

	gqlresp := fezzik.GQLResponse{
		Data:   &QueryWithInputsResponse{},
		Errors: fezzik.GQLErrors{},
	}
	err := c.gql.Query(ctx, gqlreq, gqlresp.Data, &gqlresp.Errors)
	if err != nil {
		return nil, nil, err
	}
	return gqlresp.Data.(*QueryWithInputsResponse), gqlresp.Errors, nil
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

func (c *client) OneAdd(ctx context.Context, input *OneAddInputArgs) (
	*OneAddResponse, fezzik.GQLErrors, error) {

	gqlreq := &fezzik.GQLRequest{
		Operation: "OneAdd",
		Query:     OneAddOperation,
		Variables: map[string]interface{}{
			"input": input.Input,
		},
	}

	gqlresp := fezzik.GQLResponse{
		Data:   &OneAddResponse{},
		Errors: fezzik.GQLErrors{},
	}
	err := c.gql.Query(ctx, gqlreq, gqlresp.Data, &gqlresp.Errors)
	if err != nil {
		return nil, nil, err
	}
	return gqlresp.Data.(*OneAddResponse), gqlresp.Errors, nil
}
