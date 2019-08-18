package secrets

import (
	"os"
	"strings"
)

// EnvSecrets retrieves secrets from environment variables
type EnvSecrets struct {
}

// NewEnvSecrets returns an instance
func NewEnvSecrets() *EnvSecrets {
	return &EnvSecrets{}
}

// GetSecureParameter retrieves the env variable specified by key, or error otherwise.
func (s *EnvSecrets) GetSecureParameter(key string) ([]byte, error) {
	key = strings.Replace(key, "/", "_", -1)
	data := os.Getenv(key)
	return []byte(data), nil
}

// SetSecureParameter sets the value for the key specified by environment variable
func (s *EnvSecrets) SetSecureParameter(key, value string) error {
	key = strings.Replace(key, "/", "_", -1)
	os.Setenv(key, value)
	return nil
}
