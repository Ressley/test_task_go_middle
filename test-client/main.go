package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	httpServerAddr = "http://http-grpc-client:8080"
)

func main() {
	log.Println("[info] Sending challenge")
	if err := postRequest("/challenge", map[string]string{
		"destination": "grpc-server:50051",
	}); err != nil {
		log.Printf("[error] Failed to send challenge: %v", err)
	}

	log.Println("[info] Sending message")
	data := []byte("Hello World!")
	if err := postRequest("/message", map[string]interface{}{
		"destination": "grpc-server:50051",
		"data":        data,
	}); err != nil {
		log.Printf("[error] Failed to send message: %v", err)
	}
}

func postRequest(endpoint string, payload interface{}) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(httpServerAddr+endpoint, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		log.Printf("[error] Failed while sending request: %s", body)
		return errors.New(fmt.Sprintf("Failed while sending request: %s", body))
	}

	log.Printf("[info] Response: %s\n", body)
	return nil
}
