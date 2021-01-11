package ib

import (
	"crypto/tls"
	"encoding/json"
	"log"

	resty "github.com/go-resty/resty/v2"
)

// Portfolio stores information about a specific account portfolio
type Portfolio struct {
	ID             string      `json:"id"`
	AccountID      string      `json:"accountId"`
	AccountVan     string      `json:"accountVan"`
	AccountTitle   string      `json:"accountTitle"`
	DisplayName    string      `json:"displayName"`
	AccountAlias   interface{} `json:"accountAlias"`
	AccountStatus  int64       `json:"accountStatus"`
	Currency       string      `json:"currency"`
	Type           string      `json:"type"`
	TradingType    string      `json:"tradingType"`
	Faclient       bool        `json:"faclient"`
	ClearingStatus string      `json:"clearingStatus"`
	Parent         interface{} `json:"parent"`
	Desc           string      `json:"desc"`
	Covestor       bool        `json:"covestor"`
}

// PortfoliosResponse is an array of portfolio accounts
type PortfoliosResponse []Portfolio

// Portfolios retrieves portfolios attached to the account
func Portfolios() PortfoliosResponse {
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	resp, err := client.R().Get(base + "/api/portfolio/accounts")
	if err != nil {
		log.Panic(err)
	}
	portfolios := PortfoliosResponse{}
	err = json.Unmarshal(resp.Body(), &portfolios)
	if err != nil {
		log.Panic(err)
	}
	return portfolios
}

// Position stores information about a position
type Position struct {
	AcctID            string        `json:"acctId"`
	Conid             int           `json:"conid"`
	ContractDesc      string        `json:"contractDesc"`
	Position          float64       `json:"position"`
	MktPrice          float64       `json:"mktPrice"`
	MktValue          float64       `json:"mktValue"`
	Currency          string        `json:"currency"`
	AvgCost           float64       `json:"avgCost"`
	AvgPrice          float64       `json:"avgPrice"`
	RealizedPnl       float64       `json:"realizedPnl"`
	UnrealizedPnl     float64       `json:"unrealizedPnl"`
	Exchs             interface{}   `json:"exchs"`
	Expiry            string        `json:"expiry"`
	PutOrCall         interface{}   `json:"putOrCall"`
	Multiplier        string        `json:"multiplier"`
	Strike            float64       `json:"strike"`
	ExerciseStyle     interface{}   `json:"exerciseStyle"`
	ConExchMap        []interface{} `json:"conExchMap"`
	AssetClass        string        `json:"assetClass"`
	UndConid          int           `json:"undConid"`
	Model             string        `json:"model"`
	BaseMktValue      float64       `json:"baseMktValue"`
	BaseMktPrice      float64       `json:"baseMktPrice"`
	BaseAvgCost       float64       `json:"baseAvgCost"`
	BaseAvgPrice      float64       `json:"baseAvgPrice"`
	BaseRealizedPnl   float64       `json:"baseRealizedPnl"`
	BaseUnrealizedPnl float64       `json:"baseUnrealizedPnl"`
	Time              int           `json:"time"`
	Name              string        `json:"name"`
	LastTradingDay    string        `json:"lastTradingDay"`
	Group             interface{}   `json:"group"`
	Sector            interface{}   `json:"sector"`
	SectorGroup       interface{}   `json:"sectorGroup"`
	Ticker            string        `json:"ticker"`
	Type              string        `json:"type"`
	UndComp           string        `json:"undComp,omitempty"`
	UndSym            string        `json:"undSym,omitempty"`
	FullName          string        `json:"fullName"`
	PageSize          int           `json:"pageSize"`
	ChineseName       string        `json:"chineseName,omitempty"`
}

// Positions is an array of positions
type Positions []Position

// Positions retrieves a portfolios positions
func (p Portfolio) Positions() Positions {
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	resp, err := client.R().Get(base + "/api/portfolio/" + p.AccountID + "/positions/0")
	if err != nil {
		log.Panic(err)
	}
	positions := Positions{}
	err = json.Unmarshal(resp.Body(), &positions)
	if err != nil {
		log.Panic(err)
	}
	return positions
}

// AssetClass represents the type of asset
type AssetClass string

const (
	// Futures are contracts that allow assets to be purchased at a market price for delivery and payment in the future.
	Futures AssetClass = "FUT"
	// Stocks represent a fractional ownership of a business and entitles the holder to dividends.
	Stocks AssetClass = "STK"
)

// FilterAssets filters positions by asset class
func (p Positions) FilterAssets(a AssetClass) Positions {
	tmp := make([]Position, 0)
	for _, pos := range p {
		if pos.AssetClass == string(a) {
			tmp = append(tmp, pos)
		}
	}
	return tmp
}
