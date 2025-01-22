package main

import (
	"log"
	"mydriveuploader/drive"
	"mydriveuploader/utils"
)

func main() {
	
	client, err := utils.GetDriveClient("credential.json")
	if err != nil {
		log.Fatalf("Failed to create Drive client: %v", err)
		return
	}

	folderPath := "D:/Downloads/Pinterest_files"

	if err := drive.UploadFolder(client, folderPath); err != nil {
		log.Fatalf("Failed to upload folder: %v", err)
	}
}
