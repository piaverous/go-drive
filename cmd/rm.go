package cmd

import (
	"fmt"

	"github.com/piaverous/go-drive/utils"

	"github.com/manifoldco/promptui"
	"github.com/piaverous/go-drive/drive"

	"github.com/spf13/cobra"
)

func Delete(cmd *cobra.Command, args []string) {
	svc := drive.GetService()
	r := drive.ListFiles(svc)

	if len(r.Files) == 0 {
		fmt.Println("No files found.")
	} else {
		filesMap := make(map[string]string)
		fileDisplayNames := make([]string, len(r.Files))
		for i := 0; i < len(r.Files); i++ {
			fileDisplayNames[i] = fmt.Sprintf("%15s - %5s - (%s)", utils.Teal(r.Files[i].Name), utils.ByteCountBinary(r.Files[i].Size), utils.Purple(r.Files[i].Id))
			filesMap[fileDisplayNames[i]] = r.Files[i].Id
		}

		prompt := promptui.Select{
			Label: "Select a file to delete",
			Items: fileDisplayNames,
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		drive.Delete(svc, filesMap[result])
	}
}

var deleteCmd = &cobra.Command{
	Use:   "rm",
	Short: "Deletes a file from your drive",
	Run:   Delete,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
