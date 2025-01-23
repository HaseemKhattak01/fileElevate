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
