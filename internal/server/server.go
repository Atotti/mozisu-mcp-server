// Package server provides internal server utilities
package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// Config represents server configuration
type Config struct {
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

// DefaultConfig returns the default server configuration
func DefaultConfig() Config {
	return Config{
		Port:         8080,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
}

// RunHTTPServer starts an HTTP server with graceful shutdown
func RunHTTPServer(handler http.Handler, config Config) error {
	// サーバー設定
	addr := fmt.Sprintf(":%d", config.Port)

	// タイムアウト設定付きのサーバーを作成
	server := &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
		IdleTimeout:  config.IdleTimeout,
	}

	// サーバーを非同期で起動
	go func() {
		fmt.Printf("Starting server on http://localhost:%d\n", config.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v\n", err)
		}
	}()

	// シグナル処理（Ctrl+C など）
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// シグナルを待機
	<-stop

	// グレースフルシャットダウン
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println("\nShutting down server...")
	if err := server.Shutdown(ctx); err != nil {
		return fmt.Errorf("server shutdown error: %v", err)
	}

	fmt.Println("Server gracefully stopped")
	return nil
}
