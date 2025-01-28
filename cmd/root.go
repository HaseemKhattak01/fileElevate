package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "github.com/HaseemKhattak01/mydriveuploader",
	Short: "MyDriveUploader is a CLI tool for uploading files to Google Drive",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(dropboxCmd)
	rootCmd.AddCommand(googleDriveCmd)
	rootCmd.AddCommand(serverCmd)
}
