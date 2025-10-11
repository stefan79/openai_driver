package responses

import (
	"driver/pkg/builder/responses"
	"driver/pkg/client"
	"driver/pkg/config"
	"driver/pkg/openai/responses/resp"
	"fmt"
	"strings"

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
		o, err := overlayCreateFlags(cmd, config.GetCfg(cmd))
		if err != nil {
			fmt.Printf("Error overlaying flags: %v\n", err)
			return
		}
		client, err := client.NewClient(o.openAIAPIKey, o.baseUrl, o.proxy)
		if err != nil {
			fmt.Printf("Error creating client: %v\n", err)
			return
		}
		input := responses.WithTextInput(o.prompt)
		reasoning := responses.WithReasoning(responses.EffortMinimal, responses.SummaryNone)
		r, err := client.Create(cmd.Context(), o.model, input, reasoning)
		if err != nil {
			fmt.Printf("Error creating response: %v\n", err)
			return
		}
		outputs, err := r.SelectOutput(resp.SelectMessage)
		if err != nil {
			fmt.Printf("Error selecting output: %v\n", err)
			return
		}
		for _, output := range outputs {
			fmt.Printf("%s\n", strings.Join(resp.SerializeText(&output), "\n"))
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
