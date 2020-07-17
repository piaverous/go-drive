package cmd

import (
	"fmt"

	"github.com/piaverous/go-drive/drive"
	"github.com/piaverous/go-drive/utils"

	"github.com/spf13/cobra"
)

func About(cmd *cobra.Command, args []string) {
	svc := drive.GetService()

	result := drive.About(svc)

	usage := utils.ByteCountBinary(result.StorageQuota.Usage)
	limit := utils.ByteCountBinary(result.StorageQuota.Limit)
	free := utils.ByteCountBinary(result.StorageQuota.Limit - result.StorageQuota.Usage)

	fmt.Printf("Current user is %s\n", utils.Magenta(result.User.DisplayName))
	fmt.Printf("---------------\n")
	fmt.Printf("Drive usage: %s / %s\n", utils.Purple(usage), utils.Red(limit))
	fmt.Printf("Available space : %s\n", utils.Purple(free))
}

var aboutCmd = &cobra.Command{
	Use:   "about",
	Short: "Get Metadata info about Drive session",
	Run:   About,
}

func init() {
	rootCmd.AddCommand(aboutCmd)
}
