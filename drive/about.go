package drive

import (
	"log"

	"google.golang.org/api/drive/v3"
)

func About(service *drive.Service) *drive.About {
	result, err := service.About.
		Get().
		Fields("storageQuota, user").
		Do()

	if err != nil {
		log.Fatalf("Unable to retrieve files: %v", err)
	}
	return result
}
