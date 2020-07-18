package drive

import (
	"log"

	"google.golang.org/api/drive/v3"
)

func Delete(service *drive.Service, fileId string) {
	err := service.Files.Delete(fileId).Do()

	if err != nil {
		log.Fatalf("Unable to delete file: %v", err)
	}
}
