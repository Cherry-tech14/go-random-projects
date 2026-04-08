package main

import (
	"fmt"
	"strconv"
	"strings"
)

func convertBases(hex string) string {
	word := strings.Fields(hex)
	for i := 1; i < len(word); i++ {
		switch word[i] {
		case "(hex)":
			value, err := strconv.ParseInt(word[i-1], 16, 64)
			if err != nil {
				return "invalid"
			}
			word[i-1] = strconv.FormatInt(value, 10)
			word = append(word[:i], word[i+1:]...)

		case "(bin)":
			if word[i] == "(bin)" {
				value, err := strconv.ParseInt(word[i-1], 2, 64)
				if err != nil {
					return "invalid"
				}
				word[i-1] = strconv.FormatInt(value, 10)

				word = append(word[:i], word[i+1:]...)
			}

		}
	}
	return strings.Join(word, " ")

}

func main() {
	fmt.Println(convertBases("FF (hex) files were added, and 10 (bin) years"))

}
