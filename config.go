package fezzik

type Config struct {
	Schemas    []string `yaml:"schemas"`
	Operations []string `yaml:"operations"`

	PackageName string `yaml:"package_name"`
	PackageDir  string `yaml:"package_dir"`

	ScalarBuiltIn []string          `yaml:"scalar_built_in"`
	ScalarTypeMap map[string]string `yaml:"scalar_type_map"`

	DirectiveStructSuffix string
}
