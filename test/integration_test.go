package test

import (
	"testing"

	"github.com/Atotti/mozisu-mcp-server/pkg/charcount"
)

// TestIntegration tests the integration between different components
func TestIntegration(t *testing.T) {
	// テストケース
	testCases := []struct {
		name                  string
		input                 string
		expectedTotal         int
		expectedNonWhitespace int
	}{
		{
			name:                  "Integration test - ASCII",
			input:                 "Hello, World!",
			expectedTotal:         13,
			expectedNonWhitespace: 12,
		},
		{
			name:                  "Integration test - Japanese",
			input:                 "こんにちは世界！",
			expectedTotal:         8,
			expectedNonWhitespace: 8,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 共通パッケージを使用して文字数をカウント
			result := charcount.Count(tc.input)

			// 結果を検証
			if result.TotalCount != tc.expectedTotal {
				t.Errorf("Expected total count %d, got %d", tc.expectedTotal, result.TotalCount)
			}
			if result.NonWhitespaceCount != tc.expectedNonWhitespace {
				t.Errorf("Expected non-whitespace count %d, got %d", tc.expectedNonWhitespace, result.NonWhitespaceCount)
			}
		})
	}
}
