package domain

type PriceProvider struct {
	Provider string
	// Mapping Between Network it necessary because of the providers implemented's work's well just in mainnet network.
	// This attribute will map the tokenID saved on hermez database to addresses running on current network.
	MappingBetweenNetwork map[uint]string
}
