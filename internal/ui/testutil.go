package ui

import (
	"regexp"
	"strings"
)

func Clean(inputHTML string) string {
	regexPattern := `>[[:space:]]+<`
	re := regexp.MustCompile(regexPattern)
	cleanedHTML := re.ReplaceAllString(inputHTML, "><")
	cleanedHTML = strings.TrimSpace(cleanedHTML)
	return cleanedHTML
}
