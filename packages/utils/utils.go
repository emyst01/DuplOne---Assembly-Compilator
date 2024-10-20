package utils

import (
	"strconv"
	"strings"
	"unicode"
)

func IsInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func IsValidNumberFormat(s string) bool {
	if strings.HasPrefix(s, "<") && strings.HasSuffix(s, ">") {
		numberPart := s[1 : len(s)-1]
		return IsInt(numberPart)
	}
	return false
}

func CleanString(input string) string {
	trimmed := strings.TrimSpace(input)
	fields := strings.Fields(trimmed)
	result := strings.Join(fields, " ")
	result = strings.ToLower(result)
	return result
}

func IsUint8(s string) bool {
	value, err := strconv.ParseUint(s, 10, 8)
	return err == nil && value <= 255
}
