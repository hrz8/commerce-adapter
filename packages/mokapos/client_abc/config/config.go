package config

import (
	"context"
	"log"

	"github.com/sethvargo/go-envconfig"
)

// Config is a struct that represents all available application's config sets.
type Config struct {
	BaseURL      string `env:"BASE_URL,required"`
	FunctionName string `env:"FUNCTION_NAME,required"`
}

// New is a constructor that return an instance of Config struct.
func New() *Config {
	cfg := &Config{}

	ctx := context.Background()
	if err := envconfig.Process(ctx, cfg); err != nil {
		log.Fatal("cannot load config")
	}

	return cfg
}

// GetBaseURL is a getter to get base url config.
func (c *Config) GetBaseURL() string {
	return c.BaseURL
}

// GetFunctionName is a getter to get function name config.
func (c *Config) GetFunctionName() string {
	return c.FunctionName
}
