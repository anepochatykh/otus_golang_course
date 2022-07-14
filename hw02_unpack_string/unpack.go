package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func isValid(inp string) bool {
	alreadyHasDigit := false
	for i, r := range inp {
		// check if first symbol is digit
		if (i == 0) && (unicode.IsDigit(r)) {
			return false
		}
		if unicode.IsDigit(r) {
			// check if there are two digits
			if alreadyHasDigit {
				return false
			}
			alreadyHasDigit = true
		} else {
			alreadyHasDigit = false
		}
	}
	return true
}

func Unpack(inp string) (string, error) {
	var b strings.Builder

	if !isValid(inp) {
		return "", ErrInvalidString
	}
	multiplier := -1
	prevSymbol := ""
	for _, r := range inp {
		if unicode.IsDigit(r) {
			multiplier, err := strconv.Atoi(string(r))
			if err != nil {
				return "", ErrInvalidString
			}
			for i := 0; i < multiplier; i++ {
				b.WriteString(prevSymbol)
			}
			prevSymbol = ""
		} else {
			if prevSymbol != "" && multiplier == -1 {
				b.WriteString(prevSymbol)
			}
			prevSymbol = string(r)
		}
	}
	if prevSymbol != "" {
		b.WriteString(prevSymbol)
	}

	result := b.String()
	return result, nil
}
