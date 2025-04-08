package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	// Parse command-line flags
	interactive := flag.Bool("i", false, "Run in interactive mode")
	flag.Parse()

	// Get the text to count
	var text string

	if *interactive {
		// Interactive mode
		fmt.Println("Character Count Tool")
		fmt.Println("===================")
		fmt.Println("Enter text to count characters (Ctrl+D to exit):")

		scanner := bufio.NewScanner(os.Stdin)
		for {
			fmt.Print("> ")
			if !scanner.Scan() {
				break
			}

			text = scanner.Text()
			printCharacterCounts(text)
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
			os.Exit(1)
		}
	} else {
		// Command-line mode
		args := flag.Args()
		if len(args) == 0 {
			fmt.Println("Please provide text to count characters.")
			fmt.Println("Usage: charcount [text] or charcount -i for interactive mode")
			os.Exit(1)
		}

		text = strings.Join(args, " ")
		printCharacterCounts(text)
	}
}

func printCharacterCounts(text string) {
	fmt.Println("DEBUG: Starting character count")

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

	// Write to stderr as well for debugging
	os.Stderr.WriteString(fmt.Sprintf("DEBUG: Completed character count for '%s'\n", text))
}
