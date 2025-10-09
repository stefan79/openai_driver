package main

import (
	"context"
	"driver/pkg/builder/response"
	"driver/pkg/client"
	"driver/pkg/config"
	"fmt"
	"os"
)

func main() {
	key, ok := os.LookupEnv("OPENAI_API_KEY")
	if !ok {
		fmt.Println("OPENAI_API_KEY not set")
		return
	}

	proxy := "http://localhost:8888"

	client, err := client.NewClient(&config.Config{
		OpenaiApiKey: key,
		TimeOut:      0,
		BaseUrl:      "https://api.openai.com",
		Proxy:        &proxy,
	})
	if err != nil {
		fmt.Printf("Cannot start a new client, %e\n", err)
		return
	}
	response, err := client.Create(context.Background(), "gpt-4o", response.WithTextInput("Hello, World!"))
	if err != nil {
		fmt.Printf("Cannot process the request, %e\n", err)
		return
	}

	fmt.Printf("Response: %s\n", response.Id)

}
