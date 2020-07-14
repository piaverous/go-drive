package drive

import (
	"context"
	"log"
	"sync"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
)

var once sync.Once
var instance *drive.Service

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

// GetService function uses the singleton pattern to prevent instantiating
// multiple instances of the Google Drive service. That way, we authenticate
// once for the session, not once per request
func GetService() *drive.Service {
	once.Do(func() { // <-- atomic, does not allow repeating
		instance = getDriveService() // <-- thread safe
	})
	return instance
}
