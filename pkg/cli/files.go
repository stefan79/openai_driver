package cli

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/stefan79/openai-driver/pkg/builder"
	"github.com/stefan79/openai-driver/pkg/client"
	openaifiles "github.com/stefan79/openai-driver/pkg/openai/files"
)

type FilesUploadOptions struct {
	OpenAIAPIKey string
	Proxy        *string
	BaseUrl      string
	Purpose      string
	FileName     string
	FileData     []byte
}

type FilesListOptions struct {
	OpenAIAPIKey string
	Proxy        *string
	BaseUrl      string
	Purpose      *string
}

type FilesRetrieveOptions struct {
	OpenAIAPIKey string
	Proxy        *string
	BaseUrl      string
	FileID       string
}

type FilesDeleteOptions struct {
	OpenAIAPIKey string
	Proxy        *string
	BaseUrl      string
	FileID       string
}

func NewFileClient(apiKey, baseUrl string, proxy *string) (client.Client, error) {
	return client.NewClient(apiKey, baseUrl, proxy)
}

func FilesUploadCommand(ctx context.Context, client client.Client, registry builder.FileRegistry, o *FilesUploadOptions) error {
	req, err := registry.Upload(o.Purpose, o.FileName, o.FileData)
	if err != nil {
		return err
	}
	file, err := client.UploadFile(ctx, req)
	if err != nil {
		return fmt.Errorf("Error uploading file: %v\n", err)
	}
	return printFile(file)
}

func FilesListCommand(ctx context.Context, client client.Client, o *FilesListOptions) error {
	files, err := client.ListFiles(ctx, o.Purpose)
	if err != nil {
		return fmt.Errorf("Error listing files: %v\n", err)
	}
	return printJSON(files)
}

func FilesRetrieveCommand(ctx context.Context, client client.Client, o *FilesRetrieveOptions) error {
	file, err := client.RetrieveFile(ctx, o.FileID)
	if err != nil {
		return fmt.Errorf("Error retrieving file: %v\n", err)
	}
	return printFile(file)
}

func FilesDeleteCommand(ctx context.Context, client client.Client, o *FilesDeleteOptions) error {
	result, err := client.DeleteFile(ctx, o.FileID)
	if err != nil {
		return fmt.Errorf("Error deleting file: %v\n", err)
	}
	return printJSON(result)
}

func printFile(file *openaifiles.File) error {
	return printJSON(file)
}

func printJSON(v interface{}) error {
	output, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", output)
	return nil
}
