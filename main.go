package main

import (
	"context"
	"errors"
	"fmt"
	"unicode"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// 1. MCPサーバーを作成
	s := server.NewMCPServer(
		"Demo Server 🚀", // サーバー名
		"1.0.0",         // バージョン
	)

	// 2. ツールを定義
	helloTool := mcp.NewTool(
		"hello_world",
		mcp.WithDescription("Say hello to someone"),
		// 引数 name の設定
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Name of the person to greet"),
		),
	)

	// 文字数カウントツールを定義
	countCharsTool := mcp.NewTool(
		"count_characters",
		mcp.WithDescription("Count characters in text"),
		mcp.WithString("text",
			mcp.Required(),
			mcp.Description("Text to count characters in"),
		),
	)

	// 3. ツールハンドラを登録
	s.AddTool(helloTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		// 引数の取得
		name, ok := request.Params.Arguments["name"].(string)
		if !ok {
			return nil, errors.New("name must be a string")
		}

		// 結果を返す
		return mcp.NewToolResultText(fmt.Sprintf("Hello, %s!", name)), nil
	})

	// 文字数カウントツールのハンドラを登録
	s.AddTool(countCharsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		// テキスト引数の取得
		text, ok := request.Params.Arguments["text"].(string)
		if !ok {
			return nil, errors.New("text must be a string")
		}

		// 全文字数をカウント（マルチバイト文字対応）
		totalCount := len([]rune(text))

		// 空白以外の文字数をカウント（マルチバイト文字対応）
		nonWhitespaceCount := 0
		for _, r := range text {
			if !unicode.IsSpace(r) {
				nonWhitespaceCount++
			}
		}

		// レスポンスの作成
		response := fmt.Sprintf(
			"Text: %s\nTotal characters: %d\nNon-whitespace characters: %d",
			text, totalCount, nonWhitespaceCount,
		)

		return mcp.NewToolResultText(response), nil
	})

	// 4. サーバーを起動 (stdio)
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
