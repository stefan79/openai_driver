package files

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	FilesCmd.AddCommand(UploadFileCmd)
	FilesCmd.AddCommand(ListFilesCmd)
	FilesCmd.AddCommand(RetrieveFileCmd)
	FilesCmd.AddCommand(DeleteFileCmd)
}

var FilesCmd = &cobra.Command{
	Use:   "files",
	Short: "Manage files",
	Long:  `Manage files`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use 'oai files --help' for more information")
	},
}
