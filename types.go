package fezzik

import "github.com/jensneuse/graphql-go-tools/pkg/ast"

type TypeInfo struct {
	Name           string
	IsList         bool
	IsListNullable bool
	IsTypeNullable bool
}

func getTypeInfo(ref int, doc *ast.Document) *TypeInfo {
	return getTypeInfoHelper(ref, doc, &TypeInfo{}, true)
}

func getTypeInfoHelper(ref int, doc *ast.Document, info *TypeInfo, nullable bool) *TypeInfo {
	switch doc.Types[ref].TypeKind {
	case ast.TypeKindNonNull:
		return getTypeInfoHelper(doc.Types[ref].OfType, doc, info, false)
	case ast.TypeKindList:
		info.IsList = true
		info.IsListNullable = nullable
		return getTypeInfoHelper(doc.Types[ref].OfType, doc, info, true)
	case ast.TypeKindNamed:
		info.IsTypeNullable = nullable
		info.Name = doc.Input.ByteSliceString(doc.Types[ref].Name)
	}
	return info
}

func getGoType(cfg *Config, typeName string) (string, bool) {
	var builtinTypeMap = map[string]string{
		"Boolean": "bool",
		"Float":   "float32",
		"ID":      "string",
		"Int":     "int",
		"String":  "string",
	}

	if val, found := builtinTypeMap[typeName]; found {
		return val, true
	}
	if val, found := cfg.ScalarTypeMap[typeName]; found {
		return val, true
	}
	return "", false
}
