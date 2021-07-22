package ports

import "github.com/hermeznetwork/price-updater-service/core/domain"

type MemoryRepository interface {
	LoadTokensPrices(mp map[uint]domain.Token) error
	LoadFiatPrices(mp map[string]domain.FiatPrice) error
	GetTokensPrices() ([]domain.Token, error)
	GetFiatPrices() ([]domain.FiatPrice, error)
	GetTokenPrice(tokenID uint) (domain.Token, error)
	GetFiatPrice(currency string) (domain.FiatPrice, error)
	GetAllowOriginValues() (string, error)
	LoadAllowOriginValues(originValues string) error
}
