package bitfinex

import (
	"context"
	"log"
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
			log.Println("error: ", err)
			tokenErrs = append(tokenErrs, tokenID)
			continue
		}

		if data == nil {
			log.Printf("data received empty. Skiping token: %s\n", symbol)
			tokenErrs = append(tokenErrs, tokenID)
			continue
		}

		// The token price is received inside an array in the sixth position
		result := make(map[uint]float64)
		value := data.([]interface{})[6].(float64)
		if res.StatusCode != http.StatusOK {
			tokenErrs = append(tokenErrs, tokenID)
			log.Println("http error: ", res.StatusCode, res.Status)
			continue
		}
		result[tokenID] = value

		prices = append(prices, result)
	}
	return prices, tokenErrs, nil
}

func (c *Client) GetFailedPrices(ctx context.Context, prices []map[uint]float64, tokenErrs []uint) ([]map[uint]float64, []uint, error) {
	var tokErrs []uint
	for i:=0; i<len(tokenErrs); i++ {
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
			log.Println("error: ", err)
			tokErrs = append(tokErrs, tokenID)
			continue
		}

		if data == nil {
			log.Printf("data received empty. Skiping token: %s\n", symbol)
			tokErrs = append(tokErrs, tokenID)
			continue
		}

		// The token price is received inside an array in the sixth position
		result := make(map[uint]float64)
		value := data.([]interface{})[6].(float64)
		if res.StatusCode != http.StatusOK {
			tokErrs = append(tokErrs, tokenID)
			log.Println("http error: ", res.StatusCode, res.Status)
			continue
		}
		result[tokenID] = value

		prices = append(prices, result)
	}
	return prices, tokErrs, nil
}
