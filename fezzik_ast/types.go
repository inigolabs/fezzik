package fezzik_ast

import "github.com/inigolabs/fezzik/config"

func GetGoType(cfg *config.Config, typeName string) (string, bool) {
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
