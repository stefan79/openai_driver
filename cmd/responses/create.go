package responses

import (
	"fmt"
	"os"

	"github.com/stefan79/openai-driver/pkg/builder"
	"github.com/stefan79/openai-driver/pkg/cli"
	"github.com/stefan79/openai-driver/pkg/config"
	outputter "github.com/stefan79/openai-driver/pkg/output"

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

var CreateResponseCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a response",
	Long:  `Create a response`,
	RunE: func(cmd *cobra.Command, args []string) error {
		o, err := overlayCreateFlags(cmd, config.GetCfg(cmd))
		outputter := outputter.NewDefaultOutputter()
		if err != nil {
			return fmt.Errorf("Error overlaying flags: %v\n", err)
		}
		client, err := cli.NewClient(outputter, &o)
		if err != nil {
			return fmt.Errorf("Error creating client: %v\n", err)
		}
		if err := cli.ResponsesCreateCommand(cmd.Context(), outputter, client, builder.NewRegistry(), &o); err != nil {
			return fmt.Errorf("Error creating response: %v\n", err)
		}
		return nil
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

func overlayCreateFlags(cmd *cobra.Command, base *config.Config) (cli.ResponsesCreateOptions, error) {
	o := cli.ResponsesCreateOptions{
		OpenAIAPIKey: base.OpenAI.APIKey,
		Proxy:        base.HttpConfig.Proxy,
		BaseUrl:      base.OpenAI.BaseUrl,
		Model:        base.OpenAI.Model,
		Output:       base.Console.Output,
	}
	if cmd.Flags().Changed("openai-api-key") {
		o.OpenAIAPIKey = openAIAPIKey
	}
	if cmd.Flags().Changed("proxy") {
		o.Proxy = &proxy
	}
	if cmd.Flags().Changed("model") {
		o.Model = model
	}
	if cmd.Flags().Changed("base-url") {
		o.BaseUrl = baseUrl
	}
	if cmd.Flags().Changed("output") {
		o.Output = output
	}
	if cmd.Flags().Changed("prompt") {
		o.Prompt = prompt
	} else {
		return o, fmt.Errorf("prompt is required")
	}
	if cmd.Flags().Changed("web-search-context-size") {
		webSearchContextSize, err := builder.ParseWebSearchContextSize(webSearchContextSize)
		if err != nil {
			return o, err
		}
		o.WebSearchContextSize = &webSearchContextSize
	}

	if cmd.Flags().Changed("effort") {
		effort, err := builder.ParseReasoningEffort(effort)
		if err != nil {
			return o, err
		}
		o.Effort = &effort
	}
	if cmd.Flags().Changed("summary") {
		summary, err := builder.ParseReasoningSummary(summary)
		if err != nil {
			return o, err
		}
		o.Summary = &summary
	}
	if cmd.Flags().Changed("file") {
		data, err := os.ReadFile(file)
		if err != nil {
			return o, err
		}
		o.FileName = &file
		o.FileData = data
	}
	if cmd.Flags().Changed("web-search") {
		o.WebSearch = true
	}

	return o, nil
}
