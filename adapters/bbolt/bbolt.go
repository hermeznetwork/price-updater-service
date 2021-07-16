package bbolt

import (
	"github.com/hermeznetwork/price-updater-service/config"
	bolt "go.etcd.io/bbolt"
)

type Connection struct {
	db *bolt.DB
}

func NewConnection(cfg config.BboltConfig) *Connection {
	db, err := bolt.Open(cfg.Location, cfg.Permission, nil)
	if err != nil {
		panic(err.Error())
	}
	return &Connection{db: db}
}
