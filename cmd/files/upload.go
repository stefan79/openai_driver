package files

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/stefan79/openai-driver/pkg/builder"
	"github.com/stefan79/openai-driver/pkg/cli"
	"github.com/stefan79/openai-driver/pkg/config"

	"github.com/spf13/cobra"
)

var (
	uploadOpenAIAPIKey string
	uploadProxy        string
	uploadBaseURL      string
	uploadPurpose      string
	uploadFilePath     string
)

var UploadFileCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload a file",
	Long:  `Upload a file to OpenAI`,
	RunE: func(cmd *cobra.Command, args []string) error {
		o, err := overlayUploadFlags(cmd, config.GetCfg(cmd))
		if err != nil {
			return fmt.Errorf("Error overlaying flags: %v\n", err)
		}
		client, err := cli.NewFileClient(o.OpenAIAPIKey, o.BaseUrl, o.Proxy)
		if err != nil {
			return fmt.Errorf("Error creating client: %v\n", err)
		}
		if err := cli.FilesUploadCommand(cmd.Context(), client, builder.NewFilesRegistry(), &o); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	UploadFileCmd.Flags().StringVar(&uploadOpenAIAPIKey, "openai-api-key", "", "OpenAI API key")
	UploadFileCmd.Flags().StringVar(&uploadProxy, "proxy", "", "Proxy")
	UploadFileCmd.Flags().StringVar(&uploadBaseURL, "base-url", "", "Base URL")
	UploadFileCmd.Flags().StringVar(&uploadPurpose, "purpose", "", "Purpose")
	UploadFileCmd.Flags().StringVar(&uploadFilePath, "file", "", "File path")
}

func overlayUploadFlags(cmd *cobra.Command, base *config.Config) (cli.FilesUploadOptions, error) {
	o := cli.FilesUploadOptions{
		OpenAIAPIKey: base.OpenAI.APIKey,
		Proxy:        base.HttpConfig.Proxy,
		BaseUrl:      base.OpenAI.BaseUrl,
	}
	if cmd.Flags().Changed("openai-api-key") {
		o.OpenAIAPIKey = uploadOpenAIAPIKey
	}
	if cmd.Flags().Changed("proxy") {
		o.Proxy = &uploadProxy
	}
	if cmd.Flags().Changed("base-url") {
		o.BaseUrl = uploadBaseURL
	}
	if cmd.Flags().Changed("purpose") {
		o.Purpose = uploadPurpose
	} else {
		return o, fmt.Errorf("purpose is required")
	}
	if cmd.Flags().Changed("file") {
		data, err := os.ReadFile(uploadFilePath)
		if err != nil {
			return o, err
		}
		o.FileName = filepath.Base(uploadFilePath)
		o.FileData = data
	} else {
		return o, fmt.Errorf("file is required")
	}
	return o, nil
}
