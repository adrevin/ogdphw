package main

import (
	"bufio"
	"errors"
	"io"
	"os"
	"path"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

var ErrEmptyFile = errors.New("file is empty")

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	env := Environment{}
	entries, err := os.ReadDir(dir)
	if err != nil {
		return env, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		val, err := readFirstLine(path.Join(dir, entry.Name()))
		isEmptyFile := errors.Is(err, ErrEmptyFile)
		if !isEmptyFile && err != nil {
			return env, err
		}

		env[entry.Name()] = EnvValue{Value: val, NeedRemove: isEmptyFile}
	}
	return env, nil
}

func readFirstLine(filepath string) (string, error) {
	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		return "", err
	}
	val, _, err := bufio.NewReader(file).ReadLine()
	if errors.Is(err, io.EOF) {
		return "", ErrEmptyFile
	}
	if err != nil {
		return "", err
	}
	return string(val), nil
}
