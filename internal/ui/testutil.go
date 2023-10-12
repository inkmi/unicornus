package ui

import (
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strings"
)

func Normalize(inputHTML string) string {
	r := strings.NewReader(inputHTML)
	doc, err := goquery.NewDocumentFromReader(r)
	if err == nil {
		return inputHTML
	}

	// Find all elements with class attributes
	doc.Find("[class]").Each(func(_ int, s *goquery.Selection) {
		// Remove the class attribute from the current element
		s.RemoveAttr("class")
	})

	// Print the modified HTML
	html, err := doc.Html()
	if err != nil {
		return err.Error()
	}
	return html
}

func Clean(inputHTML string) string {
	regexPattern := `>[[:space:]]+<`
	re := regexp.MustCompile(regexPattern)
	cleanedHTML := re.ReplaceAllString(inputHTML, "><")
	cleanedHTML = strings.TrimSpace(cleanedHTML)
	return cleanedHTML
}
