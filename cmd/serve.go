package cmd

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/HaseemKhattak01/mydriveuploader/drive"
	"github.com/HaseemKhattak01/mydriveuploader/dropbox"
	"github.com/HaseemKhattak01/mydriveuploader/utils"

	"github.com/spf13/cobra"
)

var (
	folderPath  string
	dropboxPath string
)

var (
	googleDriveCmd = &cobra.Command{
		Use:   "uploadG",
		Short: "Upload files from local machine to Google Drive",
		RunE:  executeDriveUpload,
	}

	dropboxCmd = &cobra.Command{
		Use:   "uploadD",
		Short: "Upload files from local machine to Dropbox",
		RunE:  executeDropboxUpload,
	}

	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Start the web server for OAuth",
		Run:   startWebServer,
	}
)

func init() {
	googleDriveCmd.Flags().StringVarP(&folderPath, "folder", "f", "", "Path to the local folder to upload")
	googleDriveCmd.MarkFlagRequired("folder")
	rootCmd.AddCommand(googleDriveCmd)

	dropboxCmd.Flags().StringVarP(&folderPath, "folder", "f", "", "Path to the local folder to upload")
	dropboxCmd.MarkFlagRequired("folder")
	dropboxCmd.Flags().StringVarP(&dropboxPath, "dropbox-path", "d", "", "Path in Dropbox to upload in")
	dropboxCmd.MarkFlagRequired("dropbox-path")
	rootCmd.AddCommand(dropboxCmd)
}

func startWebServer(cmd *cobra.Command, args []string) {
	startURL := utils.GetOAuthStartURL()
	fmt.Printf("Open the following URL in your browser to start the OAuth flow:\n%s\n", startURL)
	http.HandleFunc("/oauth/callback", utils.HandleOAuthCallback)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func executeDriveUpload(cmd *cobra.Command, args []string) error {
	client, errResp := utils.GetDriveClient()
	if errResp != nil {
		log.Printf("Error getting HTTP client: %s", errResp.Error)
		return fmt.Errorf("unable to get HTTP client: %s", errResp.Error)
	}
	if folderPath == "" {
		return errors.New("folder path is required")
	}
	if err := drive.UploadFolder(client, folderPath); err != nil {
		log.Printf("Error uploading folder to Google Drive: %v", err)
		return fmt.Errorf("error uploading folder: %v", err)
	}
	log.Printf("Folder uploaded successfully to Google Drive from %s", folderPath)
	return nil
}

func executeDropboxUpload(cmd *cobra.Command, args []string) error {
	if folderPath == "" || dropboxPath == "" {
		return errors.New("folder path and dropbox path are required")
	}
	if err := dropbox.UploadFolderToDropbox(folderPath, dropboxPath); err != nil {
		log.Printf("Error uploading folder to Dropbox: %v", err)
		return fmt.Errorf("error uploading folder to Dropbox: %v", err)
	}
	log.Printf("Folder uploaded successfully to Dropbox from %s to %s", folderPath, dropboxPath)
	return nil
}
