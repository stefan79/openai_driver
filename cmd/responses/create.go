package responses

import (
	"driver/pkg/builder/responses"
	"driver/pkg/client"
	"driver/pkg/config"
	"driver/pkg/openai/responses/resp"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	openAIAPIKey         string
	proxy                string
	baseUrl              string
	model                string
	output               string
	prompt               string
	effort               string
	summary              string
	file                 string
	webSearch            bool
	webSearchContextSize string
)

type createOptions struct {
	openAIAPIKey         string
	proxy                *string
	baseUrl              string
	model                string
	output               string
	prompt               string
	effort               *responses.Effort
	summary              *responses.Summary
	fileName             *string
	fileData             []byte
	webSearch            bool
	webSearchContextSize *responses.WebSearchContextSize
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
		options := []responses.ResponseOption{}
		options = append(options, responses.WithTextInput(o.prompt))
		if o.effort != nil {
			options = append(options, responses.WithReasoningEffort(*o.effort))
		}
		if o.summary != nil {
			options = append(options, responses.WithReasoningSummary(*o.summary))
		}
		if o.fileName != nil {
			options = append(options, responses.WithFileInput(*o.fileName, o.fileData))
		}
		if o.webSearch {
			options = append(options, responses.WithWebSearch())
		}
		if o.webSearchContextSize != nil {
			options = append(options, responses.WithWebSearchContextSize(*o.webSearchContextSize))
		}
		r, err := client.Create(cmd.Context(), o.model, options...)
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
	CreateResponseCmd.Flags().StringVar(&effort, "effort", "", "Effort")
	CreateResponseCmd.Flags().StringVar(&summary, "summary", "", "Summary")
	CreateResponseCmd.Flags().StringVar(&file, "file", "", "File Name")
	CreateResponseCmd.Flags().BoolVar(&webSearch, "web-search", true, "Web Search")
	CreateResponseCmd.Flags().StringVar(&webSearchContextSize, "web-search-context-size", "", "Web Search Context Size")
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
	if cmd.Flags().Changed("web-search-context-size") {
		webSearchContextSize, err := responses.ParseReasoningSummaryWebSearchContextSize(webSearchContextSize)
		if err != nil {
			return o, err
		}
		o.webSearchContextSize = &webSearchContextSize
	}

	if cmd.Flags().Changed("effort") {
		effort, err := responses.ParseReasoningEffort(effort)
		if err != nil {
			return o, err
		}
		o.effort = &effort
	}
	if cmd.Flags().Changed("summary") {
		summary, err := responses.ParseReasoningSummary(summary)
		if err != nil {
			return o, err
		}
		o.summary = &summary
	}
	if cmd.Flags().Changed("file") {
		data, err := os.ReadFile(file)
		if err != nil {
			return o, err
		}
		o.fileName = &file
		o.fileData = data
	}
	if cmd.Flags().Changed("web-search") {
		o.webSearch = true
	}

	return o, nil
}
