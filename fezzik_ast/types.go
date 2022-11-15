package fezzik_ast

import (
	"fmt"
	"strings"

	"github.com/inigolabs/fezzik/config"
)

func GetGoType(cfg *config.Config, typeName string) (string, bool) {
	var builtinTypeMap = map[string]string{
		"Boolean": "bool",
		"Float":   "float32",
		"Int":     "int",
		"String":  "string",
	}

	if val, found := builtinTypeMap[typeName]; found {
		return val, true
	}
	if val, found := cfg.ScalarTypeMap[typeName]; found {
		_, scalarTypeName := PkgAndType(val)
		return scalarTypeName, true
	}
	return "", false
}

func (s Source) String() string {
	return fmt.Sprintf("%s:%d", s.FileName, s.Line)
}

// take a string in the form github.com/package/foo.Bar and split it into package and type
func PkgAndType(name string) (string, string) {
	parts := strings.Split(name, ".")
	if len(parts) == 1 {
		return "", name
	}

	pkgParts := strings.Split(strings.Join(parts[:len(parts)-1], "."), "/")

	typeName := strings.Join([]string{pkgParts[len(pkgParts)-1], parts[len(parts)-1]}, ".")

	return strings.Join(parts[:len(parts)-1], "."), typeName
}
