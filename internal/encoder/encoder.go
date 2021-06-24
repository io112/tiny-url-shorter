package encoder

import (
	"errors"
	"math"
)

// Number of runes in alphabet.
var base = int64(1)

// alphabet of runes for Decode strings.
var alphabet = make(map[rune]int)

// list of runes in alphabet for Encode integers.
var alphabetRunes = make([]rune, base)

// ErrInvalidRunes is returned by Decode if there is invalid characters in given string.
var ErrInvalidRunes = errors.New("encoder: Invalid characters in string")

// ErrBelowZero is returned by Encode if given argument is below zero.
var ErrBelowZero = errors.New("encoder: Argument must be greater or equal than zero")

// init generates basic structures on package import.
func init() {
	generateAlphabet()
}

// generateAlphabet generates base, alphabet and alphabetRunes for Encode and Decode functions.
func generateAlphabet() {
	runes := append(makeRange('0', '9'), makeRange('a', 'z')...)
	runes = append(runes, makeRange('A', 'Z')...)
	base = int64(len(runes))
	alphabetRunes = runes
	for i, char := range alphabetRunes {
		alphabet[char] = i
	}
}

// makeRange generates range of runes from min rune to max.
func makeRange(min rune, max rune) []rune {
	a := make([]rune, max-min+1)
	for i := range a {
		a[i] = min + rune(i)
	}
	return a
}

// Decode decodes given string if string is valid.
func Decode(encodedString string) (*int64, error) {
	result := int64(0)
	for i, char := range encodedString {
		val, ok := alphabet[char]
		if !ok {
			return nil, ErrInvalidRunes
		}
		result += int64(val) * int64(math.Pow(float64(base), float64(len(encodedString)-1-i)))
	}
	return &result, nil
}

// Encode encodes int64 if integer greater or equal zero.
func Encode(decodedInt int64) (*string, error) {
	result := ""
	if decodedInt < 0 {
		return nil, ErrBelowZero
	}
	for decodedInt > 0 {
		result = string(alphabetRunes[decodedInt%base]) + result
		decodedInt = decodedInt / base
	}

	return &result, nil
}
