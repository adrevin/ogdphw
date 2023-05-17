package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReadDir(t *testing.T) {
	t.Run("unsupported file", func(t *testing.T) {
		evn, err := ReadDir("./testdata/env")
		require.Nil(t, err)
		require.Equal(t, 5, len(evn))
		require.Equal(t, "bar", evn["BAR"].Value)
		require.Equal(t, " ", evn["EMPTY"].Value)
		require.Equal(t, "   foo\x00with new line", evn["FOO"].Value)
		require.Equal(t, "\"hello\"", evn["HELLO"].Value)
		require.Equal(t, true, evn["UNSET"].NeedRemove)
	})
}
