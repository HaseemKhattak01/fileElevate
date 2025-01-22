package main

import (
	"log"
	"mydriveuploader/drive"
	"mydriveuploader/utils"
)

func main() {
	// Initialize the Google Drive client using credentials
	client, err := utils.GetDriveClient("credential.json")
	if err != nil {
		log.Fatalf("Failed to create Drive client: %v", err)
		return
	}

	// Specify the path to the folder you want to upload
	folderPath := "D:/Downloads/Pinterest_files"

	if err := drive.UploadFolder(client, folderPath); err != nil {
		log.Fatalf("Failed to upload folder: %v", err)
	}
}
