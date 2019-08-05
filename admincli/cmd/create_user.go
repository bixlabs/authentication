package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// createUserCmd represents the create-user command
var createUserCmd = &cobra.Command{
	Use:   "create-user",
	Short: "Create a user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create-user called")
	},
}

func init() {
	rootCmd.AddCommand(createUserCmd)
}
