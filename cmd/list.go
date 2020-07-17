package cmd

import (
	"fmt"

	"github.com/piaverous/go-drive/drive"
	"github.com/piaverous/go-drive/utils"

	"github.com/spf13/cobra"
)

func List(cmd *cobra.Command, args []string) {
	svc := drive.GetService()
	r := drive.ListFiles(svc)

	if len(r.Files) == 0 {
		fmt.Println("No files found.")
	} else {
		fmt.Println("Files:")
		for _, i := range r.Files {
			fmt.Printf("\t- %s (%s) - %s\n", utils.Teal(i.Name), utils.Purple(i.Id), i.MimeType)
		}
	}
}

var listCmd = &cobra.Command{
	Use:   "ls",
	Short: "List files in your drive",
	Run:   List,
}

func init() {
	rootCmd.AddCommand(listCmd)
}
