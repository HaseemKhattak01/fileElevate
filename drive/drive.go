package drive

import (
	"context"
	"fmt"
	"log"
	"mydriveuploader/utils"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

var folderPath string

var UploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload a folder to Google Drive",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := utils.GetDriveClient()
		if err != nil {
			return fmt.Errorf("unable to get HTTP client: %v", err)
		}
		if folderPath == "" {
			return fmt.Errorf("folder path is required")
		}
		return UploadFolder(client, folderPath)
	},
}

func init() {
	UploadCmd.Flags().StringVarP(&folderPath, "folder", "f", "", "Path to the folder to upload")
	UploadCmd.MarkFlagRequired("folder")
}

func UploadFolder(client *http.Client, folderPath string) error {
	srv, err := drive.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		return fmt.Errorf("unable to retrieve Drive client: %v", err)
	}

	return filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			f := &drive.File{Name: info.Name()}
			_, err = srv.Files.Create(f).Media(file).Do()
			if err != nil {
				return err
			}
			log.Printf("Uploaded file: %s\n", info.Name())
		}
		return nil
	})
}
