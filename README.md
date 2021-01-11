# ibclient
IBKR Client Portal Web API Wrapper for Go

The simple solution to creating trading algorithms for IBKR accounts.

## Example
Retrieve all positions in a portfolio, filter by asset class and get historical data.
```go
package main

import (
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
	var wg sync.WaitGroup
	wg.Add(1)
	go StartWebServer(&wg, broker, portfolios[0])
	positions := portfolios[0].Positions()
	futures := positions.FilterAssets(ib.Futures)
	for _, p := range futures {
		sec := broker.Security(p)
		historical := sec.Historical()
		fmt.Println(historical)
	}
}
```
