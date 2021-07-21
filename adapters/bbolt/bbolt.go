package bbolt

import (
	"github.com/hermeznetwork/price-updater-service/config"
	bolt "go.etcd.io/bbolt"
)

const defaultPermission = 0644

type Connection struct {
	db  *bolt.DB
	cfg config.BboltConfig
}

func NewConnection(cfg config.BboltConfig) *Connection {
	return &Connection{db: nil, cfg: cfg}
}

func (c *Connection) Start() {
	db, err := bolt.Open(c.cfg.Location, defaultPermission, nil)
	c.db = db
	if err != nil {
		panic(err.Error())
	}
}

func (c *Connection) End() {
	c.db.Close()
}
