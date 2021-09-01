package ports

import (
	"context"

	"github.com/hermeznetwork/price-updater-service/core/domain"
)

type TokenRepository interface {
	GetToken(ctx context.Context, tokenID uint) (domain.Token, error)
	GetTokens(ctx context.Context, fromItem uint, limit uint, order string, ids []string) ([]domain.Token, error)
	UpdateTokenPrice(ctx context.Context, tokenID uint, value float64) error
}
