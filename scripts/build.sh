#!/bin/bash

# スクリプトが存在するディレクトリの絶対パスを取得
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
# プロジェクトのルートディレクトリ（スクリプトの親ディレクトリ）
PROJECT_ROOT="$( cd "$SCRIPT_DIR/.." && pwd )"

# ビルド先ディレクトリ
BUILD_DIR="$PROJECT_ROOT/build"

# ビルド先ディレクトリが存在しない場合は作成
mkdir -p "$BUILD_DIR"

# 現在の日時を取得
TIMESTAMP=$(date +"%Y%m%d_%H%M%S")

echo "Building Mozisu MCP Server..."
echo "Project root: $PROJECT_ROOT"
echo "Build directory: $BUILD_DIR"
echo "Timestamp: $TIMESTAMP"

# MCPサーバーのビルド
echo "Building MCP server..."
go build -o "$BUILD_DIR/mozisu-mcp-server" "$PROJECT_ROOT/cmd/mcpserver"

# CLIツールのビルド
echo "Building CLI tool..."
go build -o "$BUILD_DIR/charcount" "$PROJECT_ROOT/cmd/charcount"

# Webサーバーのビルド
echo "Building Web server..."
go build -o "$BUILD_DIR/webserver" "$PROJECT_ROOT/cmd/webserver"

echo "Build completed successfully!"
echo "Binaries are available in: $BUILD_DIR"
