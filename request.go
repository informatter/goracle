package main

import (
	"fmt"
	"io"
	"net/http"
	"bytes"
	"encoding/json"
	"time"
	//"github.com/fatih/color"
)

type ResponseBody struct {
    Model              string    `json:"model"`
    CreatedAt          time.Time `json:"created_at"`
    Response           string    `json:"response"`
    Done               bool      `json:"done"`
    DoneReason         string    `json:"done_reason"`
    Context            []int     `json:"context"`
    TotalDuration      int64     `json:"total_duration"`
    LoadDuration       int64     `json:"load_duration"`
    PromptEvalDuration int64     `json:"prompt_eval_duration"`
    EvalCount          int       `json:"eval_count"`
    EvalDuration       int64     `json:"eval_duration"`
}

var responseChan = make(chan string,2)
const endpointUrl string =  "http://localhost:3000/api/generate"


// Sends an HTTP request to an LLM endpoint.
func request(model string, systemPrompt string, userPrompt string ) string{
	jsonData := fmt.Sprintf(`{"model":"%s", "system":"%s", "prompt":"%s", "stream":false}`, model, systemPrompt,userPrompt)

    // Create a new request using http
    req, err := http.NewRequest("POST", endpointUrl, bytes.NewBuffer([]byte(jsonData)))
    if err != nil {

        fmt.Println("Error creating request:", err)
        return ""
    }
    req.Header.Set("Content-Type", "application/json")


    // Send the request using http.Client
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error sending request:", err)
        return ""
    }
    defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
        fmt.Printf("Non-200 Response Status: %d\n", resp.StatusCode)
        fmt.Println("Response Body:", string(body))
    }
    if err != nil {
        fmt.Println("Error reading response body:", err)
        return ""
    }


	var responseBody ResponseBody
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		fmt.Println("Error unmarshaling response body:", err)
		return ""
	}

	return  fmt.Sprintf("%v\n", responseBody.Response)
}