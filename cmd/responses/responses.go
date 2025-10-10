package responses

import (
	"fmt"

	"github.com/spf13/cobra"
)

type CreateOptions struct {
	OpenAIAPIKey string `mapstructure:"openai_api_key"`
}

func init() {
	ResponsesCmd.AddCommand(CreateResponseCmd)
}

var ResponsesCmd = &cobra.Command{
	Use:   "responses",
	Short: "Manage responses",
	Long:  `Manage responses`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use 'oai responses --help' for more information")
	},
}
