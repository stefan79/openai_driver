package responses

import (
	"driver/pkg/builder/response"
	"driver/pkg/cli"
	"driver/pkg/client"
	"driver/pkg/config"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	openAIAPIKey string
	proxy        string
	baseUrl      string
	model        string
	output       string
	prompt       string
)

type createOptions struct {
	openAIAPIKey string
	proxy        *string
	baseUrl      string
	model        string
	output       string
	prompt       string
}

var CreateResponseCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a response",
	Long:  `Create a response`,
	Run: func(cmd *cobra.Command, args []string) {
		o := overlayCreateFlags(cmd, config.GetCfg(cmd))
		client, err := client.NewClient(o.openAIAPIKey, o.baseUrl, o.proxy)
		if err != nil {
			fmt.Printf("Error creating client: %v\n", err)
			return
		}
		input := response.WithTextInput("What is the capital of France?")
		resp, err := client.Create(cmd.Context(), o.model, input)
		if err != nil {
			fmt.Printf("Error creating response: %v\n", err)
			return
		}
		err = cli.DumpOutput(resp, o.output, os.Stdout)
		if err != nil {
			fmt.Printf("Error creating response: %v\n", err)
			return
		}
	},
}

func init() {
	CreateResponseCmd.Flags().StringVar(&openAIAPIKey, "openai-api-key", "", "OpenAI API key")
	CreateResponseCmd.Flags().StringVar(&proxy, "proxy", "", "Proxy")
	CreateResponseCmd.Flags().StringVar(&model, "model", "", "Model")
	CreateResponseCmd.Flags().StringVar(&baseUrl, "base-url", "", "Base URL")
	CreateResponseCmd.Flags().StringVar(&output, "output", "", "Output")
	CreateResponseCmd.Flags().StringVar(&prompt, "prompt", "", "Prompt")
}

func overlayCreateFlags(cmd *cobra.Command, base *config.Config) (createOptions, error) {
	o := createOptions{
		openAIAPIKey: base.OpenAI.APIKey,
		proxy:        base.HttpConfig.Proxy,
		baseUrl:      base.OpenAI.BaseUrl,
		model:        base.OpenAI.Model,
		output:       base.Console.Output,
	}
	if cmd.Flags().Changed("openai-api-key") {
		o.openAIAPIKey = openAIAPIKey
	}
	if cmd.Flags().Changed("proxy") {
		o.proxy = &proxy
	}
	if cmd.Flags().Changed("model") {
		o.model = model
	}
	if cmd.Flags().Changed("base-url") {
		o.baseUrl = baseUrl
	}
	if cmd.Flags().Changed("output") {
		o.output = output
	}
	if cmd.Flags().Changed("prompt") {
		o.prompt = prompt
	} else {
		return o, fmt.Errorf("prompt is required")
	}

	return o, nil
}
