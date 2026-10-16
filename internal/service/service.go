package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/morse-converter/pkg/morse"
)

var ErrEmptyInput = errors.New("входные данные пусты")

func DetectAndConvert(input string) (string, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return "", ErrEmptyInput
	}

	if isMorse(trimmed) {
		return morse.ToText(trimmed), nil
	}

	return morse.ToMorse(trimmed), nil
}

func isMorse(s string) bool {
	hasMorseChar := false
	for _, r := range s {
		switch r {
		case '.', '-':
			hasMorseChar = true
		case ' ', '\n', '\r', '\t':
			// разделители допустимы
		default:
			return false
		}
	}
	return hasMorseChar
}
