package drive

import (
	"log"
	"os"
	"path/filepath"

	"google.golang.org/api/drive/v3"
)

func Push(service *drive.Service, file string) *drive.File {
	ids, err := service.Files.GenerateIds().Count(1).Do()
	if err != nil {
		log.Fatalf("Unable to generate ids: %v", err)
	}
	_, fileName := filepath.Split(file)

	reader, err := os.Open(file)
	defer reader.Close()

	result, err := service.Files.Create(&drive.File{
		Id:   ids.Ids[0],
		Name: fileName,
	}).Media(reader).Do()

	if err != nil {
		log.Fatalf("Unable to create file: %v", err)
	}
	return result
}
