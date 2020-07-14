package main

import (
	"log"
	"os"

	"github.com/piaverous/go-drive/cmd"
)

func main() {
	if os.Getenv("GOOGLE_APPLICATION_CREDENTIALS") == "" {
		log.Fatal("Permission Denied - Please export GOOGLE_APPLICATION_CREDENTIALS env variable")
    }
    
	cmd.Execute()
}
