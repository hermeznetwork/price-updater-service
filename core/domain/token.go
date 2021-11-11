package domain

import 	"github.com/ethereum/go-ethereum/common"

type Token struct {
	ItemID    uint           `db:"item_id"`
	ID        uint           `db:"token_id"`
	Price     float64        `db:"usd"`
	Symbol    string         `db:"symbol"`
	Address   common.Address `db:"eth_addr"`
	BlockNum  uint           `db:"eth_block_num"`
	Name      string         `db:"name"`
	Decimals  uint           `db:"decimals"`
	UsdUpdate string         `db:"usd_update"`
}
