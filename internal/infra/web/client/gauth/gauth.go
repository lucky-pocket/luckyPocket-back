package gauth

import (
	"github.com/onee-only/gauth-go"
	"net/http"
)

func NewClient(httpClient *http.Client, opts gauth.ClientOpts) *gauth.Client {
	return gauth.NewClient(http.DefaultClient, opts)
}
