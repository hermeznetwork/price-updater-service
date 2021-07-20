package bbolt

import (
	"github.com/hermeznetwork/price-updater-service/config"
	bolt "go.etcd.io/bbolt"
)

type Connection struct {
	db *bolt.DB
}

func NewConnection(cfg config.BboltConfig) *Connection {
	// FIXME: This way to give permission not works.
	db, err := bolt.Open(cfg.Location, 0644, nil)
	if err != nil {
		panic(err.Error())
	}
	return &Connection{db: db}
}
