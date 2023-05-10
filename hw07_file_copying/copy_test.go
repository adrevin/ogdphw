package main

import (
	"errors"
	"os"
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
		require.Equal(t, dstStat.Size(), int64(256))

		dstBytes := make([]byte, 256)

		n, err := dst.Read(dstBytes)
		require.Nil(t, err)
		require.Equal(t, n, 256)
		require.Equal(t, dstBytes, bytes)

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
		require.Equal(t, dstStat.Size(), int64(128))

		dstBytes := make([]byte, 128)

		n, err := dst.Read(dstBytes)
		require.Nil(t, err)
		require.Equal(t, n, 128)
		require.Equal(t, dstBytes, bytes[128:256])

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
		require.Equal(t, dstStat.Size(), int64(10))

		dstBytes := make([]byte, 10)

		n, err := dst.Read(dstBytes)
		require.Nil(t, err)
		require.Equal(t, n, 10)
		require.Equal(t, dstBytes, bytes[48:58])

		_ = os.Remove(src.Name())
		_ = os.Remove(dstName)
	})
}
