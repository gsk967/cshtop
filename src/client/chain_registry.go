package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gsk967/cshtop/src/types"
	"github.com/tendermint/tendermint/libs/log"
)

// GetChainRegistry get the app information from cosmos chain registry
func GetChainRegistry(logger log.Logger, appName string) (types.ChainRegistryResponse, error) {
	logger = logger.With("ℹ getting the chain registry info")
	uri := fmt.Sprintf("https://raw.githubusercontent.com/cosmos/chain-registry/master/%s/chain.json", appName)
	resp, err := http.Get(uri)
	if err != nil {
		logger.Error("failed to get the chain registry info", "uri", uri, "err", err.Error())
		return types.ChainRegistryResponse{}, err
	}
	defer resp.Body.Close()
	r, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("failed to get the chain registry info", "uri", uri, "err", err.Error())
		return types.ChainRegistryResponse{}, err
	}
	var result types.ChainRegistryResponse
	err = json.Unmarshal(r, &result)
	if err != nil {
		logger.Error("failed to unmarshal the chain registry info", "uri", uri, "err", err.Error())
		return result, err
	}
	return result, err
}
