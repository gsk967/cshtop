/*
Copyright © 2023 Sai Kumar <gsk967>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cshtop",
	Short: "Block explorer for cosmos-sdk based applications.",
	Long: `Real time block explorer for cosmos-sdk based applications
It will retrive node rpc uri and node api uri from cosmos chain registry
	`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.AddCommand(
		startCmd(),
	)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
