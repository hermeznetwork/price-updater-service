package memory

import (
	"errors"
	"fmt"
	"sync"

	"github.com/hermeznetwork/price-updater-service/core/domain"
	"github.com/hermeznetwork/price-updater-service/core/ports"
)

type Memory struct {
	store sync.Map
}

func NewMemoryDB() ports.MemoryRepository {
	var m sync.Map
	return &Memory{
		store: m,
	}
}

func (m *Memory) LoadTokensPrices(mp map[uint]domain.Token) error {
	m.store.Store("tokenPrices", mp)
	return nil
}

func (m *Memory) GetTokensPrices() ([]domain.Token, error) {
	var tokens []domain.Token
	tokenInterface, ok := m.store.Load("tokenPrices")
	if !ok {
		return tokens, errors.New("tokens not found on memory")
	}
	tokenPrices, ok := tokenInterface.(map[string]domain.Token)
	if !ok {
		return tokens, errors.New("failed to load tokens on memory")
	}
	for _, v := range tokenPrices {
		tokens = append(tokens, v)
	}

	return tokens, nil
}

func (m *Memory) GetTokenPrice(tokenID uint) (domain.Token, error) {
	var token domain.Token
	tokensInterface, ok := m.store.Load("tokenPrices")
	if !ok {
		return token, errors.New("token bucket not found on memory")
	}
	tokens, ok := tokensInterface.(map[uint]domain.Token)
	if !ok {
		return token, errors.New(fmt.Sprintf("failed to load token %d on memory", tokenID))
	}
	v, ok := tokens[tokenID]
	if !ok {
		return token, errors.New(fmt.Sprintf("token %d not found on memory", tokenID))
	}
	return v, nil
}

func (m *Memory) LoadFiatPrices(mp map[string]domain.FiatPrice) error {
	m.store.Store("fiatPrices", mp)
	return nil
}

func (m *Memory) GetFiatPrices() ([]domain.FiatPrice, error) {
	var fiats []domain.FiatPrice
	fiatInterface, ok := m.store.Load("fiatPrices")
	if !ok {
		return fiats, errors.New("currencies not found on memory")
	}
	fiatPrices, ok := fiatInterface.(map[string]domain.FiatPrice)
	if !ok {
		return fiats, errors.New("failed to load currencies on memory")
	}
	for _, v := range fiatPrices {
		fiats = append(fiats, v)
	}

	return fiats, nil
}

func (m *Memory) GetFiatPrice(currency string) (domain.FiatPrice, error) {
	var fiat domain.FiatPrice
	fiatInterface, ok := m.store.Load("fiatPrices")
	if !ok {
		return fiat, errors.New(fmt.Sprintf("currency %s not found on memory", currency))
	}
	fiatPrices, ok := fiatInterface.(map[string]domain.FiatPrice)
	if !ok {
		return fiat, errors.New(fmt.Sprintf("failed to load currency %s on memory", currency))
	}

	return fiatPrices[currency], nil
}

func (m *Memory) GetAllowOriginValues() (string, error) {
	originInterface, ok := m.store.Load("origins")
	if !ok {
		return "*", nil
	}
	return originInterface.(string), nil
}

func (m *Memory) LoadAllowOriginValues(originValues string) error {
	m.store.Store("origins", originValues)
	return nil
}
