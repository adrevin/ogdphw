package main

import (
	"errors"
	"io"
	"os"

	"github.com/schollz/progressbar/v3"
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

	to, err := os.Create(toPath)
	if err != nil {
		return err
	}
	_, err = from.Seek(offset, 0)
	if err != nil {
		return err
	}

	count := fromStat.Size() - offset
	if limit > 0 && count > limit {
		count = limit
	}

	bar := progressbar.DefaultBytes(count)
	_, err = io.CopyN(io.MultiWriter(to, bar), from, count)
	if err != nil {
		return err
	}

	err = to.Close()
	if err != nil {
		return err
	}
	return nil
}
