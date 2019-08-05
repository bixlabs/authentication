package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// updateUserCmd represents the update-user command
var updateUserCmd = &cobra.Command{
	Use:   "update-user",
	Short: "Update a user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update-user called")
	},
}

func init() {
	rootCmd.AddCommand(updateUserCmd)
}
