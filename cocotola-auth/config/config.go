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
)

type ServerConfig struct {
	HTTPPort             int `yaml:"httpPort" validate:"required"`
	MetricsPort          int `yaml:"metricsPort" validate:"required"`
	ReadHeaderTimeoutSec int `yaml:"readHeaderTimeoutSec" validate:"gte=1"`
}

type CoreAPIClientConfig struct {
	Endpoint   string `yaml:"endpoint" validate:"required"`
	Username   string `yaml:"username" validate:"required"`
	Password   string `yaml:"password" validate:"required"`
	TimeoutSec int    `yaml:"timeoutSec" validate:"gte=1"`
}

type AuthAPIServerConfig struct {
	Username string `yaml:"username" validate:"required"`
	Password string `yaml:"password" validate:"required"`
}

type AuthConfig struct {
	CoreAPIClient       *CoreAPIClientConfig `yaml:"coreApiClient" validate:"required"`
	AuthAPIServer       *AuthAPIServerConfig `yaml:"authApiServer" validate:"required"`
	SigningKey          string               `yaml:"signingKey" validate:"required"`
	AccessTokenTTLMin   int                  `yaml:"accessTokenTtlMin" validate:"gte=1"`
	RefreshTokenTTLHour int                  `yaml:"refreshTokenTtlHour" validate:"gte=1"`
	GoogleProjectID     string               `yaml:"googleProjectId" validate:"required"`
	GoogleCallbackURL   string               `yaml:"googleCallbackUrl" validate:"required"`
	GoogleClientID      string               `yaml:"googleClientId" validate:"required"`
	GoogleClientSecret  string               `yaml:"googleClientSecret" validate:"required"`
	GoogleAPITimeoutSec int                  `yaml:"googleApiTimeoutSec" validate:"gte=1"`
	OwnerLoginID        string               `yaml:"ownerLoginId" validate:"required"`
	OwnerPassword       string               `yaml:"ownerPassword" validate:"required"`
}

type Config struct {
	App      *AuthConfig                `yaml:"app" validate:"required"`
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
