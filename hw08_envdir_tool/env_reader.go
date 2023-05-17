package main

import (
	"bufio"
	"bytes"
	"errors"
	"os"
	"path"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

var (
	ErrEmptyFile = errors.New("file is empty")
	replaceOld   = []byte{0}
	replaceNew   = []byte("\n")
	tmp          = []byte{32}
)

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	t := string(tmp)
	_ = t
	env := Environment{}
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		entryName := entry.Name()
		if strings.Contains(entryName, "=") {
			continue
		}
		val, err := getFirstLine(path.Join(dir, entryName))
		isEmptyFile := errors.Is(err, ErrEmptyFile)
		if !isEmptyFile && err != nil {
			return nil, err
		}

		env[entry.Name()] = EnvValue{Value: val, NeedRemove: isEmptyFile}
	}

	return env, nil
}

func getFirstLine(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return "", err
	}
	if fileInfo.Size() == 0 {
		return "", ErrEmptyFile
	}

	val, _, err := bufio.NewReader(file).ReadLine()
	if err != nil {
		return "", err
	}

	return sanitize(val), nil
}

func sanitize(b []byte) string {
	b = bytes.Replace(b, replaceOld, replaceNew, -1)
	o := strings.TrimRight(string(b), "\n")
	return strings.TrimRight(o, " ")
}
