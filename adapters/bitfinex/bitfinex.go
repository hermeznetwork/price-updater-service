package bitfinex

import (
	"context"
	"github.com/hermeznetwork/hermez-node/log"
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
	cli     *sling.Sling
	symbols map[uint]string
}

func getHTTPTransport() *http.Client {
	tr := &http.Transport{
		MaxIdleConns:       defaultMaxIdleCons,
		IdleConnTimeout:    defaultIdleConnTimeout,
		DisableCompression: true,
	}

	return &http.Client{Transport: tr}
}

func NewClient(symbols map[uint]string) ports.PriceProvider {
	slingCli := sling.New().Base(baseUrl).Client(getHTTPTransport())

	return &Client{
		cli:     slingCli,
		symbols: symbols,
	}
}

func (c *Client) GetPrices(ctx context.Context) ([]map[uint]float64, []uint, error) {
	log.Debug("Bitfinex")
	prices := make([]map[uint]float64, len(c.symbols))
	var tokenErrs []uint
	for tokenID, symbol := range c.symbols {
		url := URL + symbol + URLExtraParams
		req, err := c.cli.New().Get(url).Request()
		if err != nil {
			tokenErrs = append(tokenErrs, tokenID)
			continue
		}

		var data interface{}
		res, err := c.cli.Do(req.WithContext(ctx), &data, nil)
		if err != nil {
			log.Warn("error: ", err)
			tokenErrs = append(tokenErrs, tokenID)
			continue
		}

		if data == nil {
			log.Warn("data received empty. Skiping token: ", symbol)
			tokenErrs = append(tokenErrs, tokenID)
			continue
		}

		// The token price is received inside an array in the sixth position
		result := make(map[uint]float64)
		value := data.([]interface{})[6].(float64)
		if res.StatusCode != http.StatusOK {
			tokenErrs = append(tokenErrs, tokenID)
			log.Warn("http error: ", res.StatusCode, res.Status)
			continue
		}
		result[tokenID] = value

		prices = append(prices, result)
	}
	return prices, tokenErrs, nil
}

func (c *Client) GetFailedPrices(ctx context.Context, prices []map[uint]float64, tokenErrs []uint) ([]map[uint]float64, []uint, error) {
	log.Debug("Bitfinex")
	var tokErrs []uint
	for i := 0; i < len(tokenErrs); i++ {
		tokenID := tokenErrs[i]
		symbol := c.symbols[uint(tokenID)]
		url := URL + symbol + URLExtraParams
		req, err := c.cli.New().Get(url).Request()
		if err != nil {
			tokErrs = append(tokErrs, tokenID)
			continue
		}

		var data interface{}
		res, err := c.cli.Do(req.WithContext(ctx), &data, nil)
		if err != nil {
			log.Warn("error: ", err)
			tokErrs = append(tokErrs, tokenID)
			continue
		}

		if data == nil {
			log.Warn("data received empty. Skiping token: ", symbol)
			tokErrs = append(tokErrs, tokenID)
			continue
		}

		// The token price is received inside an array in the sixth position
		result := make(map[uint]float64)
		value := data.([]interface{})[6].(float64)
		if res.StatusCode != http.StatusOK {
			tokErrs = append(tokErrs, tokenID)
			log.Warn("http error: ", res.StatusCode, res.Status)
			continue
		}
		result[tokenID] = value

		prices = append(prices, result)
	}
	return prices, tokErrs, nil
}
