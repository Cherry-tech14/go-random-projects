package main

import (
	"strings"
)

func vowelA2An(nextWord string) string {
	nextWord = strings.ReplaceAll(nextWord, "A a", "An a")
	nextWord = strings.ReplaceAll(nextWord, "A h", "An h")
	nextWord = strings.ReplaceAll(nextWord, "A e", "An e")
	nextWord = strings.ReplaceAll(nextWord, "A h", "An h")
	nextWord = strings.ReplaceAll(nextWord, "A i", "An i")
	nextWord = strings.ReplaceAll(nextWord, "A u", "An u")
	nextWord = strings.ReplaceAll(nextWord, "A o", "An o")

	return nextWord
}
