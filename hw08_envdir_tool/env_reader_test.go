package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const testDataPath = "./testdata/env"

func TestReadDir(t *testing.T) {
	tmp, _ := os.CreateTemp(testDataPath, "B=A=R.")

	t.Run("test ReadDir method", func(t *testing.T) {
		evn, err := ReadDir(testDataPath)
		require.Nil(t, err)
		require.Equal(t, 5, len(evn))
		require.Equal(t, "bar", evn["BAR"].Value)
		require.Equal(t, " ", evn["EMPTY"].Value)
		require.Equal(t, "   foo\nwith new line", evn["FOO"].Value)
		require.Equal(t, "\"hello\"", evn["HELLO"].Value)
		require.Equal(t, true, evn["UNSET"].NeedRemove)
	})
	os.Remove(tmp.Name())
}
