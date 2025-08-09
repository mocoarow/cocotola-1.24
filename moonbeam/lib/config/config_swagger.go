package config

type SwaggerConfig struct {
	Enabled bool   `yaml:"enabled"`
	Host    string `yaml:"host"`
	Schema  string `yaml:"schema"`
}
