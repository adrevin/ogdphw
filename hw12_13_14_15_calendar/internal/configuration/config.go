package configuration

import "go.uber.org/zap"

type Config struct {
	Logger zap.Config `yaml:"logger"`
}

func NewConfig(configFile string) (Config, error) {
	config := Config{}

	err := Configure(&config, configFile)
	if err != nil {
		return config, err
	}

	return config, nil
}
