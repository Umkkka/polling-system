package config

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ListenPort      string        `default:"8080"              envconfig:"listen_port"`
	ShutdownTimeout time.Duration `default:"15s"`
	WriteTimeout    time.Duration `default:"15s"`
	ReadTimeout     time.Duration `default:"15s"`
	IdleTimeout     time.Duration `default:"60s"`
	LogLevel        string        `default:"info"              envconfig:"log_level"`
	Debug           bool          `default:"false"             envconfig:"debug"`
}

func ParseConfig() (Config, error) {
	var conf Config

	err := envconfig.Process("", &conf)
	if err != nil {
		return conf, fmt.Errorf("error parsing startup configuration: %w", err)
	}

	return conf, err
}
