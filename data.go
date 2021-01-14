package ib

import (
	"crypto/tls"
	"encoding/json"
	"log"
	"strconv"
	"strings"

	resty "github.com/go-resty/resty/v2"
	"github.com/rocketlaunchr/dataframe-go"
)

// HistoricalDataFrame stores the OCHL, Volume and Time of a security for a single frame of historical market data.
type HistoricalDataFrame struct {
	O float64 `json:"o"`
	C float64 `json:"c"`
	H float64 `json:"h"`
	L float64 `json:"l"`
	V int     `json:"v"`
	T int64   `json:"t"`
}

// ToDataFrame converts historical market data to dataframes for technical analysis.
func (hd Historical) ToDataFrame() *dataframe.DataFrame {
	o := dataframe.NewSeriesFloat64("open", nil)
	c := dataframe.NewSeriesFloat64("close", nil)
	h := dataframe.NewSeriesFloat64("high", nil)
	l := dataframe.NewSeriesFloat64("low", nil)
	v := dataframe.NewSeriesInt64("volume", nil)
	t := dataframe.NewSeriesInt64("time", nil)
	for _, f := range hd.Data {
		o.Append(f.O)
		c.Append(f.C)
		h.Append(f.H)
		l.Append(f.L)
		v.Append(f.V)
		t.Append(f.T)
	}
	df := dataframe.NewDataFrame(t, o, c, h, l, v)
	return df
}

// Historical stores historical market data returned from the brokerage server.
type Historical struct {
	Symbol            string                `json:"symbol"`
	Text              string                `json:"text"`
	PriceFactor       int                   `json:"priceFactor"`
	StartTime         string                `json:"startTime"`
	High              string                `json:"high"`
	Low               string                `json:"low"`
	TimePeriod        string                `json:"timePeriod"`
	BarLength         int                   `json:"barLength"`
	MdAvailability    string                `json:"mdAvailability"`
	MktDataDelay      int                   `json:"mktDataDelay"`
	OutsideRth        bool                  `json:"outsideRth"`
	VolumeFactor      int                   `json:"volumeFactor"`
	PriceDisplayRule  int                   `json:"priceDisplayRule"`
	PriceDisplayValue string                `json:"priceDisplayValue"`
	NegativeCapable   bool                  `json:"negativeCapable"`
	MessageVersion    int                   `json:"messageVersion"`
	Data              []HistoricalDataFrame `json:"data"`
	Points            int                   `json:"points"`
	TravelTime        int                   `json:"travelTime"`
}

// TimeUnit represents a unit of time.
type TimeUnit string

const (
	// Min - 1 to 30 minutes.
	Min TimeUnit = "min"
	// Hour - 1 to 8 hours.
	Hour TimeUnit = "h"
	// Day - 1 to 1000 days.
	Day TimeUnit = "d"
	// Week - 1 to 792 weeks.
	Week TimeUnit = "w"
	// Month - 1 to 182 months.
	Month TimeUnit = "m"
	// Year - 1 to 15 years.
	Year TimeUnit = "y"
)

// Historical retrieves historical market data for a security.
func (s Security) Historical(period int, unit TimeUnit, barSize int, barUnit TimeUnit) Historical {
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	resp, err := client.R().Get(base + "/api/iserver/marketdata/history?conid=" + strconv.Itoa(s.Conid) + "&period=" + strconv.Itoa(period) + string(unit) + "&bar=" + strconv.Itoa(barSize) + string(barUnit))
	if err != nil {
		log.Panic(err)
	}
	historical := Historical{}
	err = json.Unmarshal(resp.Body(), &historical)
	if err != nil {
		log.Panic(err)
	}
	return historical
}

type Snapshot struct {
	LastPrice                      float64 `json:"31,string,omitempty"`
	Symbol                         string  `json:"55,omitempty"`
	Text                           string  `json:"58,omitempty"`
	High                           float64 `json:"70,string,omitempty"`
	Low                            float64 `json:"71,string,omitempty"`
	Position                       float64 `json:"72,omitempty"`
	MarketValue                    string  `json:"73,omitempty"`
	AveragePrice                   float64 `json:"74,string,omitempty"`
	UnrealizedPnL                  float64 `json:"75,omitempty"`
	FormattedPosition              string  `json:"76,omitempty"`
	FormattedUnrealizedPnL         string  `json:"77,omitempty"`
	DailyPnL                       float64 `json:"78_raw,omitempty"`
	ChangePrice                    float64 `json:"82,string,omitempty"`
	ChangePercent                  float64 `json:"83,omitempty"`
	BidPrice                       float64 `json:"84,string,omitempty"`
	AskSize                        int     `json:"85,string,omitempty"`
	AskPrice                       float64 `json:"86,string,omitempty"`
	Volume                         float64 `json:"87_raw,omitempty"`
	BidSize                        int     `json:"88,string,omitempty"`
	SecurityType                   string  `json:"6070,omitempty"`
	MarketDataDeliveryMethodMarker string  `json:"6119,omitempty"`
	UnderlyingConid                int     `json:"6457,string,omitempty"`
	MarketDataAvailability         string  `json:"6509,omitempty"`
	CompanyName                    string  `json:"7051,omitempty"`
	LastSize                       int     `json:"7059,string,omitempty"`
	ContractDescription            string  `json:"7219,omitempty"`
	ListingExchange                string  `json:"7221,omitempty"`
	Industry                       string  `json:"7280,omitempty"`
	Category                       string  `json:"7281,omitempty"`
	AverageDailyVolume             string  `json:"7282,omitempty"`
	HistoricVolume30D              string  `json:"7284,omitempty"`
	DividendAmount                 float64 `json:"7286,string,omitempty"`
	DividendYieldPercentage        string  `json:"7287,omitempty"`
	DividendExDate                 string  `json:"7288,omitempty"`
	MarketCap                      string  `json:"7289,omitempty"`
	PE                             float64 `json:"7290,string,omitempty"`
	EPS                            float64 `json:"7291,string,omitempty"`
	CostBasis                      float64 `json:"7292_raw,omitempty"`
	WeekHigh52                     float64 `json:"7293,string,omitempty"`
	WeekLow52                      float64 `json:"7294,string,omitempty"`
	OpenPrice                      float64 `json:"7295,string,omitempty"`
	Conid                          int     `json:"conid"`
	ServerID                       string  `json:"server_id,omitempty"`
	Updated                        int64   `json:"_updated,omitempty"`
}

