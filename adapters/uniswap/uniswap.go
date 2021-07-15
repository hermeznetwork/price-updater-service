package uniswap

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hermeznetwork/price-updater-service/adapters/uniswap/contract"
	"github.com/hermeznetwork/price-updater-service/adapters/uniswap/hezrollup"
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

func (c *Client) GetPrices(ctx context.Context) ([]map[uint]float64, error) {
	rollup, err := hezrollup.NewHezrollup(common.HexToAddress(c.ethConf.HezRollup), c.ethConn)
	if err != nil {
		return nil, err
	}

	registerTokenCount, err := rollup.RegisterTokensCount(nil)
	if err != nil {
		return nil, err
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
			return nil, err
		}

		usdtFilter := []common.Address{
			common.HexToAddress(c.ethConf.UsdtAddress),
		}
		t, err := uniswapToken.FilterPairCreated(nil, []common.Address{address}, usdtFilter)
		if err != nil {
			return nil, err
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
	return result, nil
}
