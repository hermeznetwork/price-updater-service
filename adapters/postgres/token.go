package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/hermeznetwork/price-updater-service/core/domain"
	"github.com/hermeznetwork/price-updater-service/core/ports"
)

type TokenRepository struct {
	conn *Connection
}

func NewTokenRepository(conn *Connection) ports.TokenRepository {
	return &TokenRepository{
		conn: conn,
	}
}

func (t *TokenRepository) GetToken(ctx context.Context, tokenID uint) (domain.Token, error) {
	var token domain.Token
	err := t.conn.DB.Get(&token, "SELECT item_id, token_id, eth_block_num, decimals, COALESCE(usd_update,'1970-01-01 00:00:00'), name, symbol, COALESCE(usd,0), eth_addr FROM token WHERE token_id  = $1", tokenID)
	if err == sql.ErrNoRows {
		return token, nil
	}
	if err != nil {
		return token, err
	}
	return token, nil
}

func (t *TokenRepository) GetTokens(ctx context.Context, fromItem uint, limit uint, order string, ids []string) ([]domain.Token, error) {
	var tokens []domain.Token
	var args []interface{}
	pgPosition := 1
	query := "SELECT item_id, token_id, eth_block_num, decimals, COALESCE(usd_update,'1970-01-01 00:00:00') as usd_update, name, symbol, COALESCE(usd,0) as usd, eth_addr FROM token WHERE "

	if order == "ASC" {
		query += fmt.Sprintf("item_id >= $%d ", pgPosition)
	} else {
		query += fmt.Sprintf("item_id <= $%d ", pgPosition)
	}
	args = append(args, fromItem)

	// TODO: change to sqlx to avoid this workarround
	if len(ids) > 0 {
		var q []string
		for _ = range ids {
			pgPosition += 1
			q = append(q, fmt.Sprintf("$%d", pgPosition))
		}
		query += fmt.Sprintf("AND token_id IN (%s) ", strings.Join(q, ","))
		for _, tokenID := range ids {
			args = append(args, tokenID)
		}
	}

	// pagination
	query += "ORDER BY item_id "
	if order == "ASC" {
		query += "ASC "
	} else {
		query += "DESC "
	}
	if limit != 0 {
		query += fmt.Sprintf("LIMIT %d;", limit)
	}
	err := t.conn.DB.Select(&tokens, query, args...)
	return tokens, err
}

func (t *TokenRepository) UpdateTokenPrice(ctx context.Context, tokenID uint, value float64) error {
	query := "UPDATE token SET usd = $1 WHERE token_id = $2"
	_, err := t.conn.DB.Exec(query, value, tokenID)
	return err
}
