package parse

import (
	"errors"
	"strings"

	"github.com/inigolabs/fezzik/common"
	"github.com/inigolabs/fezzik/config"
	"github.com/inigolabs/fezzik/fezzik_ast"
	"github.com/jensneuse/graphql-go-tools/pkg/ast"
	"github.com/jensneuse/graphql-go-tools/pkg/astparser"
	"github.com/jensneuse/graphql-go-tools/pkg/astvalidation"
	"github.com/jensneuse/graphql-go-tools/pkg/astvisitor"
	"github.com/jensneuse/graphql-go-tools/pkg/operationreport"

	"github.com/rs/zerolog/log"
)

func NewOperationParser(cfg *config.Config, schema *ast.Document, doc *fezzik_ast.Document) *operationParser {

	walker := astvisitor.NewWalker(48)
	visitor := &operationVisitor{
		Walker:     &walker,
		cfg:        cfg,
		definition: schema,

		doc: doc,
	}
	walker.RegisterEnterOperationVisitor(visitor)
	walker.RegisterLeaveOperationVisitor(visitor)
	walker.RegisterSelectionSetVisitor(visitor)
	walker.RegisterFieldVisitor(visitor)
	walker.RegisterArgumentVisitor(visitor)
	walker.RegisterVariableDefinitionVisitor(visitor)

	return &operationParser{
		cfg:     cfg,
		schema:  schema,
		visitor: visitor,
		doc:     doc,
	}
}

func (p *operationParser) ParseOperation(operationStr string) error {
	parser := astparser.NewParser()
	operation := ast.NewDocument()
	operation.Input.ResetInputBytes([]byte(operationStr))
	report := operationreport.Report{}
	parser.Parse(operation, &report)
	if report.HasErrors() {
		return errors.New(report.Error())
	}

	validator := astvalidation.DefaultOperationValidator()
	report = operationreport.Report{}
	result := validator.Validate(operation, p.schema, &report)
	if result != astvalidation.Valid {
		return errors.New(report.Error())
	}

	p.visitor.operation = operation
	report = operationreport.Report{}
	p.visitor.Walker.Walk(operation, p.schema, &report)
	if report.HasErrors() {
		return errors.New(report.Error())
	}

	return nil
}

type operationParser struct {
	cfg    *config.Config
	schema *ast.Document

	visitor *operationVisitor

	doc *fezzik_ast.Document
}

type currOperation struct {
	name      string
	operation *builder
	response  *builder
	inputs    []fezzik_ast.OperationInput
}

type operationVisitor struct {
	*astvisitor.Walker

	cfg        *config.Config
	operation  *ast.Document
	definition *ast.Document

	doc *fezzik_ast.Document

	curr *currOperation
}

func (v *operationVisitor) EnterOperationDefinition(ref int) {
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

func (v *operationVisitor) LeaveOperationDefinition(ref int) {
	v.doc.Operations = append(v.doc.Operations, &fezzik_ast.OperationInfo{
		Name:         v.curr.name,
		Operation:    v.curr.operation.String(),
		ResponseType: v.curr.response.String(),
		Inputs:       v.curr.inputs,
	})
}

func (v *operationVisitor) EnterSelectionSet(ref int) {
	v.curr.operation.Writeln(" {")

	parentTypeName := v.definition.NodeNameUnsafeString(v.EnclosingTypeDefinition)
	if parentTypeName != "Query" && parentTypeName != "Mutation" {
		v.curr.response.Writeln(v.indent(), " {")
	}
}

func (v *operationVisitor) LeaveSelectionSet(ref int) {
	v.curr.operation.Write(v.indent(), "}")

	parentTypeName := v.definition.NodeNameUnsafeString(v.EnclosingTypeDefinition)
	if parentTypeName != "Query" && parentTypeName != "Mutation" {
		v.curr.response.Writeln(" }")
	}
}

func (v *operationVisitor) EnterField(ref int) {
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
	typeInfo := fezzik_ast.GetTypeInfo(fieldDefTypeRef, v.definition)

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
	if goType, ok := fezzik_ast.GetGoType(v.cfg, typeInfo.Name); ok {
		v.curr.response.Writeln(goType)
	} else if _, found := v.doc.EnumTypes[typeInfo.Name]; found {
		v.curr.response.Writeln(typeInfo.Name)
	} else {
		v.curr.response.Write("struct")
	}

	log.Debug().Str("name", string(name)).Str("type", typeInfo.Name).Msg("EnterField")
}

func (v *operationVisitor) LeaveField(ref int) {
	v.curr.operation.Writeln()
}

func (v *operationVisitor) EnterVariableDefinition(ref int) {
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
	typeInfo := fezzik_ast.GetTypeInfo(varDef.Type, v.operation)
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
	if goType, ok := fezzik_ast.GetGoType(v.cfg, typeInfo.Name); ok {
		typeBuilder.WriteString(goType)
	} else {
		typeBuilder.WriteString(typeInfo.Name)
	}

	v.curr.inputs = append(v.curr.inputs, fezzik_ast.OperationInput{
		Name: varNameStr,
		Type: typeBuilder.String(),
	})

	log.Debug().Str("name", varNameStr).Bytes("type", varType).Msg("EnterVariableDefinition")
}

func (v *operationVisitor) LeaveVariableDefinition(ref int) {
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

func (v *operationVisitor) EnterArgument(ref int) {
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

func (v *operationVisitor) LeaveArgument(ref int) {
	lastAncestor := v.Ancestors[len(v.Ancestors)-1]
	if len(v.operation.ArgumentsAfter(lastAncestor, ref)) == 0 {
		v.curr.operation.Write(")")
	}
}

func (v *operationVisitor) indent() []byte {

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

func check(err error) {
	if err != nil {
		panic(err)
	}
}
