package main

import (
	"fmt"
	"regexp"
)

func fixPunctuation(s string) string {
	w := regexp.MustCompile(`\s+([.,!?:;]+)`)
	s = w.ReplaceAllString(s, `$1`)

	m := regexp.MustCompile(`([.,!?:;]+)([^\s.,!?:;])`)
	s = m.ReplaceAllString(s, `$1 $2`)

	return s

}

func main() {
	fmt.Println(fixPunctuation("I was sitting over there ,and ...then BAMM !!"))
}
