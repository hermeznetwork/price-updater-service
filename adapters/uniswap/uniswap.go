package uniswap

import (
	"context"
	"github.com/hermeznetwork/hermez-node/log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hermeznetwork/price-updater-service/adapters/uniswap/contract"
	"github.com/hermeznetwork/price-updater-service/config"
	"github.com/hermeznetwork/price-updater-service/core/ports"
)

type Client struct {
	ethConn   *ethclient.Client
	ethConf   config.EthConfig
	addresses map[uint]string
}

// uniswapAddress works to all networks (Mainnet, Goerli, Rinkeby)
const uniswapAddress = "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"

func NewClient(cfg config.EthConfig, addresses map[uint]string) (ports.PriceProvider, error) {
	conn, err := ethclient.Dial(cfg.EthNetwork)
	if err != nil {
		return nil, err
	}
	return &Client{ethConn: conn, ethConf: cfg, addresses: addresses}, nil
}

func (c *Client) GetPrices(ctx context.Context) ([]map[uint]float64, []uint, error) {
	log.Debug("Uniswap")
	var tokenErrs []uint
	result := make([]map[uint]float64, len(c.addresses))

	for tokenID, address := range c.addresses {
		uniswapToken, err := contract.NewToken(common.HexToAddress(uniswapAddress), c.ethConn)
		if err != nil {
			log.Warn("error: ", err)
			tokenErrs = append(tokenErrs, tokenID)
			continue
		}

		usdtFilter := []common.Address{
			common.HexToAddress(c.ethConf.UsdtAddress),
		}
		addr := common.HexToAddress(address)
		var tokens = [][]common.Address{usdtFilter, {addr}}
		for i := 0; i < 2; i++ {
			t, err := uniswapToken.FilterPairCreated(nil, tokens[1-i], tokens[i])
			if err != nil {
				tokenErrs = append(tokenErrs, tokenID)
				log.Warn("error: ", err)
				continue
			}

			for t.Next() {
				pairAddress, err := uniswapToken.GetPair(nil, t.Event.Token0, t.Event.Token1)
				if err != nil {
					tokenErrs = append(tokenErrs, tokenID)
					log.Warn("error getting pair: ", err)
					continue
				}

				pricesFromPairs, err := getPriceFromPairsInfo(pairAddress, addr, c.ethConn)
				if err != nil {
					tokenErrs = append(tokenErrs, tokenID)
					log.Warn("error getting price form pairs: ", err, ". Address: ",addr, ".TokenId: ", tokenID)
					continue
				}
				float64Value, _ := pricesFromPairs.price.Float64()
				r := make(map[uint]float64)
				r[tokenID] = float64Value
				result = append(result, r)
			}
		}
	}
	return result, tokenErrs, nil
}

func (c *Client) GetFailedPrices(ctx context.Context, prices []map[uint]float64, tokenErrs []uint) ([]map[uint]float64, []uint, error) {
	log.Debug("Uniswap")
	var tokErrs []uint

	addresses := make(map[uint]common.Address)
	for tokenId, addr := range c.addresses{
		for _, tokenErr := range tokenErrs {
			if tokenErr == tokenId {
				addresses[tokenErr] = common.HexToAddress(addr)
			}
		}
	}

	for tokenID, address := range addresses {
		uniswapToken, err := contract.NewToken(common.HexToAddress(uniswapAddress), c.ethConn)
		if err != nil {
			log.Warn("error: ", err)
			tokErrs = append(tokErrs, tokenID)
			continue
		}

		usdtFilter := []common.Address{
			common.HexToAddress(c.ethConf.UsdtAddress),
		}
		var tokens = [][]common.Address{usdtFilter, {address}}
		for i := 0; i < 2; i++ {
			t, err := uniswapToken.FilterPairCreated(nil, tokens[1-i], tokens[i])
			if err != nil {
				tokErrs = append(tokErrs, tokenID)
				log.Warn("error: ", err)
				continue
			}

			for t.Next() {
				pairAddress, err := uniswapToken.GetPair(nil, t.Event.Token0, t.Event.Token1)
				if err != nil {
					tokErrs = append(tokErrs, tokenID)
					log.Warn("error getting pair: ", err)
					continue
				}
				pricesFromPairs, err := getPriceFromPairsInfo(pairAddress, address, c.ethConn)
				if err != nil {
					tokErrs = append(tokErrs, tokenID)
					log.Warn("error getting price form pairs: ", err, ". Address: ",address, ".TokenId: ", tokenID)
					continue
				}
				float64Value, _ := pricesFromPairs.price.Float64()
				r := make(map[uint]float64)
				r[tokenID] = float64Value
				prices = append(prices, r)
			}
		}
	}
	return prices, tokErrs, nil
}
