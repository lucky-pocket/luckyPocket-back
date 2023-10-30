package gauth

import (
	"net/http"

	"github.com/onee-only/gauth-go"
)

func NewClient(httpClient *http.Client, opts gauth.ClientOpts) *gauth.Client {
	return gauth.NewClient(httpClient, opts)
}
