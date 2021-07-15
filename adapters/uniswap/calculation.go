package uniswap

import (
	"errors"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hermeznetwork/price-updater-service/adapters/uniswap/artefact"
	"github.com/hermeznetwork/price-updater-service/adapters/uniswap/pair"
)

type pairsInfo struct {
	symbol0  string
	symbol1  string
	address0 common.Address
	addresl1 common.Address
	price    *big.Float
}

func getPriceFromPairsInfo(pairAddress, currentToken common.Address, conn *ethclient.Client) (*pairsInfo, error) {
	pairToken, err := pair.NewToken(pairAddress, conn)
	if err != nil {
		return nil, err
	}

	addressToken0, err := pairToken.Token0(nil)
	if err != nil {
		return nil, err
	}
	addressToken1, err := pairToken.Token1(nil)
	if err != nil {
		return nil, err
	}

	reserveInfo, err := pairToken.GetReserves(nil)
	if err != nil {
		return nil, err
	}

	reserve0 := reserveInfo.Reserve0
	reserve1 := reserveInfo.Reserve1

	token0, err := artefact.NewToken(addressToken0, conn)
	if err != nil {
		return nil, err
	}

	token1, err := artefact.NewToken(addressToken1, conn)
	if err != nil {
		return nil, err
	}

	tokenDecimals0, err := token0.Decimals(nil)
	if err != nil {
		return nil, err
	}

	tokenDecimals1, err := token1.Decimals(nil)
	if err != nil {
		return nil, err
	}

	tokenSymbol0, err := token0.Symbol(nil)
	if err != nil {
		return nil, err
	}

	tokenSymbol1, err := token1.Symbol(nil)
	if err != nil {
		return nil, err
	}

	var tokenDecimalsA uint8
	var tokenDecimalsB uint8
	var tokenReserveA big.Int
	var tokenReserveB big.Int
	var tokenSymbolA string
	var tokenSymbolB string

	tokenDecimalsA = tokenDecimals0
	tokenDecimalsB = tokenDecimals1
	tokenReserveA = *reserve0
	tokenReserveB = *reserve1
	tokenSymbolA = tokenSymbol0
	tokenSymbolB = tokenSymbol1

	if currentToken != addressToken1 {
		tokenDecimalsA = tokenDecimals1
		tokenDecimalsB = tokenDecimals0
		tokenReserveA = *reserve1
		tokenReserveB = *reserve0
		tokenSymbolA = tokenSymbol1
		tokenSymbolB = tokenSymbol0
	}

	amountIn := new(big.Int)
	expTokenA := big.NewInt(int64(tokenDecimalsA))
	amountIn.Exp(big.NewInt(10), expTokenA, nil)
	amountOut, err := getAmountOut(amountIn, tokenReserveA, tokenReserveB)
	if err != nil {
		return nil, err
	}

	amountInF := new(big.Float)
	amountInF.SetInt(amountIn)
	amountOutF := new(big.Float)
	amountOutF.SetInt(amountOut)

	finalAmountIn := new(big.Float)
	finalAmountOut := new(big.Float)

	floatDecimalsA := big.NewFloat(math.Pow10(int(tokenDecimalsA)))
	floatDecimalsB := big.NewFloat(math.Pow10(int(tokenDecimalsB)))

	finalAmountIn.Quo(amountInF, floatDecimalsA)
	finalAmountOut.Quo(amountOutF, floatDecimalsB)

	price := new(big.Float).Quo(finalAmountIn, finalAmountOut)

	result := pairsInfo{
		symbol0:  tokenSymbolA,
		symbol1:  tokenSymbolB,
		address0: addressToken0,
		addresl1: addressToken1,
		price:    price,
	}
	return &result, nil
}

func getAmountOut(amountInt *big.Int, reserveIn big.Int, reserveOut big.Int) (*big.Int, error) {
	zeroReturn := new(big.Int)
	if amountInt.Int64() <= 0 {
		return zeroReturn, errors.New("UniswapV2Library: INSUFFICIENT_INPUT_AMOUNT")
	}

	if reserveIn.Int64() <= 0 && reserveOut.Int64() <= 0 {
		return zeroReturn, errors.New("UniswapV2Library: INSUFFICIENT_LIQUIDITY")
	}

	amountInWithFee := new(big.Int).Mul(amountInt, big.NewInt(997))
	numerator := new(big.Int).Mul(amountInWithFee, &reserveOut)
	denominator := new(big.Int).Add(new(big.Int).Mul(&reserveIn, big.NewInt(1000)), amountInWithFee)

	return new(big.Int).Quo(numerator, denominator), nil
}
