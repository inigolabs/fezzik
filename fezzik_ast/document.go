package fezzik_ast

import "github.com/vektah/gqlparser/v2/ast"

func NewDocument() *Document {
	return &Document{
		InputTypes: make(map[InputTypeName]InputType),
		EnumTypes:  make(map[EnumTypeName]EnumType),
	}
}

type Document struct {
	PackageName string

	InputTypes map[InputTypeName]InputType
	EnumTypes  map[EnumTypeName]EnumType

	Operations []*OperationInfo
}

type InputTypeName = string

type InputType struct {
	Name   InputTypeName
	Fields []InputField
}

type InputField struct {
	Name      string
	Type      *TypeInfo
	TypeName  string
	StructTag string
}

type TypeInfo struct {
	Name           string
	IsList         bool
	IsListNullable bool
	IsTypeNullable bool
}

type EnumTypeName = string

type EnumType struct {
	Name   EnumTypeName
	Values []string
}

type OperationInfo struct {
	Name             string
	OperationType    ast.Operation
	Operation        string
	ResponseType     string
	ResponseSubTypes []string
	Inputs           []OperationInput
	Source           Source
}

type OperationInput struct {
	Name string
	Type string
}

type Source struct {
	FileName string
	Line     int
}
