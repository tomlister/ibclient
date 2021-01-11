# ibclient
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/tomlister/ibclient/Go)
[![GoDoc](https://img.shields.io/badge/godoc-ibclient-007d9c)](https://pkg.go.dev/github.com/tomlister/ibclient)
[![Twitter](https://img.shields.io/twitter/url?style=social&url=https%3A%2F%2Fgithub.com%2Ftomlister%2Fibclient)](https://twitter.com/intent/tweet?text=Check%20out&url=https%3A%2F%2Fgithub.com%2Ftomlister%2Fibclient)

IBKR Client Portal Web API Wrapper for Go

The simple solution to creating trading algorithms for IBKR accounts.

## Example
Retrieve all positions in a portfolio, filter by asset class and get historical data.
```go
package main

import (
	"fmt"
	"time"
	ib "github.com/tomlister/ibclient"
)

func main() {
  	ib.SetBaseURL("https://localhost:5000/v1")
	ib.Authenticate()
	ib.Schedule(func() {
		ib.KeepAlive()
	}, time.Minute)
	broker := ib.Brokers().Selected()
	portfolios := ib.Portfolios()
	positions := portfolios[0].Positions()
	futures := positions.FilterAssets(ib.Futures)
	for _, p := range futures {
		sec := broker.Security(p)
		historical := sec.Historical(2, ib.Day, 1, ib.Hour)
		fmt.Println(historical)
	}
}
```
To run this example: `go run -tags=example github.com/tomlister/ibclient/examples/historical`

## Installing
`go get github.com/tomlister/ibclient`

## Prerequisites
In order to use this package the user must have the following:
- An IBKR Pro Brokerage Account
- The IBKR Client Portal Web API running locally

## Running
Once the prerequisites are in place, the user must sign through the client portal SSO.
Logging in with a dedicated paper trading username and password is recommended.
More info about the client portal setup process can be found at https://interactivebrokers.github.io/cpwebapi/

## Roadmap
As this is a fairly new package, a lot of the features you'd expect aren't implemented yet.
Rest assured they will be implemented in due time.
- [x] Portfolio
- [x] Historical Data
- [ ] Live Data
- [ ] Trade Execution
- [ ] Scanners

## Disclaimers
This is an unofficial wrapper for IBKR's Client Portal Web API.
As using any financial products naturally comes with risk, I'm not responsible for any financial damages or losses that may occur using this library to execute trades, etc...
The licence also provides no warranty what so ever.
