package utils

import (
	"fmt"
	"github.com/gosimple/slug"
	"math/rand"
	"strings"
	"time"
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

// RemoveKhmerCharacters removes Khmer characters from a string
func RemoveKhmerCharacters(input string) string {
	return strings.Map(func(r rune) rune {
		if r >= '\u1780' && r <= '\u17FF' {
			return -1 // Remove the character
		}
		return r
	}, input)
}

// GenerateUniqueIdentifier Generates a 9-digit unique identifier
func GenerateUniqueIdentifier() string {
	rand.Int63n(time.Now().UnixNano())
	return fmt.Sprintf("%09d", rand.Intn(1e9))
}

// Init generates a slug using the simple/slug package
func Init(input string) string {
	return slug.Make(input)
}
