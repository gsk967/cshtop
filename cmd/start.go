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

			rpcNodeUri, err := cmd.Flags().GetString(FlagRPCNode)
			if err != nil {
				return err
			}

			restUri, err := cmd.Flags().GetString(FlagRest)
			if err != nil {
				return err
			}

			return src.StartBlockExplorer(appName, rpcNodeUri, restUri)
		},
	}

	cmd.Flags().StringP(FlagRPCNode, "r", "", "rpc node uri, override chain registry rpc uris")
	cmd.Flags().StringP(FlagRest, "l", "", "rest uri, override chain registry rest uris")
	cmd.Flags().StringP(FlagAppName, "a", "umee", "Application Name")

	return cmd
}
