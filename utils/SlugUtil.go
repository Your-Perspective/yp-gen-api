package utils

import (
	"regexp"
	"strings"
	"unicode"
)

// ContainsKhmer checks if a string contains Khmer characters
func ContainsKhmer(input string) bool {
	if input == "" {
		return false
	}

	// Define a function to check if a rune is in the Khmer range
	for _, r := range input {
		if r >= '\u1780' && r <= '\u17FF' {
			return true
		}
	}
	return false
}

// Init normalizes and sanitizes a string to create a slug
func Init(input string) string {
	// Normalize and remove non-ASCII characters except for Khmer characters
	normalized := removeDiacritics(input)

	// Replace spaces and slashes with hyphens
	normalized = regexp.MustCompile(`[\s/]+`).ReplaceAllString(normalized, "-")

	// Remove any character that is not a letter, number, hyphen, or Khmer character
	normalized = regexp.MustCompile(`[^a-zA-Z0-9\p{L}-]+`).ReplaceAllString(normalized, "")

	// Remove consecutive hyphens
	normalized = regexp.MustCompile(`-{2,}`).ReplaceAllString(normalized, "-")

	// Trim any leading or trailing hyphens
	normalized = strings.Trim(normalized, "-")

	return strings.ToLower(normalized)
}

// RemoveKhmerCharacters removes Khmer characters from a string
func RemoveKhmerCharacters(input string) string {
	return strings.Map(func(r rune) rune {
		if r >= '\u1780' && r <= '\u17FF' {
			return -1 // Remove the character
		}
		return r
	}, input)
}

// Helper function to remove diacritics from a string (normalization)
func removeDiacritics(input string) string {
	var output []rune
	for _, r := range input {
		if unicode.IsLetter(r) && (unicode.Is(unicode.Khmer, r) || isASCII(r)) {
			output = append(output, unicode.ToLower(r))
		}
	}
	return string(output)
}

// isASCII checks if a rune is an ASCII character
func isASCII(r rune) bool {
	return r <= 0x7F
}
