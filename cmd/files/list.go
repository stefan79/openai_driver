package files

import (
	"fmt"

	"github.com/stefan79/openai-driver/pkg/cli"
	"github.com/stefan79/openai-driver/pkg/config"
	"github.com/stefan79/openai-driver/pkg/output"

	"github.com/spf13/cobra"
)

var (
	listOpenAIAPIKey string
	listProxy        string
	listBaseURL      string
	listPurpose      string
)

var ListFilesCmd = &cobra.Command{
	Use:   "list",
	Short: "List files",
	Long:  `List uploaded files`,
	RunE: func(cmd *cobra.Command, args []string) error {
		o := overlayListFlags(cmd, config.GetCfg(cmd.Context()))
		oer := output.GetOutputter(cmd.Context())
		client, err := cli.NewFileClient(oer, o.OpenAIAPIKey, o.BaseUrl, o.Proxy)
		if err != nil {
			return fmt.Errorf("Error creating client: %v\n", err)
		}
		if err := cli.FilesListCommand(cmd.Context(), client, &o); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	ListFilesCmd.Flags().StringVar(&listOpenAIAPIKey, "openai-api-key", "", "OpenAI API key")
	ListFilesCmd.Flags().StringVar(&listProxy, "proxy", "", "Proxy")
	ListFilesCmd.Flags().StringVar(&listBaseURL, "base-url", "", "Base URL")
	ListFilesCmd.Flags().StringVar(&listPurpose, "purpose", "", "Purpose filter")
}

func overlayListFlags(cmd *cobra.Command, base *config.Config) cli.FilesListOptions {
	o := cli.FilesListOptions{
		OpenAIAPIKey: base.OpenAI.APIKey,
		Proxy:        base.HttpConfig.Proxy,
		BaseUrl:      base.OpenAI.BaseUrl,
	}
	if cmd.Flags().Changed("openai-api-key") {
		o.OpenAIAPIKey = listOpenAIAPIKey
	}
	if cmd.Flags().Changed("proxy") {
		o.Proxy = &listProxy
	}
	if cmd.Flags().Changed("base-url") {
		o.BaseUrl = listBaseURL
	}
	if cmd.Flags().Changed("purpose") {
		o.Purpose = &listPurpose
	}
	return o
}
