package domain

type PriceProvider struct {
	Provider       string
	BaseURL        string
	URL            string
	URLExtraParams string
	//SymbolsMap     symbolsMap
	//AddressesMap   addressesMap
	Symbols   string
	Addresses string
}
