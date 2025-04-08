package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

// JSON-RPC request structure
type Request struct {
	JsonRPC string      `json:"jsonrpc"`
	ID      int         `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}

// CallToolRequest parameters
type CallToolParams struct {
	Name      string                 `json:"name"`
	Arguments map[string]interface{} `json:"arguments"`
}

func RunTest() {
	// Start the MCP server as a subprocess
	cmd := exec.Command("go", "run", "../../main.go")

	// Set up pipes for stdin and stdout
	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Println("Error creating stdin pipe:", err)
		return
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error creating stdout pipe:", err)
		return
	}

	// Start the server
	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting server:", err)
		return
	}

	// Create a request to call the count_characters tool
	request := Request{
		JsonRPC: "2.0",
		ID:      1,
		Method:  "call_tool",
		Params: CallToolParams{
			Name: "count_characters",
			Arguments: map[string]interface{}{
				"text": "こんにちは世界！Hello World! 😊🚀",
			},
		},
	}

	// Marshal the request to JSON
	requestJSON, err := json.Marshal(request)
	if err != nil {
		fmt.Println("Error marshaling request:", err)
		return
	}

	// Add a newline to the request
	requestJSON = append(requestJSON, '\n')

	// Send the request to the server
	_, err = stdin.Write(requestJSON)
	if err != nil {
		fmt.Println("Error writing to stdin:", err)
		return
	}

	// Read the response
	buf := make([]byte, 4096)
	n, err := stdout.Read(buf)
	if err != nil {
		fmt.Println("Error reading from stdout:", err)
		return
	}

	// Print the response
	fmt.Println("Response from server:")
	fmt.Println(string(buf[:n]))

	// Test with another example
	request.ID = 2
	request.Method = "call_tool"
	request.Params = CallToolParams{
		Name: "count_characters",
		Arguments: map[string]interface{}{
			"text": "スペースを 含む 日本語 テキスト with English and 絵文字😊",
		},
	}

	// Marshal the request to JSON
	requestJSON, err = json.Marshal(request)
	if err != nil {
		fmt.Println("Error marshaling request:", err)
		return
	}

	// Add a newline to the request
	requestJSON = append(requestJSON, '\n')

	// Send the request to the server
	_, err = stdin.Write(requestJSON)
	if err != nil {
		fmt.Println("Error writing to stdin:", err)
		return
	}

	// Read the response
	n, err = stdout.Read(buf)
	if err != nil {
		fmt.Println("Error reading from stdout:", err)
		return
	}

	// Print the response
	fmt.Println("\nResponse from server (second test):")
	fmt.Println(string(buf[:n]))

	// Kill the server process
	cmd.Process.Kill()
}
