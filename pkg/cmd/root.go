package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use: "taskapp",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(agendasCmd)
	rootCmd.AddCommand(apiCmd)
	rootCmd.AddCommand(seedCmd)
}