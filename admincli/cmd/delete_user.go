package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteUserCmd represents the delete-user command
var deleteUserCmd = &cobra.Command{
	Use:   "delete-user",
	Short: "Delete a user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete-user called")
	},
}

func init() {
	rootCmd.AddCommand(deleteUserCmd)
}
