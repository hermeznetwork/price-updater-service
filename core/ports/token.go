package ports

import (
	"context"

	"github.com/hermeznetwork/price-updater-service/core/domain"
)

type TokenRepository interface {
	GetTokens(ctx context.Context) ([]domain.Token, error)
	UpdateTokenPrice(ctx context.Context, tokenID uint, value float64) error
}
