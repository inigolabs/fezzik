package client

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGQLError(t *testing.T) {
	type testcase struct {
		name   string
		err    *GQLError
		expect string
	}

	tests := []testcase{
		{
			name:   "Empty",
			err:    nil,
			expect: "",
		},
		{
			name: "Message",
			err: &GQLError{
				Message: "failed",
			},
			expect: "failed",
		},
		{
			name: "MessageWithOperationName",
			err: &GQLError{
				Message: "failed",
				Extensions: map[string]interface{}{
					"operationName": "op",
				},
			},
			expect: "op: failed",
		},
		{
			name: "MessageWithOperationNameAndPath",
			err: &GQLError{
				Message: "failed",
				Path:    []interface{}{"aaa", "bbb"},
				Extensions: map[string]interface{}{
					"operationName": "op",
				},
			},
			expect: "op: aaa.bbb: failed",
		},
		{
			name: "MessageWithOperationNameTypeAndPath",
			err: &GQLError{
				Message: "failed",
				Path:    []interface{}{"aaa", "bbb"},
				Extensions: map[string]interface{}{
					"operationName": "op",
					"operationType": "query",
				},
			},
			expect: "op: query.aaa.bbb: failed",
		},
		{
			name: "MessageWithOperationNameType",
			err: &GQLError{
				Message: "failed",
				Extensions: map[string]interface{}{
					"operationName": "op",
					"operationType": "query",
				},
			},
			expect: "op: failed",
		},
		{
			name: "MessageWithPath",
			err: &GQLError{
				Message: "failed",
				Path:    []interface{}{"aaa", "bbb"},
			},
			expect: "aaa.bbb: failed",
		},
		{
			name: "MessageWithTypeAndPath",
			err: &GQLError{
				Message: "failed",
				Path:    []interface{}{"aaa", "bbb"},
				Extensions: map[string]interface{}{
					"operationType": "query",
				},
			},
			expect: "query.aaa.bbb: failed",
		},
		{
			name: "MessageWithPathMultiType",
			err: &GQLError{
				Message: "failed",
				Path:    []interface{}{"aaa", 1, "bbb"},
			},
			expect: "aaa[1].bbb: failed",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := test.err.Error()
			require.Equal(t, test.expect, actual)
		})
	}
}
