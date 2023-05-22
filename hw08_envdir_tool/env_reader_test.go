package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const testDataPath = "./testdata/env"

func TestReadDir(t *testing.T) {
	tmp, err := os.CreateTemp(testDataPath, "B=A=R.")
	require.NoErrorf(t, err, "can not create temporary file")

	t.Run("test ReadDir method", func(t *testing.T) {
		evn, err := ReadDir(testDataPath)
		require.NoErrorf(t, err, "can not read dir")
		require.Equal(t, 5, len(evn))
		require.Equal(t, "bar", evn["BAR"].Value)
		require.Equal(t, "", evn["EMPTY"].Value)
		require.Equal(t, "   foo\nwith new line", evn["FOO"].Value)
		require.Equal(t, "\"hello\"", evn["HELLO"].Value)
		require.Equal(t, true, evn["UNSET"].NeedRemove)
	})
	err = os.Remove(tmp.Name())
	require.NoErrorf(t, err, "can not delete temporary file")
}
