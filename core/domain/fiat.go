package domain

type FiatPrice struct {
	Currency     string  `db:"currency"`
	BaseCurrency string  `db:"base_currency"`
	Price        float64 `db:"price"`
	LastUpdate   string  `db:"last_update"`
}
