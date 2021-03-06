package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/hermeznetwork/price-updater-service/core/domain"
	"github.com/hermeznetwork/price-updater-service/core/ports"
)

type FiatPricesRepository struct {
	conn *Connection
}

func NewFiatPricesRepository(conn *Connection) ports.FiatPriceRepository {
	return &FiatPricesRepository{
		conn: conn,
	}
}

func (f *FiatPricesRepository) GetFiatPrice(ctx context.Context, currency string) (domain.FiatPrice, error) {
	var fp domain.FiatPrice
	err := f.conn.DB.Get(&fp, "SELECT currency, base_currency, COALESCE(price,0), COALESCE(last_update,'1970-01-01 00:00:00') FROM fiat WHERE currency = $1;", currency)
	if err == sql.ErrNoRows {
		return fp, nil
	}
	if err != nil {
		return fp, err
	}
	return fp, nil
}

func (f *FiatPricesRepository) GetFiatPrices(ctx context.Context, baseCurrency string, limit uint, fromItem uint, order string, symbolFilters []string) ([]domain.FiatPrice, error) {
	var fiatPrices []domain.FiatPrice
	var args []interface{}
	pgPosition := 2
	args = append(args, baseCurrency)
	query := "SELECT currency, base_currency, COALESCE(price,0) as price, COALESCE(last_update,'1970-01-01 00:00:00') as last_update FROM fiat WHERE base_currency = $1 "

	// TODO: change to sqlx to avoid this workarround
	if len(symbolFilters) > 0 {
		var q []string
		for _ = range symbolFilters {
			q = append(q, fmt.Sprintf("$%d", pgPosition))
			pgPosition += 1
		}
		query += fmt.Sprintf("AND currency IN (%s) ", strings.Join(q, ","))
		for _, symbol := range symbolFilters {
			args = append(args, symbol)
		}
	}

	if order == "ASC" {
		query += fmt.Sprintf("AND item_id >= $%d ", pgPosition)
	} else {
		query += fmt.Sprintf("AND item_id <= $%d ", pgPosition)
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
	err := f.conn.DB.Select(&fiatPrices, query, args...)
	return fiatPrices, err
}

func (f *FiatPricesRepository) CreateFiatPrice(ctx context.Context, base_currency, currency string, price float64) error {
	query := "INSERT INTO fiat(currency, base_currency, price) VALUES ($1, $2, $3)"
	_, err := f.conn.DB.Exec(query, currency, base_currency, price)
	return err
}

func (f *FiatPricesRepository) UpdateFiatPrice(ctx context.Context, base_currency, currency string, price float64) error {
	query := "UPDATE fiat SET price = $1 WHERE currency = $2 AND base_currency = $3"
	_, err := f.conn.DB.Exec(query, price, currency, base_currency)
	return err
}
