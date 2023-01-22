package components

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gizak/termui/v3/widgets"
)

type CoinGeckoPriceResponse []struct {
	ID                    string      `json:"id,omitempty"`
	Symbol                string      `json:"symbol,omitempty"`
	Name                  string      `json:"name,omitempty"`
	Image                 string      `json:"image,omitempty"`
	CurrentPrice          float64     `json:"current_price,omitempty"`
	MarketCap             int         `json:"market_cap,omitempty"`
	MarketCapRank         int         `json:"market_cap_rank,omitempty"`
	FullyDilutedValuation int         `json:"fully_diluted_valuation,omitempty"`
	TotalVolume           int         `json:"total_volume,omitempty"`
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

var pricesDenomsByAppName map[string]string

func init() {
	pricesDenomsByAppName = map[string]string{
		"cosmoshub": "cosmos",
		"akash":     "akash-network",
		"sentinel":  "sentinel",
		"umee":      "umee",
	}
}
func fethcPrice(id string) (float64, error) {
	uri := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&ids=%s&order=market_cap_desc&per_page=100&page=1&sparkline=false", id)
	resp, err := http.Get(uri)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	r, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	var s CoinGeckoPriceResponse
	err = json.Unmarshal(r, &s)
	if err != nil {
		return 0, err
	}

	return s[0].CurrentPrice, nil
}

// PriceComponent
func PriceComponent(appName, pName string) *widgets.Paragraph {
	id, ok := pricesDenomsByAppName[appName]
	if !ok {
		fmt.Printf("There is no id for app %s", appName)
		return nil
	}
	p := widgets.NewParagraph()
	p.Title = fmt.Sprintf("%s Price", pName)

	price, err := fethcPrice(id)
	if err != nil {
		fmt.Println("Err while fetching the price", err)
	}
	p.Text = strconv.FormatFloat(price, 'g', 5, 64) + " at " + time.Now().Local().String()

	return p
}

func ChainIdComponent(appName, chainID string) *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.Title = fmt.Sprintf("%s Chain ID", appName)
	p.Text = chainID
	return p
}
