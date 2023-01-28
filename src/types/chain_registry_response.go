package types

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

func (crr ChainRegistryResponse) GetChainName() string {
	return crr.ChainName
}

func (crr ChainRegistryResponse) GetChainID() string {
	return crr.ChainID
}

func (crr ChainRegistryResponse) GetLCDUris() []string {
	uris := make([]string, 0, len(crr.Apis.Rest))
	for _, uri := range crr.Apis.Rest {
		uris = append(uris, uri.Address)
	}
	return uris
}

func (crr ChainRegistryResponse) GetGRPCUris() []string {
	uris := make([]string, 0, len(crr.Apis.Grpc))
	for _, uri := range crr.Apis.Grpc {
		uris = append(uris, uri.Address)
	}
	return uris
}

func (crr ChainRegistryResponse) GetRPCUris() []string {
	uris := make([]string, 0, len(crr.Apis.Grpc))
	for _, uri := range crr.Apis.RPC {
		uris = append(uris, uri.Address)
	}
	return uris
}
