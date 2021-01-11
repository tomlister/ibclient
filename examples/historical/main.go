// +build example

package main

import (
	"fmt"
	"time"

	ib "github.com/tomlister/ibclient"
)

func main() {
	ib.SetBaseURL("https://localhost:5000/v1")
	// Authenticate with the brokerage server
	ib.Authenticate()
	// IBKR will time out the sso session if left inactive
	// Here KeepAlive is scheduled to run async every minute
	ib.Schedule(func() {
		ib.KeepAlive()
	}, time.Minute)
	// Grab the active brokerage account
	broker := ib.Brokers().Selected()
	// Get all of the portfolios under the user account
	portfolios := ib.Portfolios()
	// Get the positions under a specific (in this case the first) portfolio
	positions := portfolios[0].Positions()
	// Filter the positions by asset class
	futures := positions.FilterAssets(ib.Futures)
	for _, p := range futures {
		// Create a new reference to a security from a position
		sec := broker.Security(p)
		// Retrieve the historical data for that security
		historical := sec.Historical()
		fmt.Println(historical)
	}
}
