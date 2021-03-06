package bbolt

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

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

func (cp *ProviderConfigRepository) ChangePriority(priority string) error {
	cp.conn.Start()
	defer cp.conn.End()
	tx, err := cp.conn.db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	bkt, err := tx.CreateBucketIfNotExists([]byte("priorityProvider"))
	if err != nil {
		return err
	}

	err = bkt.Put([]byte("priority"), []byte(priority))
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (cp *ProviderConfigRepository) PriorityProviders() ([]string, error) {
	cp.conn.Start()
	defer cp.conn.End()
	tx, err := cp.conn.db.Begin(true)
	if err != nil {
		return []string{}, err
	}

	bkt, err := tx.CreateBucketIfNotExists([]byte("priorityProvider"))
	if err != nil {
		return []string{}, err
	}

	priority := bkt.Get([]byte("priority"))
	if priority == nil {
		return []string{}, errors.New("there no priority defined")
	}
	priorityArr := strings.Split(string(priority), ",")

	return priorityArr, tx.Commit()
}
