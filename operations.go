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
	PackageName string

	Operations []*OperationInfo
}

type OperationInfo struct {
	Name          string
	Operation     string
	ResponseType  string
	Inputs        string
	InputVarNames []string
}

type currOperation struct {
	name          string
	operation     *builder
	response      *builder
	inputs        *builder
	inputVarNames []string
}

type visitor struct {
	*astvisitor.Walker

	cfg        *Config
	operation  *ast.Document
	definition *ast.Document
	inputs     *InputRenderInfo

	info *Info

	curr *currOperation
}

func NewVisitor(cfg *Config, schema *ast.Document, inputsInfo *InputRenderInfo) *visitor {
	walker := astvisitor.NewWalker(48)
	visitor := &visitor{
		Walker:     &walker,
		cfg:        cfg,
		definition: schema,
		inputs:     inputsInfo,
		info: &Info{
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
		inputs:    NewBuilder(),
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
		Name:          v.curr.name,
		Operation:     v.curr.operation.String(),
		ResponseType:  v.curr.response.String(),
		Inputs:        v.curr.inputs.String(),
		InputVarNames: v.curr.inputVarNames,
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
		v.curr.response.Write("[]")
	}
	if typeInfo.IsTypeNullable {
		v.curr.response.Write("*")
	}
	if goType, ok := getGoType(v.cfg, typeInfo.Name); ok {
		v.curr.response.Writeln(goType)
	} else if _, found := v.inputs.EnumTypes[typeInfo.Name]; found {
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
	v.curr.inputVarNames = append(v.curr.inputVarNames, varNameStr)

	typeInfo := getTypeInfo(varDef.Type, v.operation)
	typeBuilder := new(strings.Builder)
	if typeInfo.IsList {
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
	v.curr.inputs.Write(common.UppercaseFirstChar(varNameStr), " ", typeBuilder.String())
	v.curr.inputs.Writeln(" `", "json:", `"`, varNameStr, `"`, "`")

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
