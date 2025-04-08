package main

import (
	"fmt"
	"unicode"
)

func main() {
	// Test cases
	testCases := []string{
		"Hello, World!",
		"こんにちは世界！",
		"Hello 世界 😊🚀",
		"スペースを 含む 日本語 テキスト with English and 絵文字😊",
		"1234567890",
		"    Spaces at the beginning and end    ",
		"", // Empty string
	}

	fmt.Println("Character Count Test Results:")
	fmt.Println("============================")

	for _, text := range testCases {
		// Count total characters (including spaces)
		totalCount := len([]rune(text))

		// Count non-whitespace characters
		nonWhitespaceCount := 0
		for _, r := range text {
			if !unicode.IsSpace(r) {
				nonWhitespaceCount++
			}
		}

		fmt.Printf("\nText: %s\n", text)
		fmt.Printf("Total characters: %d\n", totalCount)
		fmt.Printf("Non-whitespace characters: %d\n", nonWhitespaceCount)
		fmt.Println("----------------------------")
	}

	fmt.Println("\nTest completed successfully!")
}
