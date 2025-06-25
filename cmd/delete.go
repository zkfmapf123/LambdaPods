package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var DeleteCmd = cobra.Command{
	Use:   "del",
	Short: "Delete Lambda Functions",
	Long:  "Delete Lambda Functions",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete")
	},
}

func init() {
	rootCmd.AddCommand(&DeleteCmd)
}
