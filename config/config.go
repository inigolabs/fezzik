package config

type Config struct {
	Schemas    []string `yaml:"schemas"`
	Operations []string `yaml:"operations"`

	PackageName string `default:"fezzikclient" yaml:"package_name"`
	PackageDir  string `default:"fezzikclient" yaml:"package_dir"`

	// StructTagCase controls the format of the json struct tags
	//  valid options: snake, camel, kebob
	StructTagCase string `default:"camel" yaml:"struct_tag_case"`

	ScalarTypeMap map[string]string `yaml:"scalar_type_map"`

	GenerateMocks bool `default:"false" yaml:"generate_mocks"`
	Debug         bool `default:"false" yaml:"debug"`
}
