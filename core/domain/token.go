package domain

type Token struct {
	ItemID    uint
	ID        uint
	Price     float64
	Symbol    string
	Address   string
	BlockNum  uint
	Name      string
	Decimals  uint
	UsdUpdate string
}
