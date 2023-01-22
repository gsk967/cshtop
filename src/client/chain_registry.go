package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ChainRegistryResponse struct {
	ChainName    string `json:"chain_name,omitempty"`
	Status       string `json:"status,omitempty"`
	NetworkType  string `json:"network_type,omitempty"`
	Website      string `json:"website,omitempty"`
	PrettyName   string `json:"pretty_name,omitempty"`
	ChainID      string `json:"chain_id,omitempty"`
	Bech32Prefix string `json:"bech32_prefix,omitempty"`
	DaemonName   string `json:"daemon_name,omitempty"`
	NodeHome     string `json:"node_home,omitempty"`
	Slip44       int    `json:"slip44,omitempty"`
	Apis         struct {
		RPC []struct {
			Address  string `json:"address,omitempty"`
			Provider string `json:"provider,omitempty"`
		} `json:"rpc,omitempty"`
		Rest []struct {
			Address  string `json:"address,omitempty"`
			Provider string `json:"provider,omitempty"`
		} `json:"rest,omitempty"`
		Grpc []struct {
			Address  string `json:"address,omitempty"`
			Provider string `json:"provider,omitempty"`
		} `json:"grpc,omitempty"`
	} `json:"apis,omitempty"`
	Explorers []struct {
		Kind        string `json:"kind,omitempty"`
		URL         string `json:"url,omitempty"`
		TxPage      string `json:"tx_page,omitempty"`
		AccountPage string `json:"account_page,omitempty"`
	} `json:"explorers,omitempty"`
}

// GetChainRegistry get the app information from cosmos chain registry
func GetChainRegistry(appName string) (ChainRegistryResponse, error) {
	uri := fmt.Sprintf("https://raw.githubusercontent.com/cosmos/chain-registry/master/%s/chain.json", appName)
	resp, err := http.Get(uri)
	if err != nil {
		return ChainRegistryResponse{}, err
	}
	defer resp.Body.Close()
	r, err := io.ReadAll(resp.Body)
	if err != nil {
		return ChainRegistryResponse{}, err
	}
	var result ChainRegistryResponse
	err = json.Unmarshal(r, &result)
	if err != nil {
		return result, err
	}
	return result, err
}
