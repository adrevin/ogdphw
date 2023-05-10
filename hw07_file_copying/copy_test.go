package main

import (
	"errors"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	const dstName = "/tmp/hw07_test_dst"
	bytes := make([]byte, 256)

	newSrc := func() (*os.File, error) {
		src, err := os.CreateTemp("/tmp", "hw07_test_src_")
		if err != nil {
			return nil, err
		}

		var i byte
		for {
			bytes[i] = i
			if i == 255 {
				break
			}
			i++
		}
		_, err = src.Write(bytes)
		if err != nil {
			return nil, err
		}
		err = src.Close()
		if err != nil {
			return nil, err
		}
		return src, nil
	}

	t.Run("unsupported file", func(t *testing.T) {
		err := Copy("./testdata/non-existent-file", "", 0, 0)
		require.Truef(t, errors.Is(err, ErrUnsupportedFile), "actual err - %v", err)
	})

	t.Run("offset exceeds file size", func(t *testing.T) {
		src, err := newSrc()
		require.Nil(t, err)
		err = Copy(src.Name(), "", 257, 0)
		require.Truef(t, errors.Is(err, ErrOffsetExceedsFileSize), "actual err - %v", err)
		_ = os.Remove(src.Name())
	})

	t.Run("simplest case", func(t *testing.T) {
		src, err := newSrc()
		require.Nil(t, err)

		require.Nil(t, err)
		err = Copy(src.Name(), dstName, 0, 256)
		require.Nil(t, err)

		dst, err := os.Open(dstName)
		require.Nil(t, err)

		dstStat, err := dst.Stat()
		require.Nil(t, err)
		require.Equal(t, int64(256), dstStat.Size())

		dstBytes := make([]byte, 256)

		n, err := dst.Read(dstBytes)
		require.Nil(t, err)
		require.Equal(t, 256, n)
		require.Equal(t, bytes, dstBytes)

		_ = os.Remove(src.Name())
		_ = os.Remove(dstName)
	})

	t.Run("offset case", func(t *testing.T) {
		src, err := newSrc()
		require.Nil(t, err)

		require.Nil(t, err)
		err = Copy(src.Name(), dstName, 128, 256)
		require.Nil(t, err)

		dst, err := os.Open(dstName)
		require.Nil(t, err)

		dstStat, err := dst.Stat()
		require.Nil(t, err)
		require.Equal(t, int64(128), dstStat.Size())

		dstBytes := make([]byte, 128)

		n, err := dst.Read(dstBytes)
		require.Nil(t, err)
		require.Equal(t, 128, n)
		require.Equal(t, bytes[128:256], dstBytes)

		_ = os.Remove(src.Name())
		_ = os.Remove(dstName)
	})

	t.Run("offset limit case", func(t *testing.T) {
		src, err := newSrc()
		require.Nil(t, err)

		require.Nil(t, err)
		err = Copy(src.Name(), dstName, 48, 10)
		require.Nil(t, err)

		dst, err := os.Open(dstName)
		require.Nil(t, err)

		dstStat, err := dst.Stat()
		require.Nil(t, err)
		require.Equal(t, int64(10), dstStat.Size())

		dstBytes := make([]byte, 10)

		n, err := dst.Read(dstBytes)
		require.Nil(t, err)
		require.Equal(t, 10, n)
		require.Equal(t, bytes[48:58], dstBytes)

		_ = os.Remove(src.Name())
		_ = os.Remove(dstName)
	})

	t.Run("from test.sh", func(t *testing.T) {
		err := Copy("testdata/input.txt", "out.txt", 0, 0)
		require.Nil(t, err)
		cmd := exec.Command("cmp", "out.txt", "testdata/out_offset0_limit0.txt")
		err = cmd.Run()
		require.Nil(t, err)
	})
}
