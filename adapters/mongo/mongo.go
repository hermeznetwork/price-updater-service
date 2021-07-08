package mongo

import (
	"context"
	"fmt"

	"github.com/hermeznetwork/price-updater-service/config"
	mdb "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(ctx context.Context, cfg *config.Config) *mdb.Database {
	client, err := mdb.NewClient(options.Client().ApplyURI(fmt.Sprintf("http://%s:%d", cfg.Mongo.Host, cfg.Mongo.Port)))
	if err != nil {
		panic(fmt.Errorf("can't create mongo client: %s", err.Error()))
	}

	err = client.Connect(ctx)
	if err != nil {
		panic(fmt.Errorf("can't connect mongo db: %s", err.Error()))
	}
	return client.Database(cfg.Mongo.Database)
}
