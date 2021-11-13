package fezzik

import (
	"errors"

	"github.com/jensneuse/graphql-go-tools/pkg/ast"
	"github.com/jensneuse/graphql-go-tools/pkg/astvisitor"
	"github.com/jensneuse/graphql-go-tools/pkg/operationreport"
	"github.com/rs/zerolog/log"
)

type InputRenderInfo struct {
	PackageName string

	InputTypes []InputType
	EnumTypes  map[string]EnumType
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

type EnumType struct {
	Name   string
	Values []string
}

type inputVisitor struct {
	*astvisitor.Walker

	config     *Config
	definition *ast.Document

	info *InputRenderInfo
}

func NewInputVisitor(definition *ast.Document, config *Config) *inputVisitor {
	walker := astvisitor.NewWalker(48)
	visitor := &inputVisitor{
		Walker:     &walker,
		config:     config,
		definition: definition,
		info: &InputRenderInfo{
			PackageName: config.PackageName,
			EnumTypes:   make(map[string]EnumType),
		},
	}
	walker.RegisterEnterInputObjectTypeDefinitionVisitor(visitor)
	walker.RegisterEnterEnumTypeDefinitionVisitor(visitor)
	return visitor
}

func (v *inputVisitor) Walk() error {
	report := &operationreport.Report{}
	v.Walker.Walk(v.definition, v.definition, report)
	if report.HasErrors() {
		log.Error().Str("errors", report.Error()).Msg("")
		return errors.New(report.Error())
	}
	return nil
}

func (v *inputVisitor) EnterInputObjectTypeDefinition(ref int) {
	name := v.definition.InputObjectTypeDefinitionNameString(ref)
	inputType := InputType{
		Name: name,
	}

	for _, fieldRef := range v.definition.InputObjectTypeDefinitions[ref].InputFieldsDefinition.Refs {
		fieldName := v.definition.InputValueDefinitionNameString(fieldRef)
		typeRef := v.definition.InputValueDefinitionType(fieldRef)
		typeInfo := getTypeInfo(typeRef, v.definition)

		inputField := InputField{
			Name: fieldName,
			Type: typeInfo,
		}
		if goType, ok := getGoType(v.config, typeInfo.Name); ok {
			inputField.TypeName = goType
		} else {
			inputField.TypeName = typeInfo.Name
		}
		inputType.Fields = append(inputType.Fields, inputField)
	}

	v.info.InputTypes = append(v.info.InputTypes, inputType)
}

func (v *inputVisitor) EnterEnumTypeDefinition(ref int) {
	enumType := EnumType{
		Name: v.definition.EnumTypeDefinitionNameString(ref),
	}
	refs := v.definition.EnumTypeDefinitions[ref].EnumValuesDefinition.Refs
	for _, valueRef := range refs {
		enumType.Values = append(enumType.Values,
			v.definition.EnumValueDefinitionNameString(valueRef))
	}

	v.info.EnumTypes[enumType.Name] = enumType
}

var inputsTemplate string = `
package {{ .PackageName }}

{{- range $obj := .EnumTypes}}
type {{ $obj.Name }} string
const (
	{{- range $value := $obj.Values }}
			{{ $obj.Name }}_{{ $value }} {{ $obj.Name }} = "{{ $value }}"
	{{- end }}
)

{{ end }}


{{- range $obj := .InputTypes }}
type {{$obj.Name}} struct {
{{- range $field := $obj.Fields }}
	{{- if $field.Type.IsList }}
		{{ pascal $field.Name }} {{if $field.Type.IsListNullable}}*{{end}}[]{{if $field.Type.IsTypeNullable}}*{{end}}{{$field.TypeName}} ~~json:"{{$field.Name}}"~~
	{{- else }}
		{{ pascal $field.Name }} {{if $field.Type.IsTypeNullable}}*{{end}}{{$field.TypeName}} ~~json:"{{$field.Name}}"~~
	{{- end }}
{{- end }}
}

{{ end }}
`
