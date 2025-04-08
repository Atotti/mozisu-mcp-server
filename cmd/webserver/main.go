package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Atotti/mozisu-mcp-server/internal/server"
	"github.com/Atotti/mozisu-mcp-server/pkg/charcount"
)

func main() {
	// Define HTTP handlers
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleHome)
	mux.HandleFunc("/count", handleCount)

	// サーバー設定
	config := server.DefaultConfig()
	config.Port = 8080

	// サーバーを起動
	log.Println("Starting character count web server...")
	if err := server.RunHTTPServer(mux, config); err != nil {
		log.Fatalf("Server error: %v\n", err)
	}
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

	// 共通パッケージを使用して文字数をカウント
	result := charcount.Count(request.Text)

	// Create the response
	response := struct {
		Text               string `json:"text"`
		TotalCount         int    `json:"totalCount"`
		NonWhitespaceCount int    `json:"nonWhitespaceCount"`
	}{
		Text:               result.Text,
		TotalCount:         result.TotalCount,
		NonWhitespaceCount: result.NonWhitespaceCount,
	}

	// Send the response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
	}
}
