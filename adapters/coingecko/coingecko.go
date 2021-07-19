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
	cli       *sling.Sling
	addresses map[uint]string
}

func getHTTPTransport() *http.Client {
	tr := &http.Transport{
		MaxIdleConns:       defaultMaxIdleCons,
		IdleConnTimeout:    defaultIdleConnTimeout,
		DisableCompression: true,
	}

	return &http.Client{Transport: tr}
}

func NewClient(addresses map[uint]string) ports.PriceProvider {
	slingCli := sling.New().Base(baseUrl).Client(getHTTPTransport())

	return &Client{
		cli:       slingCli,
		addresses: addresses,
	}
}

func getUrlByAdressValue(address string) string {
	if address == EmptyAddr {
		return "simple/price?ids=ethereum" + URLExtraParams
	}
	return URL + address + URLExtraParams
}

func (c *Client) GetPrices(ctx context.Context) ([]map[uint]float64, error) {
	prices := make([]map[uint]float64, len(c.addresses))
	for tokenID, address := range c.addresses {
		url := getUrlByAdressValue(address)
		req, err := c.cli.New().Get(url).Request()
		if err != nil {
			return prices, err
		}

		var data map[string]map[string]float64
		res, err := c.cli.Do(req.WithContext(ctx), &data, nil)
		if err != nil {
			return prices, err
		}
		itemResponseKey := "ethereum"
		if address != EmptyAddr {
			itemResponseKey = address
		}
		value := data[itemResponseKey]["usd"]
		if res.StatusCode != http.StatusOK {
			return prices, errors.New(fmt.Sprintf("http error: %d %s", res.StatusCode, res.Status))
		}
		result := make(map[uint]float64)
		var key uint
		if itemResponseKey != "ethereum" {
			key = tokenID
		}
		result[key] = value
		prices = append(prices, result)
	}
	return prices, nil

}
