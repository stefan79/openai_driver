package files

import (
	"fmt"

	"github.com/stefan79/openai-driver/pkg/cli"
	"github.com/stefan79/openai-driver/pkg/config"

	"github.com/spf13/cobra"
)

var (
	retrieveOpenAIAPIKey string
	retrieveProxy        string
	retrieveBaseURL      string
	retrieveFileID       string
)

var RetrieveFileCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieve file metadata",
	Long:  `Retrieve metadata for a file`,
	RunE: func(cmd *cobra.Command, args []string) error {
		o, err := overlayRetrieveFlags(cmd, config.GetCfg(cmd))
		if err != nil {
			return fmt.Errorf("Error overlaying flags: %v\n", err)
		}
		client, err := cli.NewFileClient(o.OpenAIAPIKey, o.BaseUrl, o.Proxy)
		if err != nil {
			return fmt.Errorf("Error creating client: %v\n", err)
		}
		if err := cli.FilesRetrieveCommand(cmd.Context(), client, &o); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	RetrieveFileCmd.Flags().StringVar(&retrieveOpenAIAPIKey, "openai-api-key", "", "OpenAI API key")
	RetrieveFileCmd.Flags().StringVar(&retrieveProxy, "proxy", "", "Proxy")
	RetrieveFileCmd.Flags().StringVar(&retrieveBaseURL, "base-url", "", "Base URL")
	RetrieveFileCmd.Flags().StringVar(&retrieveFileID, "file-id", "", "File identifier")
}

func overlayRetrieveFlags(cmd *cobra.Command, base *config.Config) (cli.FilesRetrieveOptions, error) {
	o := cli.FilesRetrieveOptions{
		OpenAIAPIKey: base.OpenAI.APIKey,
		Proxy:        base.HttpConfig.Proxy,
		BaseUrl:      base.OpenAI.BaseUrl,
	}
	if cmd.Flags().Changed("openai-api-key") {
		o.OpenAIAPIKey = retrieveOpenAIAPIKey
	}
	if cmd.Flags().Changed("proxy") {
		o.Proxy = &retrieveProxy
	}
	if cmd.Flags().Changed("base-url") {
		o.BaseUrl = retrieveBaseURL
	}
	if cmd.Flags().Changed("file-id") {
		o.FileID = retrieveFileID
	} else {
		return o, fmt.Errorf("file-id is required")
	}
	return o, nil
}
