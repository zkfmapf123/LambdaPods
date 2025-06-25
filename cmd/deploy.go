package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var DeployCmd = cobra.Command{
	Use:   "dep",
	Short: "Deploy Lambda Functions",
	Long:  "Create / Deploy Lambda Functions",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deploy")
	},
}

func init() {
	rootCmd.AddCommand(&DeployCmd)
}
