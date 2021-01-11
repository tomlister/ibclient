package ib

import (
	"crypto/tls"
	"log"

	resty "github.com/go-resty/resty/v2"
)

var base string

// SetBaseURL sets the base api url
func SetBaseURL(baseURL string) {
	base = baseURL
}

// KeepAlive keeps the authenticated session active.
// Inactive sessions are timed out within a few minutes.
// By calling /tickle the client portal keeps the session alive.
func KeepAlive() {
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	_, err := client.R().Post(base + "/tickle")
	if err != nil {
		log.Panic(err)
	}
}

// Authenticate initiates the brokerage session
func Authenticate() {
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	_, err := client.R().Post(base + "/api/iserver/reauthenticate")
	if err != nil {
		log.Panic(err)
	}
}
