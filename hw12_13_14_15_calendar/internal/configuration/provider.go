package configuration

import (
	"os"

	yaml "gopkg.in/yaml.v3"
)

func Configure[T interface{}](t *T, configFile string) error {
	bytes, err := os.ReadFile(configFile)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(bytes, &t)
	if err != nil {
		return err
	}

	return nil
}
