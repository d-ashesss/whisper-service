package main

import (
	"github.com/joeshaw/envdecode"
)

type Config struct {
	Port string `env:"PORT,default=8080"`
}

func NewConfig() *Config {
	var cfg Config
	_ = envdecode.Decode(&cfg)
	return &cfg
}
