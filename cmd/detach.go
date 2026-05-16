package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var detachCmd = &cobra.Command{
	Use:   "detach [fileName]",
	Short: "Detach from list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("detach called")
	},
}

func init() {
	rootCmd.AddCommand(detachCmd)
}