type Snapshots []Snapshot

// MarketDataField represents a field to request market data for.
// IB decided it was a fun idea to assign each field a number instead of a name.
// So now I must resort to this horrific mess
type MarketDataField string

const (
	LastPrice                      MarketDataField = "31"
	Symbol                         MarketDataField = "55"
	Text                           MarketDataField = "58"
	High                           MarketDataField = "70"
	Low                            MarketDataField = "71"
	Pos                            MarketDataField = "72"
	MarketValue                    MarketDataField = "73"
	AveragePrice                   MarketDataField = "74"
	UnrealizedPnL                  MarketDataField = "75"
	FormattedPosition              MarketDataField = "76"
	FormattedUnrealizedPnL         MarketDataField = "77"
	DailyPnL                       MarketDataField = "78"
	ChangePrice                    MarketDataField = "82"
	ChangePercent                  MarketDataField = "83"
	BidPrice                       MarketDataField = "84"
	AskSize                        MarketDataField = "85"
	AskPrice                       MarketDataField = "86"
	Volume                         MarketDataField = "87"
	BidSize                        MarketDataField = "88"
	Exchange                       MarketDataField = "6004"
	Conid                          MarketDataField = "6008"
	SecurityType                   MarketDataField = "6070"
	Months                         MarketDataField = "6072"
	RegularExpiry                  MarketDataField = "6073"
	MarketDataDeliveryMethodMarker MarketDataField = "6119"
	UnderlyingConid                MarketDataField = "6457"
	MarketDataAvailability         MarketDataField = "6509"
	CompanyName                    MarketDataField = "7051"
	LastSize                       MarketDataField = "7059"
	ConidExchange                  MarketDataField = "7094"
	ContractDescription            MarketDataField = "7219"
	ContractDescriptionAlt         MarketDataField = "7220"
	ListingExchange                MarketDataField = "7221"
	Industry                       MarketDataField = "7280"
	Category                       MarketDataField = "7281"
	AverageDailyVolume             MarketDataField = "7282"
	HistoricVolume30D              MarketDataField = "7284"
	PutCallRatio                   MarketDataField = "7285"
	DividendAmount                 MarketDataField = "7286"
	DividendYieldPercentage        MarketDataField = "7287"
	DividendExDate                 MarketDataField = "7288"
	MarketCap                      MarketDataField = "7289"
	PE                             MarketDataField = "7290"
	EPS                            MarketDataField = "7291"
	CostBasis                      MarketDataField = "7292"
	WeekHigh52                     MarketDataField = "7293"
	WeekLow52                      MarketDataField = "7294"
	OpenPrice                      MarketDataField = "7295"
	ClosePrice                     MarketDataField = "7296"
	Delta                          MarketDataField = "7308"
	Gamma                          MarketDataField = "7309"
	Theta                          MarketDataField = "7310"
	Vega                           MarketDataField = "7311"
	ImpliedVolatilityOption        MarketDataField = "7633"
)

// Snapshot retrieves a market data snapshot by fields
func (s Security) Snapshot(fields ...MarketDataField) Snapshots {
	fieldStrings := make([]string, 0)
	for _, f := range fields {
		fieldStrings = append(fieldStrings, string(f))
	}
	builtFields := strings.Join(fieldStrings, ",")
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	resp, err := client.R().Get(base + "/api/iserver/marketdata/snapshot?conids=" + strconv.Itoa(s.Conid) + "&fields=" + builtFields)
	if err != nil {
		log.Panic(err)
	}
	// Rerequest the snapshot
	// IBKR seems to need an initial request to initiate the market
	// data transaction and rerequesting the snapshot will
	// give us the desired data.
	resp, err = client.R().Get(base + "/api/iserver/marketdata/snapshot?conids=" + strconv.Itoa(s.Conid) + "&fields=" + builtFields)
	if err != nil {
		log.Panic(err)
	}
	snapshots := Snapshots{}
	err = json.Unmarshal(resp.Body(), &snapshots)
	if err != nil {
		log.Panic(err)
	}
	return snapshots
}
