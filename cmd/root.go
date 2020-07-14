package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "drive",
		Short: "A CLI for Google Drive",
		Long: `The Drive utility allows you to interact with Google Drive's 
API directly from your terminal. List files, upload, delete, transfer 
them using this CLI.`,
	}
)

// Execute executes the root command.
func Execute() error {
	// Do stuff that initiates other stuff here

	return rootCmd.Execute()
}
