package cmd

import (
	"github.com/spf13/cobra"

	"github.com/gsk967/cshtop/src"
)

func startCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "start",
		Short:   "Start the block explorer",
		Example: `cshtop start --app akash`,
		RunE: func(cmd *cobra.Command, args []string) error {
			appName, err := cmd.Flags().GetString(FlagAppName)
			if err != nil {
				return err
			}
			if err := src.StartBlockExplorer(appName); err != nil {
				return err
			}
			return nil
		},
	}

	cmd.Flags().StringP(FlagNode, "n", "http://localhost:26657", "App node uri")
	cmd.Flags().StringP(FlagAppName, "a", "umee", "Application Name")

	return cmd
}
