package main

import (
	"errors"
	"io"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	from, err := os.Open(fromPath)
	if err != nil {
		return ErrUnsupportedFile
	}
	fromStat, err := from.Stat()
	if err != nil {
		return ErrUnsupportedFile
	}
	if offset > fromStat.Size() {
		return ErrOffsetExceedsFileSize
	}

	// simplest io.Copy
	if offset == 0 && limit >= fromStat.Size() {
		to, err := os.Create(toPath)
		if err != nil {
			return err
		}
		_, err = io.Copy(to, from)
		if err != nil {
			return err
		}
		err = to.Close()
		if err != nil {
			return err
		}
		return nil
	}

	// with offset/limit via buffer
	to, err := os.Create(toPath)
	if err != nil {
		return err
	}
	_, err = from.Seek(offset, 0)
	if err != nil {
		return err
	}

	buf := make([]byte, 3)
	var copied int64
	for copied < limit {
		n, err := from.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		w := limit - copied
		if w > int64(n) {
			w = int64(n)
		}

		if _, err := to.Write(buf[:w]); err != nil {
			return err
		}

		copied += int64(n)
	}
	err = to.Close()
	if err != nil {
		return err
	}
	return nil
}
