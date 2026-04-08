package main

import (
	"fmt"
	"strings"
)

func upperCase(s string) string {
	word := strings.Fields(s)
	for i := 1; i < len(word); i++ {
		switch word[i] {
		case "(up)":
			word[i-1] = strings.ToUpper(word[i-1])
			word = append(word[:i], word[i+1:]...)

		case "(low)":
			word[i-1] = strings.ToLower(word[i-1])
			word = append(word[:i], word[i+1:]...)

		case "(cap)":
			word[i-1] = strings.ToLower(word[i-1])
			word[i-1] = strings.Title(word[i-1])
			word = append(word[:i], word[i+1:]...)

		}
	}
	return strings.Join(word, " ")

}

func main() {
	fmt.Println(upperCase("Ready set go (up) i should stop SHOUTING (low) Welcome to the Brooklyn BRIDGE (cap) "))
}
