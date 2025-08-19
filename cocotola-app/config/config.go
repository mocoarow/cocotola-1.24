package config

import (
	"embed"
	"os"

	// _ "embed"

	"go.yaml.in/yaml/v4"

	mblibconfig "github.com/mocoarow/cocotola-1.24/moonbeam/lib/config"
	mblibdomain "github.com/mocoarow/cocotola-1.24/moonbeam/lib/domain"
	mbliberrors "github.com/mocoarow/cocotola-1.24/moonbeam/lib/errors"

	libconfig "github.com/mocoarow/cocotola-1.24/lib/config"

	authconfig "github.com/mocoarow/cocotola-1.24/cocotola-auth/config"
	coreconfig "github.com/mocoarow/cocotola-1.24/cocotola-core/config"
)

type ServerConfig struct {
	HTTPPort             int `yaml:"httpPort" validate:"required"`
	MetricsPort          int `yaml:"metricsPort" validate:"required"`
	ReadHeaderTimeoutSec int `yaml:"readHeaderTimeoutSec" validate:"gte=1"`
}

type AuthAPIonfig struct {
	Endpoint string `yaml:"endpoint" validate:"required"`
	Username string `yaml:"username" validate:"required"`
	Password string `yaml:"password" validate:"required"`
}

type AppConfig struct {
	Auth *authconfig.AuthConfig `yaml:"auth" validate:"required"`
	Core *coreconfig.CoreConfig `yaml:"core" validate:"required"`
}

type Config struct {
	App      *AppConfig                 `yaml:"app" validate:"required"`
	Server   *ServerConfig              `yaml:"server" validate:"required"`
	DB       *mblibconfig.DBConfig      `yaml:"db" validate:"required"`
	Trace    *mblibconfig.TraceConfig   `yaml:"trace" validate:"required"`
	CORS     *mblibconfig.CORSConfig    `yaml:"cors" validate:"required"`
	Shutdown *libconfig.ShutdownConfig  `yaml:"shutdown" validate:"required"`
	Log      *mblibconfig.LogConfig     `yaml:"log" validate:"required"`
	Swagger  *mblibconfig.SwaggerConfig `yaml:"swagger" validate:"required"`
	Debug    *libconfig.DebugConfig     `yaml:"debug"`
}

//go:embed config.yml
var config embed.FS

func LoadConfig() (*Config, error) {
	filename := "config.yml"
	confContent, err := config.ReadFile(filename)
	if err != nil {
		return nil, mbliberrors.Errorf("config.ReadFile. filename: %s, err: %w", filename, err)
	}

	confContent = []byte(os.Expand(string(confContent), mblibconfig.ExpandEnvWithDefaults))
	conf := &Config{}
	if err := yaml.Unmarshal(confContent, conf); err != nil {
		return nil, mbliberrors.Errorf("yaml.Unmarshal. filename: %s, err: %w", filename, err)
	}

	if err := mblibdomain.Validator.Struct(conf); err != nil {
		return nil, mbliberrors.Errorf("Validator.Struct. filename: %s, err: %w", filename, err)
	}

	return conf, nil
}
