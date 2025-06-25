package cmd

import "github.com/spf13/cobra"

var rootCmd = cobra.Command{
	Use: "LambdaPods",
}

func initial() {
	rootCmd.PersistentFlags().StringP("path", "p", "", "[Optional] lambda path [Default] .")
}

func MustExecute() {
	initial()

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
