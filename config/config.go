package config

type Config struct {
	Schemas    []string `yaml:"schemas"`
	Operations []string `yaml:"operations"`

	PackageName string `default:"fezzikclient" yaml:"package_name"`
	PackageDir  string `default:"fezzikclient" yaml:"package_dir"`

	ScalarTypeMap map[string]string `yaml:"scalar_type_map"`

	GenerateMocks bool `default:"false" yaml:"generate_mocks"`
	Debug         bool `default:"false" yaml:"debug"`
}
