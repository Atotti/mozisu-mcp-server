package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/Atotti/mozisu-mcp-server/pkg/charcount"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// 1. MCPサーバーを作成
	s := server.NewMCPServer(
		"Mozisu MCP Server", // サーバー名
		"1.0.0",             // バージョン
	)

	// 2. 文字数カウントツールを定義
	countCharsTool := mcp.NewTool(
		"count_characters",
		mcp.WithDescription("Count characters in text"),
		mcp.WithString("text",
			mcp.Required(),
			mcp.Description("Text to count characters in"),
		),
	)

	// 3. 文字数カウントツールのハンドラを登録
	s.AddTool(countCharsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		// テキスト引数の取得
		text, ok := request.Params.Arguments["text"].(string)
		if !ok {
			return nil, errors.New("text must be a string")
		}

		// 共通パッケージを使用して文字数をカウント
		result := charcount.Count(text)

		// レスポンスの作成
		response := fmt.Sprintf(
			"Text: %s\nTotal characters: %d\nNon-whitespace characters: %d",
			result.Text, result.TotalCount, result.NonWhitespaceCount,
		)

		return mcp.NewToolResultText(response), nil
	})

	// 4. サーバーを起動 (stdio)
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
