package dropbox

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/files"
)

// UploadFolderToDropbox uploads a local folder to Dropbox
func UploadFolderToDropbox(localFolderPath, dropboxFolderPath string) error {
	accessToken := os.Getenv("DROPBOX_ACCESS_TOKEN")
	if accessToken == "" {
		return fmt.Errorf("access token is not set")
	}

	dbxConfig := dropbox.Config{
		Token:    accessToken,
		LogLevel: dropbox.LogInfo,
	}
	dbx := files.New(dbxConfig)

	return filepath.Walk(localFolderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		if err := uploadFileToDropbox(dbx, path, dropboxFolderPath, info.Name()); err != nil {
			return fmt.Errorf("failed to upload file %s: %w", info.Name(), err)
		}
		log.Printf("Uploaded file to Dropbox: %s\n", info.Name())
		return nil
	})
}

func uploadFileToDropbox(dbx files.Client, localFilePath, dropboxFolderPath, fileName string) error {
	content, err := os.Open(localFilePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}
	defer content.Close()

	dropboxPath := filepath.ToSlash(filepath.Join(dropboxFolderPath, fileName))
	if dropboxPath == "" || dropboxPath == "/" {
		return fmt.Errorf("malformed Dropbox path: %s", dropboxPath)
	}

	commitInfo := files.NewUploadArg(dropboxPath)
	_, err = dbx.Upload(commitInfo, content)
	if err != nil {
		return fmt.Errorf("failed to upload file: %w", err)
	}
	return nil
}
