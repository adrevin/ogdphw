package hw02unpackstring

import (
	"errors"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(inp string) (string, error) {
	var builder strings.Builder
	var char *rune

	for _, c := range inp {
		// it is a digit: 48 - "0", 49 - "1", 57 - "9"
		if c > 47 && c < 58 {
			// it is no previous non digit char
			if char == nil {
				return "", ErrInvalidString
			}
			// repeat previous char
			for t := 0; t < int(c-48); t++ {
				builder.WriteRune(*char)
			}
			char = nil // reset repeated char
		} else {
			// it is previous non digit char. just use it to output
			if char != nil {
				builder.WriteRune(*char)
			}
			r := c
			char = &r // remember char as previous
		}
	}

	// it is previous non digit char. just use it to output
	if char != nil {
		builder.WriteRune(*char)
	}

	return builder.String(), nil
}
