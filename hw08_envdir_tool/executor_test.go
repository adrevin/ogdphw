package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("", func(t *testing.T) {
		env, err := ReadDir(testDataPath)
		require.NoErrorf(t, err, "can not read dir")

		code := RunCmd([]string{"/bin/bash", "./testdata/echo.sh", "arg1=1", "arg2=2"}, env)
		require.Equal(t, 0, code)
	})
}
