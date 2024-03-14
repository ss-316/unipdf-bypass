// package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// )

// // Response struct represents the structure of the JSON response
// type Response struct {
// 	Code     string   `json:"Code"`
// 	Stdout   string   `json:"Stdout"`
// 	Stderr   string   `json:"Stderr"`
// 	Modified []string `json:"Modified"`
// }

// func main() {
// 	// Define the URL of unidoc playground - bypass
// 	url := "https://play.unidoc.io/api/run"

// 	// Define the headers
// 	headers := map[string]string{
// 		"authority":             "play.unidoc.io",
// 		"accept":                "*/*",
// 		"accept-language":       "en-GB,en-US;q=0.9,en;q=0.8",
// 		"content-type":          "application/json",
// 		"cookie":                "unidoc_playground_session_cookie=d4e4f704d0049c9048da",
// 		"origin":                "https://play.unidoc.io",
// 		"pass-next-get":         "on",
// 		"referer":               "https://play.unidoc.io/",
// 		"sec-ch-ua":             `"Google Chrome";v="119", "Chromium";v="119", "Not?A_Brand";v="24"`,
// 		"sec-ch-ua-mobile":      "?0",
// 		"sec-ch-ua-platform":    `"Windows"`,
// 		"sec-fetch-dest":        "empty",
// 		"sec-fetch-mode":        "cors",
// 		"sec-fetch-site":        "same-origin",
// 		"user-agent":            "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36",
// 		"x-kl-kes-ajax-request": "Ajax_Request",
// 	}

// 	// Read post payload data from the file
// 	payloadJSON, err := ioutil.ReadFile("payload.json")
// 	if err != nil {
// 		fmt.Println("Error reading payload file:", err)
// 		return
// 	}

// 	// Create a new HTTP request with the payload data
// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadJSON))
// 	if err != nil {
// 		fmt.Println("Error creating HTTP request:", err)
// 		return
// 	}

// 	// Set headers
// 	// req.Header.Set("Content-Type", "application/json")
// 	// req.Header.Set("authority", "play.unidoc.io")
// 	// req.Header.Set("accept", "*/*")
// 	// req.Header.Set("accept-language", "en-GB,en-US;q=0.9,en;q=0.8")
// 	// req.Header.Set("cookie", "unidoc_playground_session_cookie=d4e4f704d0049c9048da")
// 	// req.Header.Set("origin", "https://play.unidoc.io")
// 	// req.Header.Set("pass-next-get", "on")
// 	// req.Header.Set("referer", "https://play.unidoc.io/")
// 	// req.Header.Set("sec-ch-ua", "\"Google Chrome\";v=\"119\", \"Chromium\";v=\"119\", \"Not?A_Brand\";v=\"24\"")
// 	// req.Header.Set("sec-ch-ua-mobile", "?0")
// 	// req.Header.Set("sec-ch-ua-platform", "\"Windows\"")
// 	// req.Header.Set("sec-fetch-dest", "empty")
// 	// req.Header.Set("sec-fetch-mode", "cors")
// 	// req.Header.Set("sec-fetch-site", "same-origin")
// 	// req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36")
// 	// req.Header.Set("x-kl-kes-ajax-request", "Ajax_Request")

// 	for key, value := range headers {
// 		req.Header.Set(key, value)
// 	}

// 	// Create a new HTTP client
// 	client := &http.Client{}

// 	// Send the HTTP request
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println("Error sending HTTP request:", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	// Read the response body
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("Error reading response body:", err)
// 		return
// 	}

// 	// Parse the JSON response
// 	var response Response
// 	err = json.Unmarshal(body, &response)
// 	if err != nil {
// 		fmt.Println("Error parsing JSON response:", err)
// 		return
// 	}

// 	// Print Stdout as output text content
// 	fmt.Println("Output text content:")
// 	fmt.Println(response.Code)
// }

package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Payload struct represents the structure of the payload JSON
type Payload struct {
	Code  string `json:"Code"`
	Files []File `json:"Files"`
}

// File struct represents the structure of the file details in the payload JSON
type File struct {
	Filename string `json:"Filename"`
	B64str   string `json:"B64str"`
	MimeType string `json:"MimeType"`
	Size     int    `json:"Size"`
}

// Response struct represents the structure of the JSON response
type Response struct {
	Code     string   `json:"Code"`
	Stdout   string   `json:"Stdout"`
	Stderr   string   `json:"Stderr"`
	Modified []string `json:"Modified"`
}

func main() {

	pdfFileName := "check.pdf"
	withoutExtension := strings.Split(pdfFileName, ".pdf")[0]
	codeFileName := "code.old.txt"

	// Read the content of the local file
	codeContent, err := ioutil.ReadFile(codeFileName)
	if err != nil {
		fmt.Println("Error reading code file:", err)
		return
	}

	// Read the content of the PDF file
	pdfContent, err := ioutil.ReadFile(pdfFileName)
	if err != nil {
		fmt.Println("Error reading PDF file:", err)
		return
	}

	// Convert the PDF content to base64
	pdfBase64 := base64.StdEncoding.EncodeToString(pdfContent)

	// Create a new payload
	payload := Payload{
		Code: string(codeContent),
		Files: []File{
			{
				Filename: pdfFileName,
				B64str:   pdfBase64,
				MimeType: "application/pdf",
				Size:     len(pdfContent),
			},
		},
	}

	// Convert the payload to JSON
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error encoding payload to JSON:", err)
		return
	}

	// Define the URL of unidoc playground
	url := "https://play.unidoc.io/api/run"

	// Create a new HTTP request with the payload data
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadJSON))
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Create a new HTTP client
	client := &http.Client{}

	// Send the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("[-] Error sending HTTP request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("[-] Error reading response body:", err)
		return
	}

	// Parse the JSON response
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("[-] Error parsing JSON response:", err)
		return
	}

	// Print Stdout where the text is present
	fmt.Println("[+] Output text content:")
	fmt.Println(response.Stdout)

	// Save Stdout content to a file with the name of the PDF file in txt document
	err = ioutil.WriteFile(withoutExtension+".txt", []byte(response.Stdout), 0644)
	if err != nil {
		fmt.Println("[-] Error writing text content to file:", err)
		return
	}

	fmt.Println("----------------------------")
	fmt.Printf("%s\n", response.Stdout)
	fmt.Println("----------------------------")
	fmt.Printf("Text content saved to %s.txt\n", withoutExtension)
}
