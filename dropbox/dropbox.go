package dropbox

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"mydriveuploader/config"

	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/files"
)

func UploadFolderToDropbox(localFolderPath, dropboxFolderPath string) error {
	dbxConfig := dropbox.Config{
		Token:    config.GetDropBoxToken(),
		LogLevel: dropbox.LogInfo, // Change to dropbox.LogOff for no logging
	}
	dbx := files.New(dbxConfig)

	return filepath.Walk(localFolderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			content, err := os.Open(path)
			if err != nil {
				return fmt.Errorf("failed to read file: %w", err)
			}
			defer content.Close()

			// Ensure the Dropbox path uses forward slashes and is well-formed
			dropboxPath := filepath.ToSlash(filepath.Join(dropboxFolderPath, info.Name()))
			if dropboxPath == "" || dropboxPath == "/" {
				return fmt.Errorf("malformed Dropbox path: %s", dropboxPath)
			}
			commitInfo := files.NewUploadArg(dropboxPath)
			_, err = dbx.Upload(commitInfo, content)
			if err != nil {
				return fmt.Errorf("failed to upload file: %w", err)
			}
			log.Printf("Uploaded file to Dropbox: %s\n", info.Name())
		}
		return nil
	})
}
