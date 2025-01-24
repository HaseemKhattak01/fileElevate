package cmd

import (
	"errors"
	"fmt"
	"mydriveuploader/drive"
	"mydriveuploader/dropbox"
	"mydriveuploader/utils"

	"github.com/spf13/cobra"
)

var folderPath string
var dropboxPath string

var serveCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload files from local machine to Google Drive",
	RunE:  executeDriveUpload,
}

var dropboxCmd = &cobra.Command{
	Use:   "uploadD",
	Short: "Upload files from local machine to DropBox",
	RunE:  executeDropboxUpload,
}

func init() {
	serveCmd.Flags().StringVarP(&folderPath, "folder", "f", "", "Path to the local folder to upload")
	serveCmd.MarkFlagRequired("folder")
	rootCmd.AddCommand(serveCmd)

	dropboxCmd.Flags().StringVarP(&folderPath, "folder", "f", "", "path to the local folder to upload")
	dropboxCmd.MarkFlagRequired("folder")
	dropboxCmd.Flags().StringVarP(&dropboxPath, "dropbox-path", "d", "", "path in dropbox to upload in")
	dropboxCmd.MarkFlagRequired("dropbox-path")
	rootCmd.AddCommand(dropboxCmd)
}

func executeDriveUpload(cmd *cobra.Command, args []string) error {
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
}

func executeDropboxUpload(cmd *cobra.Command, args []string) error {
	if folderPath == "" || dropboxPath == "" {
		return errors.New("folder path and dropbox path are required")
	}
	err := dropbox.UploadFolderToDropbox(folderPath, dropboxPath)
	if err != nil {
		return fmt.Errorf("error uploading folder to Dropbox: %v", err)
	}
	return nil
}
