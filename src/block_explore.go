package src

import (
	"os"

	"github.com/gsk967/cshtop/src/client"
	"github.com/gsk967/cshtop/src/utils"
	"github.com/tendermint/tendermint/libs/log"
)

// StartBlockExplorer
func StartBlockExplorer(appName, rpcNodeUri, restUri string) error {
	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))
	crr, err := client.GetChainRegistry(logger, appName)
	if err != nil {
		return err
	}
	var restUris []string
	if len(restUri) == 0 {
		restUris = utils.GetUris(crr.GetRestUris(), "443")
	} else {
		restUris = []string{restUri}
	}
	// Get the validators
	vals := client.GetMapValidators(client.GetValidators(logger, restUris))

	var rpcUris []string
	if len(rpcNodeUri) == 0 {
		rpcUris = utils.GetUris(crr.GetRPCUris(), "443")
	} else {
		rpcUris = []string{rpcNodeUri}
	}
	// Get the rpc client
	client := client.TMClient(logger, rpcUris)

	// draw main screen
	DrawMainMenu(logger, crr.GetChainName(), crr.PrettyName, crr.GetChainID(), client, vals)
	return nil
}
