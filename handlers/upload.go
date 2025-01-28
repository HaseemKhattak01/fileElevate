package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/HaseemKhattak01/mydriveuploader/dropbox"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		httpError(w, "Failed to parse form data", http.StatusBadRequest, err)
		return
	}

	localFolderPath, dropboxFolderPath := r.FormValue("localFolderPath"), r.FormValue("dropboxFolderPath")
	if localFolderPath == "" || dropboxFolderPath == "" {
		httpError(w, "Both localFolderPath and dropboxFolderPath are required", http.StatusBadRequest, nil)
		return
	}

	if err := dropbox.UploadFolderToDropbox(localFolderPath, dropboxFolderPath); err != nil {
		httpError(w, fmt.Sprintf("Failed to upload folder: %v", err), http.StatusInternalServerError, err)
		return
	}

	log.Printf("Folder uploaded successfully from %s to %s", localFolderPath, dropboxFolderPath)
	fmt.Fprintln(w, "Folder uploaded successfully")
}

func httpError(w http.ResponseWriter, message string, code int, err error) {
	if err != nil {
		log.Printf("%s: %v", message, err)
	} else {
		log.Println(message)
	}
	http.Error(w, message, code)
}
