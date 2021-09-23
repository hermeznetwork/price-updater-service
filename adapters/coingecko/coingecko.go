package coingecko

import (
	"context"
	"github.com/hermeznetwork/hermez-node/log"
	"net/http"
	"time"
	"strings"

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

func (c *Client) GetPrices(ctx context.Context) ([]map[uint]float64, []uint, error) {
	log.Debug("CoinGecko")
	var tokenErrs []uint
	prices := make([]map[uint]float64, len(c.addresses))
	for tokenID, address := range c.addresses {
		url := getUrlByAdressValue(address)
		req, err := c.cli.New().Get(url).Request()
		if err != nil {
			log.Warn("error: ", err)
			tokenErrs = append(tokenErrs, tokenID)
			continue
		}

		var data map[string]map[string]float64
		res, err := c.cli.Do(req.WithContext(ctx), &data, nil)
		if err != nil {
			log.Warn("error: ", err)
			tokenErrs = append(tokenErrs, tokenID)
			continue
		}
		itemResponseKey := "ethereum"
		if address != EmptyAddr {
			itemResponseKey = strings.ToLower(address)
		}
		value := data[itemResponseKey]["usd"]
		if res.StatusCode != http.StatusOK || len(data) == 0 {
			tokenErrs = append(tokenErrs, tokenID)
			log.Warn("http error: ", res.StatusCode, ". TokenId: ", tokenID, ". Data received: ", data)
			continue
		}
		result := make(map[uint]float64)
		var key uint
		if itemResponseKey != "ethereum" {
			key = tokenID
		}
		result[key] = value
		prices = append(prices, result)
	}
	return prices, tokenErrs, nil
}

func (c *Client) GetFailedPrices(ctx context.Context, prices []map[uint]float64, tokenErrs []uint) ([]map[uint]float64, []uint, error) {
	log.Debug("CoinGecko")
	var tokErrs []uint
	for i := 0; i < len(tokenErrs); i++ {
		tokenID := tokenErrs[i]
		address := c.addresses[uint(tokenID)]
		url := getUrlByAdressValue(address)
		req, err := c.cli.New().Get(url).Request()
		if err != nil {
			log.Warn("error: ", err)
			tokErrs = append(tokErrs, tokenID)
			continue
		}

		var data map[string]map[string]float64
		res, err := c.cli.Do(req.WithContext(ctx), &data, nil)
		if err != nil {
			log.Warn("error: ", err)
			tokErrs = append(tokErrs, tokenID)
			continue
		}
		itemResponseKey := "ethereum"
		if address != EmptyAddr {
			itemResponseKey = strings.ToLower(address)
		}
		value := data[itemResponseKey]["usd"]
		if res.StatusCode != http.StatusOK || len(data) == 0 {
			tokErrs = append(tokErrs, tokenID)
			log.Warn("http error: ", res.StatusCode, ". TokenId: ", tokenID, ". Data received: ", data)
			continue
		}
		result := make(map[uint]float64)
		var key uint
		if itemResponseKey != "ethereum" {
			key = tokenID
		}
		result[key] = value
		prices = append(prices, result)
	}
	return prices, tokErrs, nil
}
