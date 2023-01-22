package src

import (
	"github.com/gsk967/cshtop/src/client"
)

// StartBlockExplorer
func StartBlockExplorer(appName string) error {
	crr, err := client.GetChainRegistry(appName)
	if err != nil {
		return err
	}

	DrawMainMenu(crr.ChainName, crr.PrettyName, crr.ChainID, crr.Apis.RPC[0].Address)
	return nil
}
