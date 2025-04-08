package main

import (
	"testing"
	"unicode"
)

// TestCountCharacters tests the character counting functionality
func TestCountCharacters(t *testing.T) {
	// テストケース
	testCases := []struct {
		name                  string
		input                 string
		expectedTotal         int
		expectedNonWhitespace int
	}{
		{
			name:                  "ASCII only",
			input:                 "Hello, World!",
			expectedTotal:         13,
			expectedNonWhitespace: 12,
		},
		{
			name:                  "Japanese characters",
			input:                 "こんにちは世界！",
			expectedTotal:         8,
			expectedNonWhitespace: 8,
		},
		{
			name:                  "Mixed with emojis",
			input:                 "Hello 世界 😊🚀",
			expectedTotal:         11,
			expectedNonWhitespace: 9,
		},
		{
			name:                  "Empty string",
			input:                 "",
			expectedTotal:         0,
			expectedNonWhitespace: 0,
		},
		{
			name:                  "Whitespace only",
			input:                 "   \t\n",
			expectedTotal:         5,
			expectedNonWhitespace: 0,
		},
	}

	// 各テストケースを実行
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 文字数をカウント
			totalCount := len([]rune(tc.input))

			// 空白以外の文字数をカウント
			nonWhitespaceCount := 0
			for _, r := range tc.input {
				if !unicode.IsSpace(r) {
					nonWhitespaceCount++
				}
			}

			// 結果を検証
			if totalCount != tc.expectedTotal {
				t.Errorf("Expected total count %d, got %d", tc.expectedTotal, totalCount)
			}
			if nonWhitespaceCount != tc.expectedNonWhitespace {
				t.Errorf("Expected non-whitespace count %d, got %d", tc.expectedNonWhitespace, nonWhitespaceCount)
			}
		})
	}
}
