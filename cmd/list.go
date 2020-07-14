package cmd

import (
	"fmt"

	"github.com/piaverous/go-drive/drive"

	"github.com/spf13/cobra"
)

func List(cmd *cobra.Command, args []string) {
	svc := drive.GetService()
	r := drive.ListFiles(svc)

	fmt.Println("Files:")
	if len(r.Files) == 0 {
		fmt.Println("No files found.")
	} else {
		for _, i := range r.Files {
			fmt.Printf("\t- %s (%s)\n", i.Name, i.Id)
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
