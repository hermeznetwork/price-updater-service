package fiat

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dghubble/sling"
	"github.com/hermeznetwork/price-updater-service/core/ports"
)

const (
	baseURL                = "https://api.exchangeratesapi.io/v1/"
	baseCurrency           = "USD"
	currencies             = "CNY,EUR,JPY,GBP"
	defaultMaxIdleCons     = 10
	defaultIdleConnTimeout = 2 * time.Second
	EmptyAddr              = "0x0000000000000000000000000000000000000000"
	FFAddr                 = "0xffffffffffffffffffffffffffffffffffffffff"
)

type fiatResult struct {
	Base  string             `json:"base"`
	Rates map[string]float64 `json:"rates"`
}

type Client struct {
	cli    *sling.Sling
	apiKey string
}

func getHTTPTransport() *http.Client {
	tr := &http.Transport{
		MaxIdleConns:       defaultMaxIdleCons,
		IdleConnTimeout:    defaultIdleConnTimeout,
		DisableCompression: true,
	}

	return &http.Client{Transport: tr}
}

func NewClient(apiKey string) ports.FiatProvider {
	slingCli := sling.New().Base(baseURL).Client(getHTTPTransport())

	return &Client{
		cli:    slingCli,
		apiKey: apiKey,
	}
}

func (c *Client) GetPrices(ctx context.Context) (map[string]float64, error) {
	var rates map[string]float64
	url := "latest?base=" + baseCurrency + "&symbols=" + currencies + "&access_key=" + c.apiKey
	req, err := c.cli.New().Get(url).Request()
	if err != nil {
		return rates, err
	}

	fiatResp := fiatResult{}
	res, err := c.cli.Do(req.WithContext(ctx), &fiatResp, nil)
	if err != nil {
		return rates, err
	}

	if res.StatusCode != http.StatusOK {
		return rates, errors.New(fmt.Sprintf("http error: %d %s", res.StatusCode, res.Status))
	}

	rates = fiatResp.Rates
	fmt.Printf("%+v\n", rates)
	return rates, nil
}
