package uniswap

import (
	"context"
	"math/big"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hermeznetwork/price-updater-service/adapters/uniswap/contract"
	"github.com/hermeznetwork/price-updater-service/config"
	"github.com/hermeznetwork/price-updater-service/core/ports"
)

type Client struct {
	ethConn *ethclient.Client
	ethConf config.EthConfig
}

// uniswapAddress works to all networks (Mainnet, Goerli, Rinkeby)
const uniswapAddress = "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"

func NewClient(cfg config.EthConfig) (ports.PriceProvider, error) {
	conn, err := ethclient.Dial(cfg.EthNetwork)
	if err != nil {
		return nil, err
	}
	return &Client{ethConn: conn, ethConf: cfg}, nil
}

func (c *Client) GetPrices(ctx context.Context) ([]map[uint]float64, []uint, error) {
	var tokenErrs []uint
	rollup, err := contract.NewHezrollup(common.HexToAddress(c.ethConf.HezRollup), c.ethConn)
	if err != nil {
		return nil, nil, err
	}

	registerTokenCount, err := rollup.RegisterTokensCount(nil)
	if err != nil {
		return nil, nil, err
	}

	addresses := make(map[uint]common.Address)
	for i := int64(0); i < registerTokenCount.Int64(); i++ {
		tokenAddress, err := rollup.TokenList(nil, big.NewInt(i))
		if err != nil {
			// TODO: logging to fail
			continue
		}
		addresses[uint(i)] = tokenAddress
	}

	result := make([]map[uint]float64, len(addresses))

	for tokenID, address := range addresses {
		uniswapToken, err := contract.NewToken(common.HexToAddress(uniswapAddress), c.ethConn)
		if err != nil {
			log.Println("error: ", err)
			tokenErrs = append(tokenErrs, tokenID)
			continue
		}

		usdtFilter := []common.Address{
			common.HexToAddress(c.ethConf.UsdtAddress),
		}
		t, err := uniswapToken.FilterPairCreated(nil, []common.Address{address}, usdtFilter)
		if err != nil {
			log.Println("error: ", err)
			tokenErrs = append(tokenErrs, tokenID)
			continue
		}

		for t.Next() {
			pairAddress, err := uniswapToken.GetPair(nil, t.Event.Token0, t.Event.Token1)
			if err != nil {
				// TODO: logging information about try
				continue
			}

			pricesFromPairs, err := getPriceFromPairsInfo(pairAddress, address, c.ethConn)
			if err != nil {
				// TODO: logging information about try
				continue
			}
			float64Value, _ := pricesFromPairs.price.Float64()
			r := make(map[uint]float64)
			r[tokenID] = float64Value
			result = append(result, r)
		}
	}
	return result, tokenErrs, nil
}

func (c *Client) GetFailedPrices(ctx context.Context, prices []map[uint]float64, tokenErrs []uint) ([]map[uint]float64, []uint, error) {
	var tokErrs []uint
	rollup, err := contract.NewHezrollup(common.HexToAddress(c.ethConf.HezRollup), c.ethConn)
	if err != nil {
		return nil, nil, err
	}

	addresses := make(map[uint]common.Address)
	for i := 0; i < len(tokenErrs); i++ {
		tokenAddress, err := rollup.TokenList(nil, big.NewInt(int64(i)))
		if err != nil {
			// TODO: logging to fail
			continue
		}
		addresses[uint(i)] = tokenAddress
	}

	for tokenID, address := range addresses {
		uniswapToken, err := contract.NewToken(common.HexToAddress(uniswapAddress), c.ethConn)
		if err != nil {
			log.Println("error: ", err)
			tokErrs = append(tokErrs, tokenID)
			continue
		}

		usdtFilter := []common.Address{
			common.HexToAddress(c.ethConf.UsdtAddress),
		}
		t, err := uniswapToken.FilterPairCreated(nil, []common.Address{address}, usdtFilter)
		if err != nil {
			log.Println("error: ", err)
			tokErrs = append(tokErrs, tokenID)
			continue
		}

		for t.Next() {
			pairAddress, err := uniswapToken.GetPair(nil, t.Event.Token0, t.Event.Token1)
			if err != nil {
				// TODO: logging information about try
				continue
			}

			pricesFromPairs, err := getPriceFromPairsInfo(pairAddress, address, c.ethConn)
			if err != nil {
				// TODO: logging information about try
				continue
			}
			float64Value, _ := pricesFromPairs.price.Float64()
			r := make(map[uint]float64)
			r[tokenID] = float64Value
			prices = append(prices, r)
		}
	}
	return prices, tokErrs, nil
}