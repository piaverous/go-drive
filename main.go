package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
)

func getDriveService() *drive.Service {
	ctx := context.Background()
	client, err := google.DefaultClient(
		ctx,
		drive.DriveScope,
	)

	svc, err := drive.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
	}
	return svc
}

func main() {
	if os.Getenv("GOOGLE_APPLICATION_CREDENTIALS") == "" {
		log.Fatal("Permission Denied - Please export GOOGLE_APPLICATION_CREDENTIALS env variable")
	}

	svc := getDriveService()
	r, err := svc.Files.
		List().
		PageSize(10).
		Fields("nextPageToken, files(id, name)").
		Do()

	if err != nil {
		log.Fatalf("Unable to retrieve files: %v", err)
	}

	fmt.Println("Files:")
	if len(r.Files) == 0 {
		fmt.Println("No files found.")
	} else {
		for _, i := range r.Files {
			fmt.Printf("\t- %s (%s)\n", i.Name, i.Id)
		}
	}

	result, err := svc.About.Get().Fields("storageQuota, user").Do()
	fmt.Printf("%v\n", result.User)
	fmt.Printf("%v\n", result.StorageQuota)
}
