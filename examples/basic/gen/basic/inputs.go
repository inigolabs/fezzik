package basic

import "time"

type OneEnum string

const (
	OneEnum_ValA OneEnum = "ValA"
	OneEnum_ValB OneEnum = "ValB"
	OneEnum_ValC OneEnum = "ValC"
)

type __DirectiveLocation string

const (
	__DirectiveLocation_QUERY                  __DirectiveLocation = "QUERY"
	__DirectiveLocation_MUTATION               __DirectiveLocation = "MUTATION"
	__DirectiveLocation_SUBSCRIPTION           __DirectiveLocation = "SUBSCRIPTION"
	__DirectiveLocation_FIELD                  __DirectiveLocation = "FIELD"
	__DirectiveLocation_FRAGMENT_DEFINITION    __DirectiveLocation = "FRAGMENT_DEFINITION"
	__DirectiveLocation_FRAGMENT_SPREAD        __DirectiveLocation = "FRAGMENT_SPREAD"
	__DirectiveLocation_INLINE_FRAGMENT        __DirectiveLocation = "INLINE_FRAGMENT"
	__DirectiveLocation_VARIABLE_DEFINITION    __DirectiveLocation = "VARIABLE_DEFINITION"
	__DirectiveLocation_SCHEMA                 __DirectiveLocation = "SCHEMA"
	__DirectiveLocation_SCALAR                 __DirectiveLocation = "SCALAR"
	__DirectiveLocation_OBJECT                 __DirectiveLocation = "OBJECT"
	__DirectiveLocation_FIELD_DEFINITION       __DirectiveLocation = "FIELD_DEFINITION"
	__DirectiveLocation_ARGUMENT_DEFINITION    __DirectiveLocation = "ARGUMENT_DEFINITION"
	__DirectiveLocation_INTERFACE              __DirectiveLocation = "INTERFACE"
	__DirectiveLocation_UNION                  __DirectiveLocation = "UNION"
	__DirectiveLocation_ENUM                   __DirectiveLocation = "ENUM"
	__DirectiveLocation_ENUM_VALUE             __DirectiveLocation = "ENUM_VALUE"
	__DirectiveLocation_INPUT_OBJECT           __DirectiveLocation = "INPUT_OBJECT"
	__DirectiveLocation_INPUT_FIELD_DEFINITION __DirectiveLocation = "INPUT_FIELD_DEFINITION"
)

type __TypeKind string

const (
	__TypeKind_SCALAR       __TypeKind = "SCALAR"
	__TypeKind_OBJECT       __TypeKind = "OBJECT"
	__TypeKind_INTERFACE    __TypeKind = "INTERFACE"
	__TypeKind_UNION        __TypeKind = "UNION"
	__TypeKind_ENUM         __TypeKind = "ENUM"
	__TypeKind_INPUT_OBJECT __TypeKind = "INPUT_OBJECT"
	__TypeKind_LIST         __TypeKind = "LIST"
	__TypeKind_NON_NULL     __TypeKind = "NON_NULL"
)

type OneInput struct {
	OneInt             *int       `json:"oneInt"`
	OneIntMust         int        `json:"oneIntMust"`
	OneIntList         *[]*int    `json:"oneIntList"`
	OneIntMustList     *[]int     `json:"oneIntMustList"`
	OneIntMustListMust []int      `json:"oneIntMustListMust"`
	OneStr             *string    `json:"oneStr"`
	OneBool            *bool      `json:"oneBool"`
	OneEnum            *OneEnum   `json:"oneEnum"`
	OneTime            *time.Time `json:"oneTime"`
}
