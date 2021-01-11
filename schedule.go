package ib

import "time"

// Schedule runs a function periodically
func Schedule(callback func(), interval time.Duration) chan struct{} {
	ticker := time.NewTicker(interval)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				callback()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	return quit
}
