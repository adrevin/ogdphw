package configuration

type Config struct {
	Logger LoggerConf
	// TODO
}

type LoggerConf struct {
	Level string
}

func NewConfig(configFile string) (Config, error) {
	config := Config{}

	err := Configure(&config, configFile)
	if err != nil {
		return config, err
	}

	return config, nil
}
