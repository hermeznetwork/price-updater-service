package bitfinex

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
	baseUrl                = "https://api-pub.bitfinex.com/v2/"
	URL                    = "ticker/t"
	URLExtraParams         = "USD"
	defaultMaxIdleCons     = 10
	defaultIdleConnTimeout = 2 * time.Second
	EmptyAddr              = "0x0000000000000000000000000000000000000000"
	FFAddr                 = "0xffffffffffffffffffffffffffffffffffffffff"
)

type Client struct {
	cli *sling.Sling
}

func getHTTPTransport() *http.Client {
	tr := &http.Transport{
		MaxIdleConns:       defaultMaxIdleCons,
		IdleConnTimeout:    defaultIdleConnTimeout,
		DisableCompression: true,
	}

	return &http.Client{Transport: tr}
}

func NewClient() ports.PriceProvider {
	slingCli := sling.New().Base(baseUrl).Client(getHTTPTransport())

	return &Client{
		cli: slingCli,
	}
}

func (c *Client) GetPrices(ctx context.Context) (float64, error) {
	symbol := "ETH"
	url := URL + symbol + URLExtraParams
	req, err := c.cli.New().Get(url).Request()
	if err != nil {
		return 0, err
	}

	var data interface{}
	res, err := c.cli.Do(req.WithContext(ctx), &data, nil)
	if err != nil {
		return 0, err
	}

	// The token price is received inside an array in the sixth position
	result := data.([]interface{})[6].(float64)
	if res.StatusCode != http.StatusOK {
		return 0, errors.New(fmt.Sprintf("http error: %s %d", res.StatusCode, res.Status))
	}
	return result, err
}
