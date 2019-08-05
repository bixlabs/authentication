package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// findUserCmd represents the find-user command
var findUserCmd = &cobra.Command{
	Use:   "find-user",
	Short: "Find a user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("find-user called")
	},
}

func init() {
	rootCmd.AddCommand(findUserCmd)
}
