package config

import (
	"fmt"
	"os"
)

// Config holds runtime configuration values.
type Config struct {
	Env  string
	Port string
}

// Load gathers configuration from environment variables with sane defaults.
func Load() Config {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return Config{
		Env:  env,
		Port: port,
	}
}

// Addr returns the TCP listen address for the HTTP server.
func (c Config) Addr() string {
	return fmt.Sprintf(":%s", c.Port)
}
