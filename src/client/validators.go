package client

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/gsk967/cshtop/src/types"
	"github.com/tendermint/tendermint/libs/log"
)

func getValidators(logger log.Logger, lcd string, uriQuery map[string]string, validators []types.Validator) ([]types.Validator, int) {
	// uri := fmt.Sprintf("%s?%s", lcd, uriQuery)
	uri, err := url.Parse(lcd)
	if err != nil {
		logger.Error("👎🏻 failed to parse uri ", "uri", uri.String(), "err", err.Error())
	}
	query := uri.Query()
	for k, v := range uriQuery {
		query.Set(k, v)
	}
	uri.RawQuery = query.Encode()
	resp, err := http.Get(uri.String())
	if err != nil {
		logger.Error("👎🏻 failed to get validators from uri ", "uri", uri.String(), "err", err.Error())
		return validators, http.StatusInternalServerError
	}

	if resp.StatusCode == 200 {
		defer resp.Body.Close()
		r, err := io.ReadAll(resp.Body)
		if err != nil {
			logger.Error("👎🏻 failed to read the validators from uri ", "uri", uri.String(), "err", err.Error())
		}
		var s types.ValidatorsResp
		err = json.Unmarshal(r, &s)
		if err != nil {
			logger.Error("👎🏻 failed to unmarshal the validators from uri ", "uri", uri.String(), "err", err.Error())
		}
		logger.Info("👍 found the validators", "uri", uri, "count", len(s.Validators))
		validators = append(validators, s.Validators...)

		if len(s.Pagination.NextKey) != 0 {
			// uriQuery = fmt.Sprintf("status=BOND_STATUS_BONDED&pagination.key=%s", s.Pagination.NextKey)
			uriQuery = map[string]string{
				"status":         "BOND_STATUS_BONDED",
				"pagination.key": s.Pagination.NextKey,
			}
			return getValidators(logger, lcd, uriQuery, validators)
		}
	} else {
		return validators, resp.StatusCode
	}

	return validators, resp.StatusCode
}

// GetValidators will returns bonded validators from lcd list
func GetValidators(logger log.Logger, lcdUris []string) []types.Validator {
	uriQuery := map[string]string{
		"status": "BOND_STATUS_BONDED",
	}
	for _, lcdUri := range lcdUris {
		lcdUri := fmt.Sprintf("%s/cosmos/staking/v1beta1/validators", lcdUri)
		validators, statusCode := getValidators(logger, lcdUri, uriQuery, []types.Validator{})
		if statusCode == 200 {
			logger.Info("👍 found the validators", "uri", lcdUri, "total_active", len(validators))
			return validators
		} else {
			logger.Error("👎🏻 failed to get validators from uri ", "uri", lcdUri)
			continue
		}
	}

	return []types.Validator{}
}

// GetMapValidators
func GetMapValidators(validators []types.Validator) types.ValidatorMap {
	vals := make(types.ValidatorMap)
	for _, val := range validators {
		e, _ := base64.StdEncoding.DecodeString(val.ConsensusPubkey.Key)
		pk := &ed25519.PubKey{Key: e}
		vals[pk.Address().String()] = val.Description.Moniker
	}
	return vals
}
