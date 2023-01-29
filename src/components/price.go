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
	var s types.CoinGeckoPriceResponse
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

	if id == "no_price" {
		p.Text = "💁🏻‍♂️ Not listed. Coming soon..."
	} else {
		price, ut, err := fethcPrice(logger, id)
		if err != nil {
			logger.Error("error while fetching the price", "err", err.Error())
			return p
		}
		p.Text = strconv.FormatFloat(price, 'g', 5, 64) + " at " + ut.Local().String() + "\nfrom CoinGecko"
	}

	return p
}

func ChainIdComponent(appName, chainID string) *widgets.Paragraph {
	p := widgets.NewParagraph()
	p.Title = fmt.Sprintf("%s chain-id", appName)
	p.Text = chainID
	return p
}
