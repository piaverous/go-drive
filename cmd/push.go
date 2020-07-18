package cmd

import (
	"errors"

	"github.com/piaverous/go-drive/drive"

	"github.com/spf13/cobra"
)

func Push(cmd *cobra.Command, args []string) {
	svc := drive.GetService()
	drive.Push(svc, args[0])
}

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push files to your drive",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("Requires exactly one argument: file to push")
		}
		return nil
	},
	Run: Push,
}

func init() {
	rootCmd.AddCommand(pushCmd)
}
