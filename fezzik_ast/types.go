package fezzik_ast

import (
	"fmt"

	"github.com/inigolabs/fezzik/config"
)

var builtinTypeMap = map[string]string{
	"Boolean": "bool",
	"Float":   "float32",
	"Int":     "int",
	"String":  "string",
}

func GetGoType(cfg *config.Config, typeName string) (string, bool) {
	if val, found := builtinTypeMap[typeName]; found {
		return val, true
	}
	if val, found := cfg.ScalarTypeMap[typeName]; found {
		return val, true
	}
	return "", false
}

func (s Source) String() string {
	return fmt.Sprintf("%s:%d", s.FileName, s.Line)
}
