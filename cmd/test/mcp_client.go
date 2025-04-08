package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os/exec"
	"time"
)

// MCP protocol request
type MCPRequest struct {
	JsonRPC string      `json:"jsonrpc"`
	ID      int         `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}

// MCP protocol response
type MCPResponse struct {
	JsonRPC string          `json:"jsonrpc"`
	ID      int             `json:"id"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   *struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

func main() {
	fmt.Println("Starting MCP protocol test...")

	// Start the MCP server
	cmd := exec.Command("go", "run", "../../main.go")

	stdin, _ := cmd.StdinPipe()
	stdout, _ := cmd.StdoutPipe()

	cmd.Start()

	// Give the server a moment to start up
	time.Sleep(500 * time.Millisecond)

	// Create a reader for the stdout
	reader := bufio.NewReader(stdout)

	// First, try to list available tools
	listToolsRequest := MCPRequest{
		JsonRPC: "2.0",
		ID:      1,
		Method:  "listTools",
		Params:  struct{}{},
	}

	// Send the request
	requestJSON, _ := json.Marshal(listToolsRequest)
	fmt.Printf("Request: %s\n", requestJSON)
	io.WriteString(stdin, string(requestJSON)+"\n")

	// Read the response
	responseStr, _ := reader.ReadString('\n')
	fmt.Printf("Response: %s\n\n", responseStr)

	// Try to parse the response
	var response MCPResponse
	json.Unmarshal([]byte(responseStr), &response)

	// If we got an error, try some other common method names
	if response.Error != nil {
		methods := []string{
			"list_tools",
			"tools.list",
			"get_tools",
			"tools",
		}

		for i, method := range methods {
			listToolsRequest.ID = i + 2
			listToolsRequest.Method = method

			requestJSON, _ := json.Marshal(listToolsRequest)
			fmt.Printf("Trying method: %s\n", method)
			fmt.Printf("Request: %s\n", requestJSON)

			io.WriteString(stdin, string(requestJSON)+"\n")
			responseStr, _ := reader.ReadString('\n')
			fmt.Printf("Response: %s\n\n", responseStr)
		}
	}

	// Now try to call our count_characters tool directly
	callToolRequest := MCPRequest{
		JsonRPC: "2.0",
		ID:      10,
		Method:  "callTool",
		Params: map[string]interface{}{
			"name": "count_characters",
			"arguments": map[string]interface{}{
				"text": "こんにちは世界！Hello World! 😊🚀",
			},
		},
	}

	// Try different method names for calling a tool
	callMethods := []string{
		"callTool",
		"call_tool",
		"tool.call",
		"execute",
		"run",
	}

	for i, method := range callMethods {
		callToolRequest.ID = 10 + i
		callToolRequest.Method = method

		requestJSON, _ := json.Marshal(callToolRequest)
		fmt.Printf("Trying method: %s\n", method)
		fmt.Printf("Request: %s\n", requestJSON)

		io.WriteString(stdin, string(requestJSON)+"\n")
		responseStr, _ := reader.ReadString('\n')
		fmt.Printf("Response: %s\n\n", responseStr)
	}

	// Kill the server
	cmd.Process.Kill()
	fmt.Println("Test completed.")
}
