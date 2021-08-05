package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/hermeznetwork/price-updater-service/core/domain"
	"github.com/hermeznetwork/price-updater-service/core/ports"
)

type tokenFromDB struct {
	ItemID    uint
	ID        uint
	BlockNum  uint
	Name      string
	Symbol    string
	Price     float64
	Address   common.Address
	Decimals  uint
	UsdUpdate string
}

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
	stmt := t.conn.db.QueryRowContext(ctx, "SELECT item_id, token_id, eth_block_num, decimals, COALESCE(usd_update,'1970-01-01 00:00:00'), name, symbol, COALESCE(usd,0), eth_addr FROM token WHERE token_id  = $1", tokenID)
	tdb := tokenFromDB{}
	err := stmt.Scan(&tdb.ItemID, &tdb.ID, &tdb.BlockNum, &tdb.Decimals, &tdb.UsdUpdate, &tdb.Name, &tdb.Symbol, &tdb.Price, &tdb.Address)
	if err == sql.ErrNoRows {
		return token, nil
	}
	if err != nil {
		return token, err
	}
	token = dbToDomainToken(tdb)
	return token, nil
}

func (t *TokenRepository) GetTokens(ctx context.Context, fromItem uint, limit uint, order string) ([]domain.Token, error) {
	var tokens []domain.Token
	var args []interface{}
	query := "SELECT item_id, token_id, eth_block_num, decimals, COALESCE(usd_update,'1970-01-01 00:00:00'), name, symbol, COALESCE(usd,0), eth_addr FROM token WHERE "
	if order == "ASC" {
		query += "item_id >= $1 "
	} else {
		query += "item_id <= $1 "
	}
	args = append(args, fromItem)

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
	stmt, err := t.conn.db.QueryContext(ctx, query, args...)
	if err != nil {
		return tokens, err
	}

	for stmt.Next() {
		tdb := tokenFromDB{}
		err = stmt.Scan(&tdb.ItemID, &tdb.ID, &tdb.BlockNum, &tdb.Decimals, &tdb.UsdUpdate, &tdb.Name, &tdb.Symbol, &tdb.Price, &tdb.Address)
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
		ItemID:    tdb.ItemID,
		ID:        tdb.ID,
		Price:     tdb.Price,
		Symbol:    tdb.Symbol,
		Address:   tdb.Address.String(),
		BlockNum:  tdb.BlockNum,
		Name:      tdb.Name,
		Decimals:  tdb.Decimals,
		UsdUpdate: tdb.UsdUpdate,
	}
}
