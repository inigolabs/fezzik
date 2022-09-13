// MIT License
//
// Copyright (c) 2017 Dmitri Shuralyov
// Copyright (c) 2020 Hasura
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package subscription

// OptionType represents the logic of graphql query construction
type OptionType string

const (
	// optionTypeOperationName is private because it's option is built-in and unique
	optionTypeOperationName      OptionType = "operation_name"
	OptionTypeOperationDirective OptionType = "operation_directive"
)

// Option abstracts an extra render interface for the query string
// They are optional parts. By default GraphQL queries can request data without them
type Option interface {
	// Type returns the supported type of the renderer
	// available types: operation_name and operation_directive
	Type() OptionType
	// String returns the query component string
	String() string
}

// operationNameOption represents the operation name render component
type operationNameOption struct {
	name string
}

func (ono operationNameOption) Type() OptionType {
	return optionTypeOperationName
}

func (ono operationNameOption) String() string {
	return ono.name
}

// OperationName creates the operation name option
func OperationName(name string) Option {
	return operationNameOption{name}
}
