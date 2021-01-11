# ibclient
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
		historical := sec.Historical()
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
Once the user has the prerequisites in order, the user must sign through the client portal sso page.
It is recommended to login through your accounts dedicated paper trading username and password.
To find out more about the client portal setup process visit https://interactivebrokers.github.io/cpwebapi/

## Roadmap
As this is a fairly new package, a lot of the features you'd expect aren't implemented yet.
Rest assured they will be implemented in due time.
- [x] Portfolio
- [x] Historical Data
- [ ] Trade Execution

## Disclaimers
This is an unofficial wrapper for IBKR's Client Portal Web API.
As using any financial products naturally comes with risk, I'm not responsible for any financial damages or losses that may occur using this library to execute trades, etc...
The licence also provides no warranty what so ever.
