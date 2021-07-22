package bbolt

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hermeznetwork/price-updater-service/core/domain"
	"github.com/hermeznetwork/price-updater-service/core/ports"
)

type ProviderConfigRepository struct {
	conn *Connection
}

type bboltPriceProvider struct {
	Provider string          `json:"provider"`
	Config   map[uint]string `json:"config"`
}

func NewProviderConfigRepository(conn *Connection) ports.ConfigProviderRepository {
	return &ProviderConfigRepository{
		conn: conn,
	}
}

func (cp *ProviderConfigRepository) CurrentProvider() (string, error) {
	cp.conn.Start()
	defer cp.conn.End()
	tx, err := cp.conn.db.Begin(true)
	if err != nil {
		return "", err
	}

	bkt, err := tx.CreateBucketIfNotExists([]byte("runningProvider"))
	if err != nil {
		return "", err
	}

	currentProvider := bkt.Get([]byte("running"))
	if currentProvider == nil {
		return "", errors.New("there no provider running")
	}
	return string(currentProvider), tx.Commit()
}

func (cp *ProviderConfigRepository) SaveConfig(provider string, data domain.PriceProvider) error {
	cp.conn.Start()
	defer cp.conn.End()
	tx, err := cp.conn.db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	model := bboltPriceProvider{
		Provider: data.Provider,
		Config:   data.MappingBetweenNetwork,
	}

	bkt, err := tx.CreateBucketIfNotExists([]byte("configProvider"))
	if err != nil {
		return err
	}
	buf, err := json.Marshal(model)
	if err != nil {
		return err
	}
	err = bkt.Put([]byte(provider), []byte(buf))
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (cp *ProviderConfigRepository) LoadConfig(provider string) (domain.PriceProvider, error) {
	cp.conn.Start()
	defer cp.conn.End()
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

	providerConfigRaw := bucket.Get([]byte(provider))
	if providerConfigRaw == nil {
		return pp, fmt.Errorf("configuration to provider %s not found", provider)
	}
	var bpp bboltPriceProvider
	err = json.Unmarshal(providerConfigRaw, &bpp)
	if err != nil {
		return pp, err
	}
	pp.Provider = bpp.Provider
	pp.MappingBetweenNetwork = bpp.Config
	return pp, tx.Commit()
}

func (cp *ProviderConfigRepository) ChangeRunningProvider(provider string) error {
	cp.conn.Start()
	defer cp.conn.End()
	tx, err := cp.conn.db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	bkt, err := tx.CreateBucketIfNotExists([]byte("runningProvider"))
	if err != nil {
		return err
	}

	err = bkt.Put([]byte("running"), []byte(provider))
	if err != nil {
		return err
	}

	return tx.Commit()
}
