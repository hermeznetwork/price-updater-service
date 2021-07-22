package memory

import (
	"testing"

	"github.com/hermeznetwork/price-updater-service/core/domain"
)

func TestLoadToMemory(t *testing.T) {
	newMemory := NewMemoryDB()

	tokens := make(map[uint]domain.Token)
	tokenID := uint(1)
	tokens[tokenID] = domain.Token{
		ItemID:    2,
		ID:        1,
		Price:     9.9,
		Symbol:    "",
		Address:   "",
		BlockNum:  0,
		Name:      "",
		Decimals:  0,
		UsdUpdate: "",
	}
	newMemory.LoadTokensPrices(tokens)

	token, err := newMemory.GetTokenPrice(tokenID)
	if err != nil {
		t.Error(err.Error())
	}

	if token != tokens[tokenID] {
		t.Error("test fail: the tokens are different")
	}

	tokenID = uint(5)
	_, err = newMemory.GetTokenPrice(tokenID)
	if err == nil {
		t.Error("test fail: expected fail because this tokenID doesn't exist", err.Error())
	}
}
