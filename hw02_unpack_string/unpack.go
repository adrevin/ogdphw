package hw02unpackstring

import (
	"errors"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(inp string) (string, error) {
	var builder strings.Builder
	var char *rune = nil
	for _, c := range inp {
		if c > 47 && c < 58 { // it is a digit
			if char == nil { // it is no previous non digit char
				return "", ErrInvalidString
			}
			for t := 0; t < int(c-48); t++ { // repeat previous char
				builder.WriteRune(*char)
			}
			char = nil // reset repeated char
		} else {
			if char != nil {
				builder.WriteRune(*char) // it is previous non digit char. just use it to output
			}
			r := c
			char = &r // remember char as previous
		}
	}
	if char != nil {
		builder.WriteRune(*char) // it is previous non digit char. just use it to output
	}
	return builder.String(), nil
}
