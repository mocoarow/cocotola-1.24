package config

import "github.com/gin-contrib/cors"

type CORSConfig struct {
	AllowOrigins []string `yaml:"allowOrigins"`
}

func InitCORS(cfg *CORSConfig) cors.Config {
	if len(cfg.AllowOrigins) == 1 && cfg.AllowOrigins[0] == "*" {
		return cors.Config{
			AllowAllOrigins: true,
			AllowMethods:    []string{"*"},
			AllowHeaders:    []string{"*"},
		}
	}

	return cors.Config{
		AllowOrigins: cfg.AllowOrigins,
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
	}
}
