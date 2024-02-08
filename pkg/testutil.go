package pkg

import (
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strings"
)

func Normalize(inputHtml string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(inputHtml))
	if err != nil {
		return err.Error()
	}

	doc.Find("div").Each(func(index int, div *goquery.Selection) {
		// Move the children of the div element to the parent
		parent := div.Parent()
		parent.AppendSelection(div.ChildrenFiltered("*"))
		div.Remove()
	})
	// Also remove all class
	doc.Find("[class]").Each(func(_ int, s *goquery.Selection) {
		// Remove the class attribute from the current element
		s.RemoveAttr("class")
	})

	html, err := doc.Find("body").Html()
	if err != nil {
		return err.Error()
	}
	return html
}

func RemoveClassAndStyle(inputHTML string) string {
	r := strings.NewReader(inputHTML)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return err.Error()
	}

	// Find all elements with class attributes
	doc.Find("[class]").Each(func(_ int, s *goquery.Selection) {
		// Remove the class attribute from the current element
		s.RemoveAttr("class")
	})
	doc.Find("[style]").Each(func(_ int, s *goquery.Selection) {
		// Remove the class attribute from the current element
		s.RemoveAttr("style")
	})
	// Print the modified HTML
	html, err := doc.Find("body").Html()
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
