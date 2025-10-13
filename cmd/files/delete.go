package files

import (
	"fmt"

	"github.com/stefan79/openai-driver/pkg/cli"
	"github.com/stefan79/openai-driver/pkg/config"
	"github.com/stefan79/openai-driver/pkg/output"

	"github.com/spf13/cobra"
)

var (
	deleteOpenAIAPIKey string
	deleteProxy        string
	deleteBaseURL      string
	deleteFileID       string
)

var DeleteFileCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a file",
	Long:  `Delete a file by id`,
	RunE: func(cmd *cobra.Command, args []string) error {
		o, err := overlayDeleteFlags(cmd, config.GetCfg(cmd.Context()))
		if err != nil {
			return fmt.Errorf("Error overlaying flags: %v\n", err)
		}
		oer := output.GetOutputter(cmd.Context())
		client, err := cli.NewFileClient(oer, o.OpenAIAPIKey, o.BaseUrl, o.Proxy)
		if err != nil {
			return fmt.Errorf("Error creating client: %v\n", err)
		}
		if err := cli.FilesDeleteCommand(cmd.Context(), client, &o); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	DeleteFileCmd.Flags().StringVar(&deleteOpenAIAPIKey, "openai-api-key", "", "OpenAI API key")
	DeleteFileCmd.Flags().StringVar(&deleteProxy, "proxy", "", "Proxy")
	DeleteFileCmd.Flags().StringVar(&deleteBaseURL, "base-url", "", "Base URL")
	DeleteFileCmd.Flags().StringVar(&deleteFileID, "file-id", "", "File identifier")
}

func overlayDeleteFlags(cmd *cobra.Command, base *config.Config) (cli.FilesDeleteOptions, error) {
	o := cli.FilesDeleteOptions{
		OpenAIAPIKey: base.OpenAI.APIKey,
		Proxy:        base.HttpConfig.Proxy,
		BaseUrl:      base.OpenAI.BaseUrl,
	}
	if cmd.Flags().Changed("openai-api-key") {
		o.OpenAIAPIKey = deleteOpenAIAPIKey
	}
	if cmd.Flags().Changed("proxy") {
		o.Proxy = &deleteProxy
	}
	if cmd.Flags().Changed("base-url") {
		o.BaseUrl = deleteBaseURL
	}
	if cmd.Flags().Changed("file-id") {
		o.FileID = deleteFileID
	} else {
		return o, fmt.Errorf("file-id is required")
	}
	return o, nil
}
