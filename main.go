package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("incomplete command")
	}
	input := os.Args[1]
	output := os.Args[2]
	data, err := os.ReadFile(input)
	if err != nil {
		fmt.Println("error reading file", err)
	}

	testword := convertBases(string(data))
	result := changeCase(testword)
	result = fixPunctuation(result)
	result = fixQuotes(result)
	result = vowelA2An(result)
	result = upperCase(result)

	os.Create("result.txt")

	finalOutput := os.WriteFile(output, []byte(result), 0644)

	if finalOutput != nil {
		fmt.Println("result not found", finalOutput)
	}
}
