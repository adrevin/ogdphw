package configuration

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestConfigurationProvider(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		tmp, err := os.CreateTemp("/tmp", "config.yml.")
		require.NoErrorf(t, err, "can not create temporary file")
		_, err = tmp.WriteString(`
logger:
  level: debug
  encoding: console
  encoderConfig:
    timeKey:        "T"
    levelKey:       "L"
    nameKey:        "N"
    callerKey:      "C"
    messageKey:     "M"
    stacktraceKey:  "S"
httpServer:
  host: "0.0.0.0"
  port: 5000
  readTimeout: 15s
  writeTimeout: 60s
  idleTimeout: 5s
storage:
  usePostgres: false
  postgresConnection: ""
shutdownTimeout: 6s
`)
		require.NoErrorf(t, err, "can not write temporary file")

		config, err := NewConfig(tmp.Name())
		require.NoError(t, err)
		require.Equal(t, zap.NewAtomicLevelAt(zap.DebugLevel), config.Logger.Level)
		require.Equal(t, "T", config.Logger.EncoderConfig.TimeKey)
		require.Equal(t, false, config.Storage.UsePostgresStorage)
		require.Equal(t, 5000, config.HTTPServer.Port)
		require.Equal(t, 6*time.Second, config.ShutdownTimeout)

		err = os.Remove(tmp.Name())
		require.NoErrorf(t, err, "can not delete temporary file")
	})
}
