package drive

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func UploadFolder(client *http.Client, folderPath string) error {
	srv, err := drive.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		return fmt.Errorf("unable to retrieve Drive client: %v", err)
	}

	return filepath.Walk(folderPath, func(path string, info os.FileInfo, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if info.IsDir() {
			return nil
		}

		if err := uploadFile(srv, path, info.Name()); err != nil {
			return fmt.Errorf("failed to upload file %s: %v", info.Name(), err)
		}
		log.Printf("Uploaded file: %s", info.Name())
		return nil
	})
}

func uploadFile(srv *drive.Service, filePath, fileName string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("unable to open file %s: %v", fileName, err)
	}
	defer file.Close()

	driveFile := &drive.File{Name: fileName}
	if _, err := srv.Files.Create(driveFile).Media(file).Do(); err != nil {
		return fmt.Errorf("unable to create file %s in Drive: %v", fileName, err)
	}
	return nil
}
