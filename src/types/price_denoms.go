package types

import "time"

var PricesDenomsByAppName map[string]string

func init() {
	PricesDenomsByAppName = map[string]string{
		"cosmoshub":   "cosmos",
		"akash":       "akash-network",
		"sentinel":    "sentinel",
		"umee":        "umee",
		"agoric":      "agoric",
		"assetmantle": "assetmantle",
		"axelar":      "axelar",
		"bitcanna":    "bitcanna",
		"bitsong":     "bitsong",
		"cerberus":    "cerberus-2",
		"chihuahua":   "chihuahua-token",
		"comdex":      "comdex",
		"crescent":    "crescent-network",
		// "cryptoorgchain": "crypto-com-chain",
		"desmos":      "desmos",
		"osmosis":     "osmosis",
		"regen":       "regen",
		"juno":        "juno-network",
		"evmos":       "evmos",
		"fetchhub":    "fetch-ai",
		"starname":    "starname",
		"stride":      "stride",
		"stargaze":    "stargaze",
		"persistence": "persistence",
		"kava":        "kava",
		"teritori":    "teritori",
		"rizon":       "rizon",
		"passage":     "no_price",
		"emoney":      "no_price",
		"nyx":         "no_price",
	}
}

type CoinGeckoPriceResponse []struct {
	ID                    string      `json:"id,omitempty"`
	Symbol                string      `json:"symbol,omitempty"`
	Name                  string      `json:"name,omitempty"`
	Image                 string      `json:"image,omitempty"`
	CurrentPrice          float64     `json:"current_price,omitempty"`
	MarketCap             float64     `json:"market_cap,omitempty"`
	MarketCapRank         int         `json:"market_cap_rank,omitempty"`
	FullyDilutedValuation int         `json:"fully_diluted_valuation,omitempty"`
	TotalVolume           float64     `json:"total_volume,omitempty"`
	High24H               float64     `json:"high_24h,omitempty"`
	Low24H                float64     `json:"low_24h,omitempty"`
	PriceChange24H        float64     `json:"price_change_24h,omitempty"`
	CirculatingSupply     float64     `json:"circulating_supply,omitempty"`
	TotalSupply           float64     `json:"total_supply,omitempty"`
	MaxSupply             interface{} `json:"max_supply,omitempty"`
	Ath                   float64     `json:"ath,omitempty"`
	AthChangePercentage   float64     `json:"ath_change_percentage,omitempty"`
	AthDate               time.Time   `json:"ath_date,omitempty"`
	Atl                   float64     `json:"atl,omitempty"`
	AtlChangePercentage   float64     `json:"atl_change_percentage,omitempty"`
	AtlDate               time.Time   `json:"atl_date,omitempty"`
	Roi                   interface{} `json:"roi,omitempty"`
	LastUpdated           time.Time   `json:"last_updated,omitempty"`
}
