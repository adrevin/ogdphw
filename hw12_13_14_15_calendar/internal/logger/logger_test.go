package logger

import (
	"bytes"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/configuration"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

const yaml = `
logger:
  level: info
  encoding: console
  encoderConfig:
    timeKey:        ""
    levelKey:       "L"
    nameKey:        ""
    callerKey:      ""
    messageKey:     "M"
    stacktraceKey:  ""`

const requiredOut = `info	informational message
warn	warning message
error	error message	{"error": "error"}
`

func TestLogger(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		tmp, err := os.CreateTemp("/tmp", "config.yml.")
		require.NoErrorf(t, err, "can not create temporary file")
		_, err = tmp.WriteString(yaml)
		require.NoErrorf(t, err, "can not write temporary file")

		config, err := configuration.NewConfig(tmp.Name())
		require.NoErrorf(t, err, "can not delete temporary file")

		err = os.Remove(tmp.Name())
		require.NoErrorf(t, err, "can not delete temporary file")

		old := os.Stdout
		r, w, err := os.Pipe()
		require.NoError(t, err)
		os.Stdout = w

		print()

		outC := make(chan string)
		go func() {
			var buf bytes.Buffer
			io.Copy(&buf, r)
			outC <- buf.String()
		}()

		logger := New(config.Logger)
		logger.Debug("debug message")
		logger.Info("informational message")
		logger.Warn("warning message")
		logger.Error("error message", zap.Error(errors.New("error")))

		w.Close()
		os.Stdout = old
		out := <-outC

		require.Equal(t, requiredOut, out)
	})
}
