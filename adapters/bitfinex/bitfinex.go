package bitfinex

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
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
	cli     *sling.Sling
	symbols []string
}

func getHTTPTransport() *http.Client {
	tr := &http.Transport{
		MaxIdleConns:       defaultMaxIdleCons,
		IdleConnTimeout:    defaultIdleConnTimeout,
		DisableCompression: true,
	}

	return &http.Client{Transport: tr}
}

func NewClient(symbols []string) ports.PriceProvider {
	slingCli := sling.New().Base(baseUrl).Client(getHTTPTransport())

	return &Client{
		cli:     slingCli,
		symbols: symbols,
	}
}

func (c *Client) GetPrices(ctx context.Context) ([]map[uint]float64, error) {
	prices := make([]map[uint]float64, len(c.symbols))

	for _, symbolRaw := range c.symbols {
		symbolSplited := strings.Split(symbolRaw, "=")
		tokenID, err := strconv.Atoi(symbolSplited[0])
		if err != nil {
			return prices, err
		}
		symbol := symbolSplited[1]
		url := URL + symbol + URLExtraParams
		req, err := c.cli.New().Get(url).Request()
		if err != nil {
			return prices, err
		}

		var data interface{}
		res, err := c.cli.Do(req.WithContext(ctx), &data, nil)
		if err != nil {
			return prices, err
		}

		if data == nil {
			log.Printf("skiping token: %s\n", symbol)
			continue
		}

		// The token price is received inside an array in the sixth position
		result := make(map[uint]float64)
		value := data.([]interface{})[6].(float64)
		if res.StatusCode != http.StatusOK {
			return prices, errors.New(fmt.Sprintf("http error: %d %s", res.StatusCode, res.Status))
		}
		result[uint(tokenID)] = value

		prices = append(prices, result)
	}
	return prices, nil
}
