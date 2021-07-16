package bbolt

import (
	"errors"

	"github.com/hermeznetwork/price-updater-service/core/domain"
	"github.com/hermeznetwork/price-updater-service/core/ports"
)

type ConfigProviderRepository struct {
	conn *Connection
}

func NewConfigProviderRepository(conn *Connection) ports.ConfigProviderRepository {
	return &ConfigProviderRepository{
		conn: conn,
	}
}

func (cp *ConfigProviderRepository) SaveConfig() error {
	return errors.New("not implemented yet")
}

func (cp *ConfigProviderRepository) LoadConfig() (domain.PriceProvider, error) {
	var pp domain.PriceProvider
	tx, err := cp.conn.db.Begin(true)
	if err != nil {
		return pp, err
	}
	defer tx.Rollback()

	bucket, err := tx.CreateBucketIfNotExists([]byte("configProvider"))
	if err != nil {
		return pp, err
	}

	v := bucket.Get([]byte("coingecko"))
	if len(v) == 0 {
		return pp, errors.New("provider not found")
	}
	return pp, err
}
