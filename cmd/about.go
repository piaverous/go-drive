package cmd

import (
	"fmt"

	"github.com/piaverous/go-drive/drive"

	"github.com/spf13/cobra"
)

func About(cmd *cobra.Command, args []string) {
	svc := drive.GetService()

	result := drive.About(svc)
	fmt.Printf("%v\n", result.User)
	fmt.Printf("%v\n", result.StorageQuota)
}

var aboutCmd = &cobra.Command{
	Use:   "about",
	Short: "Get Metadata info about Drive session",
	Run:   About,
}

func init() {
	rootCmd.AddCommand(aboutCmd)
}
