package config

import (
	"os"
	"strings"
)

// ExpandEnvWithDefaults expands environment variables in the format VAR_NAME:-default_value.
func ExpandEnvWithDefaults(varName string) string {
	// Check if it contains :-
	if strings.Contains(varName, ":-") {
		parts := strings.SplitN(varName, ":-", 2)
		name := parts[0]
		defaultValue := parts[1]

		if value := os.Getenv(name); value != "" {
			return value
		}

		return defaultValue
	}

	// Simple variable expansion
	return os.Getenv(varName)
}
