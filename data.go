package ib

import (
	"crypto/tls"
	"encoding/json"
	"log"
	"strconv"

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

// Historical retrieves historical market data for a security.
func (s Security) Historical() Historical {
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	resp, err := client.R().Get(base + "/api/iserver/marketdata/history?conid=" + strconv.Itoa(s.Conid) + "&period=2d")
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
