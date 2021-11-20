package fezzik

import (
	"errors"
	"strings"

	"github.com/inigolabs/fezzik/common"
	"github.com/jensneuse/graphql-go-tools/pkg/ast"
	"github.com/jensneuse/graphql-go-tools/pkg/astvisitor"
	"github.com/jensneuse/graphql-go-tools/pkg/operationreport"

	"github.com/rs/zerolog/log"
)

type Info struct {
	Debug       bool
	PackageName string

	Operations []*OperationInfo
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

type currOperation struct {
	name      string
	operation *builder
	response  *builder
	inputs    []OperationInput
}

type visitor struct {
	*astvisitor.Walker

	cfg        *Config
	operation  *ast.Document
	definition *ast.Document

	inputsInfo *InputRenderInfo

	info *Info

	curr *currOperation
}

func NewVisitor(cfg *Config, schema *ast.Document, inputsInfo *InputRenderInfo) *visitor {
	walker := astvisitor.NewWalker(48)
	visitor := &visitor{
		Walker:     &walker,
		cfg:        cfg,
		definition: schema,
		inputsInfo: inputsInfo,
		info: &Info{
			Debug:       cfg.Debug,
			PackageName: cfg.PackageName,
		},
	}
	walker.RegisterEnterOperationVisitor(visitor)
	walker.RegisterLeaveOperationVisitor(visitor)
	walker.RegisterSelectionSetVisitor(visitor)
	walker.RegisterFieldVisitor(visitor)
	walker.RegisterArgumentVisitor(visitor)
	walker.RegisterVariableDefinitionVisitor(visitor)
	return visitor
}

func (v *visitor) AddInfo(queryStr string, query *ast.Document) error {
	v.operation = query
	report := operationreport.Report{}
	v.Walker.Walk(query, v.definition, &report)
	if report.HasErrors() {
		log.Error().Str("errors", report.Error()).Msg("")
		return errors.New(report.Error())
	}

	return nil
}

func (v *visitor) EnterOperationDefinition(ref int) {
	v.curr = &currOperation{
		operation: NewBuilder(),
		response:  NewBuilder(),
	}

	hasName := v.operation.OperationDefinitions[ref].Name.Length() > 0
	if !hasName {
		v.StopWithExternalErr(operationreport.ExternalError{
			Message: "operation must be named",
		})
	}

	v.curr.name = string(v.operation.Input.ByteSlice(v.operation.OperationDefinitions[ref].Name))

	switch v.operation.OperationDefinitions[ref].OperationType {
	case ast.OperationTypeQuery:
		v.curr.operation.Write(v.indent(), "query ", v.curr.name)
	case ast.OperationTypeMutation:
		v.curr.operation.Write(v.indent(), "mutation ", v.curr.name)
	case ast.OperationTypeSubscription:
		v.StopWithExternalErr(operationreport.ExternalError{
			Message: "subscriptions not supported",
		})
	}

	log.Debug().Str("name", v.curr.name).Msg("EnterOperationDefinition")
}

func (v *visitor) LeaveOperationDefinition(ref int) {
	v.info.Operations = append(v.info.Operations, &OperationInfo{
		Name:         v.curr.name,
		Operation:    v.curr.operation.String(),
		ResponseType: v.curr.response.String(),
		Inputs:       v.curr.inputs,
	})
}

func (v *visitor) EnterSelectionSet(ref int) {
	v.curr.operation.Writeln(" {")

	parentTypeName := v.definition.NodeNameUnsafeString(v.EnclosingTypeDefinition)
	if parentTypeName != "Query" && parentTypeName != "Mutation" {
		v.curr.response.Writeln(v.indent(), " {")
	}
}

func (v *visitor) LeaveSelectionSet(ref int) {
	v.curr.operation.Write(v.indent(), "}")

	parentTypeName := v.definition.NodeNameUnsafeString(v.EnclosingTypeDefinition)
	if parentTypeName != "Query" && parentTypeName != "Mutation" {
		v.curr.response.Writeln(" }")
	}
}

func (v *visitor) EnterField(ref int) {
	// render operation
	field := v.operation.Fields[ref]
	name := v.operation.Input.ByteSliceString(field.Name)
	v.curr.operation.Write(v.indent(), name)

	// render response
	fieldDefRef, exist := v.FieldDefinition(ref)
	if !exist {
		panic("field not found") // TODO - better error message, don't panic
	}
	fieldDefTypeRef := v.definition.FieldDefinitionType(fieldDefRef)
	typeInfo := getTypeInfo(fieldDefTypeRef, v.definition)

	v.curr.response.Write(common.UppercaseFirstChar(name), " ")
	if typeInfo.IsList {
		if typeInfo.IsListNullable {
			v.curr.response.Write("*")
		}
		v.curr.response.Write("[]")
	}
	if typeInfo.IsTypeNullable {
		v.curr.response.Write("*")
	}
	if goType, ok := getGoType(v.cfg, typeInfo.Name); ok {
		v.curr.response.Writeln(goType)
	} else if _, found := v.inputsInfo.EnumTypes[typeInfo.Name]; found {
		v.curr.response.Writeln(typeInfo.Name)
	} else {
		v.curr.response.Write("struct")
	}

	log.Debug().Str("name", string(name)).Str("type", typeInfo.Name).Msg("EnterField")
}

func (v *visitor) LeaveField(ref int) {
	v.curr.operation.Writeln()
}

func (v *visitor) EnterVariableDefinition(ref int) {
	// render operation
	if !v.operation.VariableDefinitionsBefore(ref) {
		v.curr.operation.Write(v.indent(), "(")
	}

	varDef := v.operation.VariableDefinitions[ref]
	varName, err := v.operation.PrintValueBytes(varDef.VariableValue, nil)
	varNameStr := string(varName[1:])
	check(err)
	varType, err := v.operation.PrintTypeBytes(varDef.Type, nil)
	check(err)
	v.curr.operation.Write("$", varNameStr, " : ", varType)

	if v.operation.VariableDefinitions[ref].DefaultValue.IsDefined {
		defaultStr, err := v.operation.PrintValueBytes(varDef.DefaultValue.Value, nil)
		check(err)
		v.curr.operation.Write(" = ", defaultStr)
	}

	if v.operation.VariableDefinitions[ref].HasDirectives {
		v.curr.operation.Write(" ")
	}

	// render inputs
	typeInfo := getTypeInfo(varDef.Type, v.operation)
	typeBuilder := new(strings.Builder)
	if typeInfo.IsList {
		if typeInfo.IsListNullable {
			typeBuilder.WriteString("*")
		}
		typeBuilder.WriteString("[]")
	}
	if typeInfo.IsTypeNullable {
		typeBuilder.WriteString("*")
	}
	if goType, ok := getGoType(v.cfg, typeInfo.Name); ok {
		typeBuilder.WriteString(goType)
	} else {
		typeBuilder.WriteString(typeInfo.Name)
	}

	v.curr.inputs = append(v.curr.inputs, OperationInput{
		Name: varNameStr,
		Type: typeBuilder.String(),
	})

	log.Debug().Str("name", varNameStr).Bytes("type", varType).Msg("EnterVariableDefinition")
}

func (v *visitor) LeaveVariableDefinition(ref int) {
	if !v.operation.VariableDefinitionsAfter(ref) {
		v.curr.operation.Write(")")
	} else {
		v.curr.operation.Write(", ")
	}

	varDefRef := v.operation.VariableDefinitions[ref]
	varName, err := v.operation.PrintValueBytes(varDefRef.VariableValue, nil)
	check(err)
	varType, err := v.operation.PrintTypeBytes(varDefRef.Type, nil)
	check(err)
	log.Debug().Bytes("name", varName).Bytes("type", varType).Msg("LeaveVariableDefinition")
}

func (v *visitor) EnterArgument(ref int) {
	lastAncestor := v.Ancestors[len(v.Ancestors)-1]
	if len(v.operation.ArgumentsBefore(lastAncestor, ref)) == 0 {
		v.curr.operation.Write("(")
	} else {
		v.curr.operation.Write(", ")
	}

	argName := v.operation.Input.ByteSlice(v.operation.Arguments[ref].Name)
	argVal, err := v.operation.PrintValueBytes(v.operation.Arguments[ref].Value, nil)
	check(err)
	v.curr.operation.Write(argName, ": ", argVal)
}

func (v *visitor) LeaveArgument(ref int) {
	lastAncestor := v.Ancestors[len(v.Ancestors)-1]
	if len(v.operation.ArgumentsAfter(lastAncestor, ref)) == 0 {
		v.curr.operation.Write(")")
	}
}

func (v *visitor) indent() []byte {

	if len(v.Ancestors) == 0 {
		return []byte("")
	}

	switch v.Ancestors[0].Kind {
	case ast.NodeKindOperationDefinition,
		ast.NodeKindFragmentDefinition:
	default:
		return []byte("   ")
	}

	buf := []byte{}
	for i := range v.Ancestors {
		if v.Ancestors[i].Kind == ast.NodeKindSelectionSet {
			buf = append(buf, []byte("   ")...)
		}
	}
	return buf
}

var operationsTemplate string = `
package {{ .PackageName }}

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/machinebox/graphql"
	"github.com/mitchellh/mapstructure"
)

{{ $root := . }}

{{- range $operation := .Operations }}

var {{ $operation.Name }}Operation string = ~~
{{ $operation.Operation }}
~~

{{ if len $operation.Inputs }}
type {{ $operation.Name }}InputArgs struct {
{{- range $val := $operation.Inputs }}
{{ pascal $val.Name }} {{ $val.Type}} ~~json:"{{ $val.Name }}"~~
{{- end }}
}
{{ end }}

type {{ $operation.Name }}Response struct {
{{ $operation.ResponseType }}
}

{{ if len $operation.Inputs }}
func (c *client) {{ $operation.Name }}(ctx context.Context, input *{{ $operation.Name }}InputArgs) (
	*{{ $operation.Name }}Response, fezzik.GQLErrors, error) {
{{ else }}
func (c *client) {{ $operation.Name }}(ctx context.Context) (*{{ $operation.Name }}Response, fezzik.GQLErrors, error) {
{{ end }}

	gqlreq := &fezzik.GQLRequest{
		Operation: "{{ $operation.Name}}",
		Query: {{ $operation.Name }}Operation,
		Variables: map[string]interface{}{
			{{- range $val := $operation.Inputs }}	
			"{{ $val.Name }}" : input.{{ pascal $val.Name }},
			{{- end }}
		},
	}

	gqlresp := fezzik.GQLResponse{
		Data: &{{ $operation.Name }}Response{},
		Errors: fezzik.GQLErrors{},
	}
	err := c.gql.Query(ctx, gqlreq, gqlresp.Data, &gqlresp.Errors)
	if err != nil {
		return nil, nil, err
	}
	return gqlresp.Data.(*{{ $operation.Name }}Response), gqlresp.Errors, nil
}

{{- end }}
`

var clientTemplate string = `
package {{ .PackageName }}

import "github.com/machinebox/graphql"

type Client interface {
{{- range $operation := .Operations }}
{{- if len $operation.Inputs }}
	{{ $operation.Name }}(ctx context.Context, input *{{ $operation.Name }}InputArgs) (*{{ $operation.Name }}Response, fezzik.GQLErrors, error)
{{- else }}
	{{ $operation.Name }}(ctx context.Context) (*{{ $operation.Name }}Response, fezzik.GQLErrors, error)
{{- end }}	
{{- end }}
}

func NewClient(url string, httpclient *http.Client) Client {
	gqlclient := fezzik.NewGQLClient(url, httpclient)
	return &client{
		gql: gqlclient,
	}
}

type client struct {
	gql *fezzik.GQLClient
}
`

var TEMP string = `
	req := NewRequest({{ $operation.Name }}Operation)
	{{- range $val := $operation.Inputs }}
	q.Var("{{ $val.Name }}", input.{{ pascal $val.Name }})
	{{- end}}
	var resp map[string]interface{}
	err := c.gql.Run(ctx, q, &resp)

	{{- if $root.Debug }}
	log.Debug().Interface("resp", resp).Err(err).Msg("{{ $operation.Name }}")
	{{- end }}
	
	if err != nil {
		return nil, err
	}

	output := {{ $operation.Name }}Response{}
	err = mapstructure.Decode(resp, &output)
	
	{{- if $root.Debug }}
	log.Debug().Interface("output", output).Err(err).Msg("{{ $operation.Name }}")
	{{- end }}
	
	if err != nil {
		return nil, err
	}
	return &output, err
`
