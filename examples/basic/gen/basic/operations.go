// Code generated by github.com/inigolabs/fezzik, DO NOT EDIT.

package basic

import (
	"context"
	"encoding/json"

	"github.com/inigolabs/fezzik/client"
	"github.com/inigolabs/fezzik/examples/basic/types"
)

type OneAllTypesOne struct {
	OneInt             *int
	OneIntMust         int
	OneIntList         *[]*int
	OneIntMustList     *[]int
	OneIntMustListMust []int
}

// OneAllTypesResponse response type for OneAllTypes
type OneAllTypesResponse struct {
	One *OneAllTypesOne
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

	resp := &client.GQLResponse{
		Data: &OneAllTypesResponse{},
	}

	err := c.gql.Query(ctx, gqlreq, resp)
	if err != nil {
		return nil, err
	}

	var data *OneAllTypesResponse
	if resp.Data != nil {
		data = resp.Data.(*OneAllTypesResponse)
	}

	if resp.Errors == nil {
		return data, nil
	}

	return data, resp.Errors
}

type OneWithSubSelectionsOne struct {
	Two *OneWithSubSelectionsOneTwo
}

type OneWithSubSelectionsOneTwo struct {
	TwoInt *int
	TwoStr *string
	Three  *OneWithSubSelectionsOneTwoThree
}

type OneWithSubSelectionsOneTwoThree struct {
	ThreeInt *int
}

// OneWithSubSelectionsResponse response type for OneWithSubSelections
type OneWithSubSelectionsResponse struct {
	One *OneWithSubSelectionsOne
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

	resp := &client.GQLResponse{
		Data: &OneWithSubSelectionsResponse{},
	}

	err := c.gql.Query(ctx, gqlreq, resp)
	if err != nil {
		return nil, err
	}

	var data *OneWithSubSelectionsResponse
	if resp.Data != nil {
		data = resp.Data.(*OneWithSubSelectionsResponse)
	}

	if resp.Errors == nil {
		return data, nil
	}

	return data, resp.Errors
}

type OneWithAliasOne struct {
	OneInt *int
}

type OneWithAliasTwo struct {
	OneInt *int
}

// OneWithAliasResponse response type for OneWithAlias
type OneWithAliasResponse struct {
	One *OneWithAliasOne
	Two *OneWithAliasTwo
}

// OneWithAlias from examples/basic/operations/operations.graphql:27
func (c *gqlclient) OneWithAlias(ctx context.Context) (*OneWithAliasResponse, error) {

	var oneWithAliasOperation string = `
	query OneWithAlias {
		one {
			oneInt
		}
		two : one {
			oneInt
		}
	}`

	gqlreq := &client.GQLRequest{
		OperationName: "OneWithAlias",
		Query:         oneWithAliasOperation,
		Variables:     map[string]interface{}{},
	}

	resp := &client.GQLResponse{
		Data: &OneWithAliasResponse{},
	}

	err := c.gql.Query(ctx, gqlreq, resp)
	if err != nil {
		return nil, err
	}

	var data *OneWithAliasResponse
	if resp.Data != nil {
		data = resp.Data.(*OneWithAliasResponse)
	}

	if resp.Errors == nil {
		return data, nil
	}

	return data, resp.Errors
}

type QueryWithInputsOne struct {
	OneInt *int
}

type QueryWithInputsTwo struct {
	TwoInt *int
}

// QueryWithInputsResponse response type for QueryWithInputs
type QueryWithInputsResponse struct {
	One *QueryWithInputsOne
	Two *QueryWithInputsTwo
}

// QueryWithInputs from examples/basic/operations/operations.graphql:38
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

	resp := &client.GQLResponse{
		Data: &QueryWithInputsResponse{},
	}

	err := c.gql.Query(ctx, gqlreq, resp)
	if err != nil {
		return nil, err
	}

	var data *QueryWithInputsResponse
	if resp.Data != nil {
		data = resp.Data.(*QueryWithInputsResponse)
	}

	if resp.Errors == nil {
		return data, nil
	}

	return data, resp.Errors
}

// OneAddResponse response type for OneAdd
type OneAddResponse struct {
	OneAdd *string
}

// OneAdd from examples/basic/operations/operations.graphql:49
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

	resp := &client.GQLResponse{
		Data: &OneAddResponse{},
	}

	err := c.gql.Query(ctx, gqlreq, resp)
	if err != nil {
		return nil, err
	}

	var data *OneAddResponse
	if resp.Data != nil {
		data = resp.Data.(*OneAddResponse)
	}

	if resp.Errors == nil {
		return data, nil
	}

	return data, resp.Errors
}

// TwoAddResponse response type for TwoAdd
type TwoAddResponse struct {
	TwoAdd *string
}

// TwoAdd from examples/basic/operations/operations.graphql:54
func (c *gqlclient) TwoAdd(ctx context.Context,
	input *types.TwoInput,
) (*TwoAddResponse, error) {

	var twoAddOperation string = `
	mutation TwoAdd($input : TwoInput) {
		twoAdd(input: $input)
	}`

	gqlreq := &client.GQLRequest{
		OperationName: "TwoAdd",
		Query:         twoAddOperation,
		Variables: map[string]interface{}{
			"input": input,
		},
	}

	resp := &client.GQLResponse{
		Data: &TwoAddResponse{},
	}

	err := c.gql.Query(ctx, gqlreq, resp)
	if err != nil {
		return nil, err
	}

	var data *TwoAddResponse
	if resp.Data != nil {
		data = resp.Data.(*TwoAddResponse)
	}

	if resp.Errors == nil {
		return data, nil
	}

	return data, resp.Errors
}

// UpdatedResponse response type for Updated
type UpdatedResponse struct {
	Updated *types.Result
}

// Updated from examples/basic/operations/operations.graphql:58
func (c *gqlSubscriptionClient) Updated(fn func(out *UpdatedResponse, err error) error) (string, error) {

	var updatedOperation string = `
	subscription Updated {
		updated {
			id
		}
	}`

	var variables = map[string]interface{}{}

	return c.gql.Exec(updatedOperation, variables, func(in []byte, err error) error {
		if err != nil {
			return fn(nil, err)
		}

		var out *UpdatedResponse
		if err = json.Unmarshal(in, &out); err != nil {
			return err
		}

		return fn(out, nil)
	})
}

// ChangedResponse response type for Changed
type ChangedResponse struct {
	Changed *types.Result
}

// Changed from examples/basic/operations/operations.graphql:64
func (c *gqlSubscriptionClient) Changed(
	input *string,
	fn func(out *ChangedResponse, err error) error,
) (string, error) {

	var changedOperation string = `
	subscription Changed($input : String) {
		changed(input: $input) {
			id
		}
	}`

	var variables = map[string]interface{}{
		"input": input,
	}

	return c.gql.Exec(changedOperation, variables, func(in []byte, err error) error {
		if err != nil {
			return fn(nil, err)
		}

		var out *ChangedResponse
		if err = json.Unmarshal(in, &out); err != nil {
			return err
		}

		return fn(out, nil)
	})
}
