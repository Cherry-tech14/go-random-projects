package main

import (
	"regexp"
)

func fixPunctuation(s string) string {
	w := regexp.MustCompile(`\s+([.,!?:;]+)`)
	s = w.ReplaceAllString(s, `$1`)

	m := regexp.MustCompile(`([.,!?:;]+)([^\s.,!?:;])`)
	s = m.ReplaceAllString(s, `$1 $2`)

	return s

}
