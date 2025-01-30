package cmd

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/HaseemKhattak01/mydriveuploader/config"
	"github.com/HaseemKhattak01/mydriveuploader/drive"
	"github.com/HaseemKhattak01/mydriveuploader/dropbox"
	"github.com/HaseemKhattak01/mydriveuploader/utils"
	"github.com/spf13/cobra"
)

var (
	FolderPath  string
	DropboxPath string
)

var (
	googleDriveCmd = &cobra.Command{
		Use:   "uploadG",
		Short: "Upload files from local machine to Google Drive",
		RunE:  ExecuteDriveUpload,
	}

	dropboxCmd = &cobra.Command{
		Use:   "uploadD",
		Short: "Upload files from local machine to Dropbox",
		RunE:  ExecuteDropboxUpload,
	}

	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Start the web server for OAuth",
		Run:   StartWebServer,
	}
)

func init() {
	googleDriveCmd.Flags().StringVarP(&FolderPath, "folder", "f", "", "Path to the local folder to upload")
	googleDriveCmd.MarkFlagRequired("folder")
	rootCmd.AddCommand(googleDriveCmd)

	dropboxCmd.Flags().StringVarP(&FolderPath, "folder", "f", "", "Path to the local folder to upload")
	dropboxCmd.MarkFlagRequired("folder")
	dropboxCmd.Flags().StringVarP(&DropboxPath, "dropbox-path", "d", "", "Path in Dropbox to upload in")
	dropboxCmd.MarkFlagRequired("dropbox-path")
	rootCmd.AddCommand(dropboxCmd)
}

func StartWebServer(cmd *cobra.Command, args []string) {
	startURL := utils.GetOAuthStartURL()
	log.Printf("Starting OAuth process automatically with URL: %s", startURL)

	http.HandleFunc("/oauth/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		if code == "" {
			http.Error(w, "Code not found", http.StatusBadRequest)
			return
		}

		config.LoadConfig()
		cfg := config.GetConfig()
		token, err := utils.ExchangeCodeForToken(cfg.AppKey, cfg.AppSecret, code, cfg.RedirectURL)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to get token: %v", err), http.StatusInternalServerError)
			return
		}

		tokenFile := "DropBoxToken.json"
		if err := utils.SaveToken(tokenFile, token); err != nil {
			http.Error(w, fmt.Sprintf("Failed to save token: %v", err), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Access Token saved successfully")
		log.Println("Access Token saved successfully")
	})

	http.HandleFunc("/oauth/start", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, startURL, http.StatusTemporaryRedirect)
	})

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func ExecuteDriveUpload(cmd *cobra.Command, args []string) error {
	client, errResp := utils.GetDriveClient()
	if errResp != nil {
		log.Printf("Error getting HTTP client: %s", errResp.Error)
		return fmt.Errorf("unable to get HTTP client: %s", errResp.Error)
	}
	if FolderPath == "" {
		return errors.New("folder path is required")
	}
	if err := drive.UploadFolder(client, FolderPath); err != nil {
		log.Printf("Error uploading folder to Google Drive: %v", err)
		return fmt.Errorf("error uploading folder: %v", err)
	}
	log.Printf("Folder uploaded successfully to Google Drive from %s", FolderPath)
	return nil
}

func ExecuteDropboxUpload(cmd *cobra.Command, args []string) error {
	if FolderPath == "" || DropboxPath == "" {
		return errors.New("folder path and dropbox path are required")
	}
	if err := dropbox.UploadFolderToDropbox(FolderPath, DropboxPath); err != nil {
		log.Printf("Error uploading folder to Dropbox: %v", err)
		return fmt.Errorf("error uploading folder to Dropbox: %v", err)
	}
	log.Printf("Folder uploaded successfully to Dropbox from %s to %s", FolderPath, DropboxPath)
	return nil
}
