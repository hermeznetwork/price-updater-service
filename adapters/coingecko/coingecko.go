package coingecko

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
	baseUrl                = "https://api.coingecko.com/api/v3/"
	URL                    = "simple/token_price/ethereum?contract_addresses="
	URLExtraParams         = "&vs_currencies=usd"
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

func getUrlByAdressValue(address string) string {
	if address == EmptyAddr {
		return "simple/price?ids=ethereum" + URLExtraParams
	}
	return URL + address + URLExtraParams
}

func (c *Client) GetPrices(ctx context.Context) (float64, error) {
	// Get the address from somewhere
	address := "0x0000000000000000000000000000000000000000"
	url := getUrlByAdressValue(address)
	req, err := c.cli.New().Get(url).Request()
	if err != nil {
		return 0, err
	}

	var data map[string]map[string]float64
	res, err := c.cli.Do(req.WithContext(ctx), &data, nil)
	if err != nil {
		return 0, err
	}
	itemResponseKey := "ethereum"
	if address != EmptyAddr {
		itemResponseKey = address
	}
	result := data[itemResponseKey]["usd"]
	if res.StatusCode != http.StatusOK {
		return 0, errors.New(fmt.Sprintf("http error: %s %s", res.StatusCode, res.Status))
	}
	return result, nil
}
