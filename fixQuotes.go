package main

import (
	"fmt"
	"regexp"
)

func fixQuotes(s string) string {
	w := regexp.MustCompile(`\s*'\s*`)
	return w.ReplaceAllString(s, "'")
}

func main() {
	fmt.Println(fixQuotes("I am exactly how they describe me: ' awesome '"))
}
