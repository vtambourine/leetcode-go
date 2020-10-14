package detect_capital

import (
	"regexp"
)

func detectCapitalUse(word string) bool {
	allCaps := regexp.MustCompile(`^[A-Z]*$`)
	allLow := regexp.MustCompile(`^[a-z]*$`)
	firstCap := regexp.MustCompile(`^[A-Z][a-z]*$`)

	return allCaps.MatchString(word) || allLow.MatchString(word) || firstCap.MatchString(word)
}