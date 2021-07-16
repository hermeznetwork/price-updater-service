package postgres

import (
	"context"

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
	stmt := f.conn.db.QueryRowContext(ctx, "SELECT currency, base_currency, price FROM fiat WHERE currency = $1;", currency)
	err := stmt.Scan(&fp.Currency, &fp.BaseCurrency, &fp.Price)
	if err != nil {
		return fp, err
	}
	return fp, nil
}

func (f *FiatPricesRepository) GetFiatPrices(ctx context.Context, baseCurrency string) ([]domain.FiatPrice, error) {
	var fiatPrices []domain.FiatPrice
	stmt, err := f.conn.db.Query("SELECT currency, base_currency, price FROM fiat WHERE base_currency = $1;", baseCurrency)
	if err != nil {
		return fiatPrices, err
	}

	for stmt.Next() {
		fp := domain.FiatPrice{}
		err = stmt.Scan(&fp.Currency, &fp.BaseCurrency, &fp.Price)
		if err != nil {
			return fiatPrices, err
		}
		fiatPrices = append(fiatPrices, fp)
	}

	return fiatPrices, nil
}

func (f *FiatPricesRepository) CreateFiatPrice(ctx context.Context, base_currency, currency string, price float64) error {
	stmt, err := f.conn.db.PrepareContext(ctx, "INSERT INTO fiat(currency, base_currency, price) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(ctx, currency, base_currency, price)
	return err
}

func (f *FiatPricesRepository) UpdateFiatPrice(ctx context.Context, base_currency, currency string, price float64) error {
	stmt, err := f.conn.db.PrepareContext(ctx, "UPDATE fiat SET price = $1 WHERE currency = $2 AND base_currency = $3")
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(ctx, price, currency, base_currency)
	return err
}
