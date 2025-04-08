# Mozisu MCP Server

[![CI](https://github.com/Atotti/mozisu-mcp-server/actions/workflows/ci.yml/badge.svg)](https://github.com/Atotti/mozisu-mcp-server/actions/workflows/ci.yml)

文字数をカウントして返すことで、LLMが正確な文字数で文章を作成できるようにするMCPサーバーです。
![Demo on Claude Desktop](image.png)
## 機能

- **文字数カウント**: テキストの文字数をカウントして返します
  - 全文字数（スペースを含む）
  - 空白以外の文字数
- **マルチバイト文字対応**: 日本語や絵文字などのUnicode文字を正確にカウント
- **複数の利用方法**: MCPサーバー、コマンドラインツール、Webインターフェースから利用可能

## インストール

```bash
# リポジトリのクローン
git clone https://github.com/Atotti/mozisu-mcp-server.git
cd mozisu-mcp-server

# 依存関係のインストール
go mod download

# ビルド
task build
```

## Setup for Claude Desktop
```json
{
	"mcpServers": {
	  "mozisu-mcp-server": {
		"command": "/Users/ayu/mycode/mozisu-mcp-server/bin/mozisu-mcp-server", // ビルド済みファイル
		"args": []
	  }
	}
}
```

## 使用方法

### MCPサーバーとして使用

```bash
go run cmd/mcpserver/main.go
```

または、ビルド済みのバイナリを使用:

```bash
./bin/mozisu-mcp-server
```

これにより、LLMが`count_characters`ツールを使用して文字数カウント機能を利用できます。

### コマンドラインツールとして使用

```bash
# ビルド済みのバイナリを使用
./bin/charcount "カウントしたいテキスト"

# または直接実行
go run cmd/charcount/main.go "カウントしたいテキスト"
```

対話モードで使用:

```bash
./bin/charcount -i
```

### Webインターフェースとして使用

```bash
# ビルド済みのバイナリを使用
./bin/webserver

# または直接実行
go run cmd/webserver/main.go
```

その後、ブラウザで http://localhost:8080 にアクセスします。

[MIT License](LICENSE)
