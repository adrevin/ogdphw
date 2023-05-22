package configuration

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestConfigurationProvider(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		tmp, err := os.CreateTemp("/tmp", "config.yml.")
		require.NoErrorf(t, err, "can not create temporary file")
		_, err = tmp.WriteString(`
logger:
  level: info

`)
		require.NoErrorf(t, err, "can not write temporary file")

		config, err := NewConfig(tmp.Name())
		require.NoError(t, err)
		require.Equal(t, zap.NewAtomicLevelAt(zap.InfoLevel), config.Logger.Level)

		err = os.Remove(tmp.Name())
		require.NoErrorf(t, err, "can not delete temporary file")
	})
}
