package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// resetPasswordCmd represents the reset-password command
var resetPasswordCmd = &cobra.Command{
	Use:   "reset-password",
	Short: "Reset a user password",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("reset-password called")
	},
}

func init() {
	rootCmd.AddCommand(resetPasswordCmd)
}
