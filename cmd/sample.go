package cmd

import (
	"context"
	"driver/pkg/builder/response"
	"driver/pkg/client"
	"driver/pkg/config"
)

func main() {
	client := client.NewClient(&config.Config{
		OpenaiApiKey: "",
		TimeOut:      0,
		BaseUrl:      "",
	})
	client.Create(context.Background(), response.WithTextInput("Hello, World!"))
}
