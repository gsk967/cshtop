package components

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gizak/termui/v3/widgets"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/gsk967/cshtop/src/types"
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

func fethcPrice(logger log.Logger, id string) (float64, time.Time, error) {
	uri := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&ids=%s&order=market_cap_desc&per_page=100&page=1&sparkline=false", id)
	resp, err := http.Get(uri)
	if err != nil {
		logger.Error("error at get the price from coingecko", "uri", uri, "err", err.Error())
		return 0, time.Time{}, err
	}
	defer resp.Body.Close()
	r, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("error at read the price from coingecko", "uri", uri, "err", err.Error())
		return 0, time.Time{}, err
	}
	var s CoinGeckoPriceResponse
	err = json.Unmarshal(r, &s)
	if err != nil {
		logger.Error("error at unmarshal the price from coingecko", "uri", uri, "err", err.Error())
		return 0, time.Time{}, err
	}
	// logger.Debug("get the price from coingecko", "price", s[0].CurrentPrice)
	return s[0].CurrentPrice, s[0].LastUpdated, nil
}

// PriceComponent
func PriceComponent(logger log.Logger, appName, pName string) *widgets.Paragraph {
	id, ok := types.PricesDenomsByAppName[appName]
	if !ok {
		logger.Error("there is no CoinGecko id", "app", appName)
		return nil
	}

	p := widgets.NewParagraph()
	p.Title = fmt.Sprintf("%s Price", pName)

	price, ut, err := fethcPrice(logger, id)
	if err != nil {
		logger.Error("error while fetching the price", "err", err.Error())
		return p
	}

	p.Text = strconv.FormatFloat(price, 'g', 5, 64) + " at " + ut.Local().String() + "\nfrom CoinGecko"
	return p
}

func ChainIdComponent(appName, chainID string) *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.Title = fmt.Sprintf("%s chain-id", appName)
	p.Text = chainID
	return p
}
