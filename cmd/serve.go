package cmd

import (
	"errors"
	"fmt"
	"mydriveuploader/drive"
	"mydriveuploader/utils"

	"github.com/spf13/cobra"
)

var folderPath string

var serveCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload files from local machine to Google Drive",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, errResp := utils.GetDriveClient()
		if errResp != nil {
			return fmt.Errorf("unable to get HTTP client: %s", errResp.Error)
		}
		if folderPath == "" {
			return errors.New("folder path is required")
		}
		err := drive.UploadFolder(client, folderPath)
		if err != nil {
			return fmt.Errorf("error uploading folder: %v", err)
		}
		return nil
	},
}

func init() {
	serveCmd.Flags().StringVarP(&folderPath, "folder", "f", "", "Path to the local folder to upload")
	serveCmd.MarkFlagRequired("folder")
	rootCmd.AddCommand(serveCmd)
}
