package configuration

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfigurationProvider(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		tmp, err := os.CreateTemp("/tmp", "config.yml.")
		require.NoErrorf(t, err, "can not create temporary file")
		tmp.WriteString(`
logger:
  level: INFO`)

		config, err := NewConfig(tmp.Name())
		require.NoError(t, err)
		require.Equal(t, "INFO", config.Logger.Level)

		err = os.Remove(tmp.Name())
		require.NoErrorf(t, err, "can not delete temporary file")
	})
}
