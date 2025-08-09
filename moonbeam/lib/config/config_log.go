package config

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
)

type LogConfig struct {
	Level    string `yaml:"level"`
	Platform string `yaml:"platform"`
}

func newReplaceAttr(platform string) func([]string, slog.Attr) slog.Attr {
	switch platform {
	case "gcp":
		return func(groups []string, a slog.Attr) slog.Attr {
			switch a.Key {
			case slog.LevelKey:
				return slog.Attr{Key: "severity", Value: a.Value}
			case slog.MessageKey:
				return slog.Attr{Key: "message", Value: a.Value}
			}

			return a
		}
	}
	return nil
}

func InitLog(cfg *LogConfig) {
	defaultLogLevel := stringToLogLevel(cfg.Level)

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:       defaultLogLevel,
		ReplaceAttr: newReplaceAttr(cfg.Platform),
	})))
}

func stringToLogLevel(value string) slog.Level {
	switch strings.ToLower(value) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		slog.Info(fmt.Sprintf("Unsupported log level: %s", value))
		return slog.LevelWarn
	}
}
