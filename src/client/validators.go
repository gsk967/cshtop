package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Validator struct {
	OperatorAddress string `json:"operator_address,omitempty"`
	ConsensusPubkey struct {
		Type string `json:"@type,omitempty"`
		Key  string `json:"key,omitempty"`
	} `json:"consensus_pubkey,omitempty"`
	Jailed          bool   `json:"jailed,omitempty"`
	Status          string `json:"status,omitempty"`
	Tokens          string `json:"tokens,omitempty"`
	DelegatorShares string `json:"delegator_shares,omitempty"`
	Description     struct {
		Moniker         string `json:"moniker,omitempty"`
		Identity        string `json:"identity,omitempty"`
		Website         string `json:"website,omitempty"`
		SecurityContact string `json:"security_contact,omitempty"`
		Details         string `json:"details,omitempty"`
	} `json:"description,omitempty"`
	UnbondingHeight string    `json:"unbonding_height,omitempty"`
	UnbondingTime   time.Time `json:"unbonding_time,omitempty"`
	Commission      struct {
		CommissionRates struct {
			Rate          string `json:"rate,omitempty"`
			MaxRate       string `json:"max_rate,omitempty"`
			MaxChangeRate string `json:"max_change_rate,omitempty"`
		} `json:"commission_rates,omitempty"`
		UpdateTime time.Time `json:"update_time,omitempty"`
	} `json:"commission,omitempty"`
	MinSelfDelegation string `json:"min_self_delegation,omitempty"`
}

type ValidatorsResp struct {
	Validators []Validator `json:"validators,omitempty"`
	Pagination struct {
		NextKey string `json:"next_key,omitempty"`
		Total   string `json:"total,omitempty"`
	} `json:"pagination,omitempty"`
}

func getValidators(lcd, uriQuery string, validators []Validator) ([]Validator, int) {
	uri := fmt.Sprintf("%s?%s", lcd, uriQuery)
	fmt.Println("uri ", uri)
	resp, err := http.Get(uri)
	if err != nil {
		fmt.Println("err ", err)
		return validators, http.StatusInternalServerError
	}

	if resp.StatusCode == 200 {
		defer resp.Body.Close()
		r, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("err ", err)
		}
		var s ValidatorsResp
		err = json.Unmarshal(r, &s)
		if err != nil {
			fmt.Println("err ", err)
		}
		validators = append(validators, s.Validators...)
		if len(s.Pagination.NextKey) != 0 {
			uriQuery = fmt.Sprintf("status=BOND_STATUS_BONDED&pagination.key=%s", s.Pagination.NextKey)
			getValidators(lcd, uriQuery, validators)
		}
	} else {
		return validators, resp.StatusCode
	}

	return validators, resp.StatusCode
}

// GetValidators will returns bonded validators from lcd list
func GetValidators(lcdUris []string) []Validator {
	uriQuery := "status=BOND_STATUS_BONDED"
	for _, lcdUri := range lcdUris {
		lcdUri := fmt.Sprintf("%s/cosmos/staking/v1beta1/validators", lcdUri)
		validators, statusCode := getValidators(lcdUri, uriQuery, []Validator{})
		if statusCode == 200 {
			fmt.Println("founnd the validators from uri ", lcdUri)
			return validators
		} else {
			fmt.Println("not founnd the validators from uri ", lcdUri)
			continue
		}
	}

	return []Validator{}
}
