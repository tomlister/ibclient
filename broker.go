package ib

import (
	"crypto/tls"
	"encoding/json"
	"log"

	resty "github.com/go-resty/resty/v2"
)

// BrokerAccounts stores information about a users available brokerage accounts
type BrokerAccounts struct {
	Accounts     []string `json:"accounts"`
	ChartPeriods struct {
		STK   []string `json:"STK"`
		CFD   []string `json:"CFD"`
		OPT   []string `json:"OPT"`
		FOP   []string `json:"FOP"`
		WAR   []string `json:"WAR"`
		IOPT  []string `json:"IOPT"`
		FUT   []string `json:"FUT"`
		CASH  []string `json:"CASH"`
		IND   []string `json:"IND"`
		BOND  []string `json:"BOND"`
		FUND  []string `json:"FUND"`
		CMDTY []string `json:"CMDTY"`
		PHYSS []string `json:"PHYSS"`
	} `json:"chartPeriods"`
	SelectedAccount string `json:"selectedAccount"`
	AllowFeatures   struct {
		ShowGFIS               bool `json:"showGFIS"`
		AllowFXConv            bool `json:"allowFXConv"`
		AllowTypeAhead         bool `json:"allowTypeAhead"`
		SnapshotRefreshTimeout int  `json:"snapshotRefreshTimeout"`
		LiteUser               bool `json:"liteUser"`
		ShowWebNews            bool `json:"showWebNews"`
		Research               bool `json:"research"`
		DebugPnl               bool `json:"debugPnl"`
		ShowTaxOpt             bool `json:"showTaxOpt"`
	} `json:"allowFeatures"`
}

// BrokerAccount stores the id of a brokerage account
type BrokerAccount struct {
	ID string
}

// Brokers retrieves all of an accounts brokerage accounts
func Brokers() BrokerAccounts {
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	resp, err := client.R().Get(base + "/api/iserver/accounts")
	if err != nil {
		log.Panic(err)
	}
	brokerAccounts := BrokerAccounts{}
	err = json.Unmarshal(resp.Body(), &brokerAccounts)
	if err != nil {
		log.Panic(err)
	}
	return brokerAccounts
}

// Selected returns the default/active brokerage account
func (ba BrokerAccounts) Selected() BrokerAccount {
	brokerAccount := BrokerAccount{
		ID: ba.SelectedAccount,
	}
	return brokerAccount
}
