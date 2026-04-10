package main

import (
	"regexp"
)

func fixQuotes(s string) string {
	w := regexp.MustCompile(`\s*'\s*`)
	return w.ReplaceAllString(s, "'")
}
