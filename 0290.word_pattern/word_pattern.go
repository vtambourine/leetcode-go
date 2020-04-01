package word_pattern

import (
	"strings"
)

func wordPattern(pattern string, str string) bool {
	dict := strings.Split(str, " ")

	if len(dict) != len(pattern) {
		return false
	}

	chars := make(map[rune]string)
	words := make(map[string]rune)

	for i, p := range pattern {
		if c, ok := chars[p]; ok {
			if c != dict[i] {
				return false
			}
		} else if w, ok := words[dict[i]]; ok {
			if w != p {
				return false
			}
		}

		chars[p] = dict[i]
		words[dict[i]] = p
	}

	return true
}