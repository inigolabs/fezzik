package basic

import (
	"context"

	"github.com/machinebox/graphql"
	"github.com/mitchellh/mapstructure"
	"github.com/rs/zerolog/log"
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
		OneIntList         []*int
		OneIntMustList     []int
		OneIntMustListMust []int
	}
}

func (c *client) OneAllTypes(ctx context.Context) (*OneAllTypesResponse, error) {

	q := graphql.NewRequest(OneAllTypesOperation)
	var resp map[string]interface{}
	err := c.gql.Run(ctx, q, &resp)
	log.Debug().Interface("resp", resp).Err(err).Msg("OneAllTypes")
	if err != nil {
		return nil, err
	}

	output := OneAllTypesResponse{}
	err = mapstructure.Decode(resp, &output)
	log.Debug().Interface("output", output).Err(err).Msg("OneAllTypes")
	if err != nil {
		return nil, err
	}
	return &output, err
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

func (c *client) OneWithSubSelections(ctx context.Context) (*OneWithSubSelectionsResponse, error) {

	q := graphql.NewRequest(OneWithSubSelectionsOperation)
	var resp map[string]interface{}
	err := c.gql.Run(ctx, q, &resp)
	log.Debug().Interface("resp", resp).Err(err).Msg("OneWithSubSelections")
	if err != nil {
		return nil, err
	}

	output := OneWithSubSelectionsResponse{}
	err = mapstructure.Decode(resp, &output)
	log.Debug().Interface("output", output).Err(err).Msg("OneWithSubSelections")
	if err != nil {
		return nil, err
	}
	return &output, err
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
	*QueryWithInputsResponse, error) {

	q := graphql.NewRequest(QueryWithInputsOperation)
	q.Var("input_one", input.InputOne)
	q.Var("input_two", input.InputTwo)
	var resp map[string]interface{}
	err := c.gql.Run(ctx, q, &resp)
	log.Debug().Interface("resp", resp).Err(err).Msg("QueryWithInputs")
	if err != nil {
		return nil, err
	}

	output := QueryWithInputsResponse{}
	err = mapstructure.Decode(resp, &output)
	log.Debug().Interface("output", output).Err(err).Msg("QueryWithInputs")
	if err != nil {
		return nil, err
	}
	return &output, err
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
	*OneAddResponse, error) {

	q := graphql.NewRequest(OneAddOperation)
	q.Var("input", input.Input)
	var resp map[string]interface{}
	err := c.gql.Run(ctx, q, &resp)
	log.Debug().Interface("resp", resp).Err(err).Msg("OneAdd")
	if err != nil {
		return nil, err
	}

	output := OneAddResponse{}
	err = mapstructure.Decode(resp, &output)
	log.Debug().Interface("output", output).Err(err).Msg("OneAdd")
	if err != nil {
		return nil, err
	}
	return &output, err
}
