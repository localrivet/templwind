package config

import "gopkg.in/yaml.v3"

type Config struct {
	Port int `yaml:"Port"`
}

// This function now explicitly expects a pointer to Config rather than any type.
func LoadConfigFromYamlBytes(bytes []byte, cfg *Config) error {
	err := yaml.Unmarshal(bytes, cfg)
	if err != nil {
		return err
	}

	return nil
}
