package types

import "time"

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

type ValidatorMap map[string]string
