package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var InfoCmd = cobra.Command{
	Use:   "i",
	Short: "Info Lambda Functions",
	Long:  "Info Lambda Functions",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("info")
	},
}

func init() {
	rootCmd.AddCommand(&InfoCmd)
}
