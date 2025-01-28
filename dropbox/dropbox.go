package dropbox

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/HaseemKhattak01/mydriveuploader/utils"
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/v6/dropbox/files"
)

func UploadFolderToDropbox(localFolderPath, dropboxFolderPath string) error {
	token, err := utils.TokenFromFile("DropBoxToken.json")
	if err != nil {
		return fmt.Errorf("failed to read access token: %w", err)
	}

	dbx := files.New(dropbox.Config{
		Token:    token.AccessToken,
		LogLevel: dropbox.LogInfo,
	})

	return filepath.Walk(localFolderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
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

	_, err = dbx.Upload(files.NewUploadArg(dropboxPath), content)
	if err != nil {
		return fmt.Errorf("failed to upload file: %w", err)
	}
	return nil
}
