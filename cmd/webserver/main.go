package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"unicode"
)

// CharacterCount represents the result of counting characters
type CharacterCount struct {
	Text               string `json:"text"`
	TotalCount         int    `json:"totalCount"`
	NonWhitespaceCount int    `json:"nonWhitespaceCount"`
}

func main() {
	// Define HTTP handlers
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/count", handleCount)

	// サーバー設定
	port := 8080
	addr := fmt.Sprintf(":%d", port)

	// タイムアウト設定付きのサーバーを作成
	server := &http.Server{
		Addr:         addr,
		Handler:      nil, // DefaultServeMux を使用
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// サーバーを非同期で起動
	go func() {
		fmt.Printf("Starting character count server on http://localhost:%d\n", port)
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
		log.Fatalf("Server shutdown error: %v\n", err)
	}

	fmt.Println("Server gracefully stopped")
}

// handleHome serves the home page with a simple form
func handleHome(w http.ResponseWriter, r *http.Request) {
	html := `
<!DOCTYPE html>
<html>
<head>
    <title>Character Count Tool</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            max-width: 800px;
            margin: 0 auto;
            padding: 20px;
        }
        textarea {
            width: 100%;
            height: 150px;
            margin-bottom: 10px;
        }
        button {
            padding: 8px 16px;
            background-color: #4CAF50;
            color: white;
            border: none;
            cursor: pointer;
        }
        #result {
            margin-top: 20px;
            padding: 10px;
            border: 1px solid #ddd;
            border-radius: 4px;
            display: none;
        }
    </style>
</head>
<body>
    <h1>Character Count Tool</h1>
    <p>Enter text below to count characters:</p>

    <textarea id="text" placeholder="Enter text here..."></textarea>
    <button onclick="countCharacters()">Count Characters</button>

    <div id="result"></div>

    <script>
        function countCharacters() {
            const text = document.getElementById('text').value;

            fetch('/count', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ text: text }),
            })
            .then(response => response.json())
            .then(data => {
                const resultDiv = document.getElementById('result');
                resultDiv.innerHTML =
                    '<h2>Results:</h2>' +
                    '<p><strong>Text:</strong> ' + data.text + '</p>' +
                    '<p><strong>Total characters:</strong> ' + data.totalCount + '</p>' +
                    '<p><strong>Non-whitespace characters:</strong> ' + data.nonWhitespaceCount + '</p>';
                resultDiv.style.display = 'block';
            })
            .catch(error => {
                console.error('Error:', error);
                alert('An error occurred while counting characters.');
            });
        }
    </script>
</body>
</html>
`
	w.Header().Set("Content-Type", "text/html")
	if _, err := w.Write([]byte(html)); err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

// handleCount processes the character count request
func handleCount(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the request body
	var request struct {
		Text string `json:"text"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Count characters
	text := request.Text
	totalCount := len([]rune(text))

	nonWhitespaceCount := 0
	for _, r := range text {
		if !unicode.IsSpace(r) {
			nonWhitespaceCount++
		}
	}

	// Create the response
	result := CharacterCount{
		Text:               text,
		TotalCount:         totalCount,
		NonWhitespaceCount: nonWhitespaceCount,
	}

	// Send the response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
	}
}
