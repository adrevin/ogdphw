package hw02unpackstring

import (
	"errors"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(inp string) (string, error) {
	var builder strings.Builder
	var char string
	for _, c := range inp {
		if c > 47 && c < 58 { // it is a digit
			if char == "" { // it is no previous non digit char
				return "", ErrInvalidString
			}
			builder.WriteString(strings.Repeat(char, int(c-48))) // repeat previous char
			char = ""                                            // reset repeated char
		} else {
			builder.WriteString(char) // it is previous non digit char. just use it to output
			char = string(c)          // remember char as previous
		}
	}
	builder.WriteString(char) // process last char
	return builder.String(), nil
}
