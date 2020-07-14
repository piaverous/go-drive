package drive

import (
	"log"

	"google.golang.org/api/drive/v3"
)

func ListFiles(service *drive.Service) *drive.FileList {
	result, err := service.Files.
		List().
		PageSize(10).
		Fields("nextPageToken, files(id, name)").
		Do()

	if err != nil {
		log.Fatalf("Unable to retrieve files: %v", err)
	}
	return result
}
