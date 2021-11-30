package fezzik_ast

func NewDocument() *Document {
	return &Document{
		EnumTypes: make(map[string]EnumType),
	}
}

type Document struct {
	PackageName string

	InputTypes []InputType
	EnumTypes  map[string]EnumType

	Operations []*OperationInfo
}

type InputType struct {
	Name   string
	Fields []InputField
}

type InputField struct {
	Name     string
	Type     *TypeInfo
	TypeName string
}

type TypeInfo struct {
	Name           string
	IsList         bool
	IsListNullable bool
	IsTypeNullable bool
}

type EnumType struct {
	Name   string
	Values []string
}

type OperationInfo struct {
	Name         string
	Operation    string
	ResponseType string
	Inputs       []OperationInput
}

type OperationInput struct {
	Name string
	Type string
}
