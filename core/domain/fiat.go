package domain

type FiatPrice struct {
	Currency     string
	BaseCurrency string
	Price        float64
	LastUpdate   string
}
