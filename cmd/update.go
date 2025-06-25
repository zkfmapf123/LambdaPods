package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var UpdateCmd = cobra.Command{
	Use:   "u",
	Short: "Update Lambda Functions",
	Long:  "Update Lambda Functions",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update")
	},
}

func init() {
	rootCmd.AddCommand(&UpdateCmd)
}
