package configuration

import (
	"time"

	"go.uber.org/zap"
)

type Config struct {
	Logger  zap.Config           `yaml:"logger"`
	Server  ServerConfiguration  `yaml:"server"`
	Storage StorageConfiguration `yaml:"storage"`
}

type ServerConfiguration struct {
	Host            string        `yaml:"host"`
	Port            int           `yaml:"port"`
	ShutdownTimeout time.Duration `yaml:"shutdownTimeout"`
	IdleTimeout     time.Duration `yaml:"idleTimeout"`
	WriteTimeout    time.Duration `yaml:"writeTimeout"`
	ReadTimeout     time.Duration `yaml:"readTimeout"`
}

type StorageConfiguration struct {
	UsePostgresStorage bool   `yaml:"usePostgres"`
	PostgresConnection string `yaml:"postgresConnection"`
}

func NewConfig(configFile string) (Config, error) {
	config := Config{}

	err := Configure(&config, configFile)
	if err != nil {
		return config, err
	}

	return config, nil
}
