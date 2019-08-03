//Package isogram checks string to see if it's non-repeating patterns excluding punctuation.
package isogram

import "unicode"

//IsIsogram checks a string for duplicate rune values.
func IsIsogram(input string) bool {

	word := make(map[rune]struct{})
	var LETTER rune

	for _, letter := range input {
		LETTER = unicode.ToUpper(letter)

		if !unicode.IsLetter(LETTER) {
			continue
		}

		if _, ok := word[LETTER]; ok {
			return false
		}
		word[LETTER] = struct{}{}
	}
	return true
}
