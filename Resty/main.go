package main

import (
	"crypto/tls"
	"encoding/json"
	"log"

	"gopkg.in/resty.v1"
)

func httpCall(url string) string {
	// GET request
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		Get(url)

	// explore response object
	log.Printf("Error: %v", err)
	log.Printf("Response Status Code: %v", resp.StatusCode())
	log.Printf("Response Status: %v", resp.Status())
	log.Printf("Response Time: %v", resp.Time())
	log.Printf("Response Received At: %v", resp.ReceivedAt())
	log.Printf("Response Body: %v", resp) // or resp.String() or string(resp.Body())

	return string(resp.Body())
}

func toJson(text string) map[string]interface{} {
	var result map[string]interface{}
	json.Unmarshal([]byte(text), &result)
	return result
}

func main() {
	resty.SetDebug(true)
	resty.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	response := httpCall("http://httpbin.org/get")
	json := toJson(response)
	log.Println("Origin: ", json["origin"])

}
