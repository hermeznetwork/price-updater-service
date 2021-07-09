package postgres

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hermeznetwork/price-updater-service/core/domain"
	"github.com/hermeznetwork/price-updater-service/core/ports"
)

type tokenFromDB struct {
	Symbol  string
	Price   float64
	Address common.Address
}

type TokenRepository struct {
	conn *Connection
}

func NewTokenRepository(conn *Connection) ports.TokenRepository {
	return &TokenRepository{
		conn: conn,
	}
}

func (t *TokenRepository) GetTokens(ctx context.Context) ([]domain.Token, error) {
	var tokens []domain.Token
	stmt, err := t.conn.db.QueryContext(ctx, "SELECT symbol, usd, eth_addr FROM token")
	if err != nil {
		return tokens, err
	}

	for stmt.Next() {
		tdb := tokenFromDB{}
		err = stmt.Scan(&tdb.Symbol, &tdb.Price, &tdb.Address)
		if err != nil {
			return tokens, err
		}
		tokens = append(tokens, dbToDomainToken(tdb))
	}
	return tokens, nil
}

func (t *TokenRepository) UpdateTokenPrice(ctx context.Context, tokenID uint, value float64) error {
	stmt, err := t.conn.db.PrepareContext(ctx, "UPDATE token SET usd = $1 WHERE token_id = $2")
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(ctx, value, tokenID)
	return err
}

func dbToDomainToken(tdb tokenFromDB) domain.Token {
	return domain.Token{
		Symbol:  tdb.Symbol,
		Price:   tdb.Price,
		Address: tdb.Address.String(),
	}
}
