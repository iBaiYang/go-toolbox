package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{}

func init() {
	rootCmd.AddCommand(wordCmd)
	rootCmd.AddCommand(timeCmd)
	rootCmd.AddCommand(sqlCmd)
	rootCmd.AddCommand(jsonCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
