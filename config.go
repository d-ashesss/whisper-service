package main

import (
	"github.com/joeshaw/envdecode"
)

// Config contains app configuration options.
type Config struct {
	Port string `env:"PORT,default=8080"`
}

// NewConfig loads app configuration.
func NewConfig() *Config {
	var cfg Config
	_ = envdecode.Decode(&cfg)
	return &cfg
}
