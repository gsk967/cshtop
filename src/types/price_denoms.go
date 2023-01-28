package types

var PricesDenomsByAppName map[string]string

func init() {
	PricesDenomsByAppName = map[string]string{
		"cosmoshub": "cosmos",
		"akash":     "akash-network",
		"sentinel":  "sentinel",
		"umee":      "umee",
	}
}
