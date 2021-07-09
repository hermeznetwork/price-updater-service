package mongo

import (
	"context"
	"fmt"

	"github.com/hermeznetwork/price-updater-service/config"
	mdb "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(ctx context.Context, cfg *config.Config) *mdb.Database {
	url := fmt.Sprintf("mongodb+srv://%s:%s\\@%s/%s?retryWrites=true&w=majority", cfg.Mongo.User, cfg.Mongo.Password, cfg.Mongo.Host, cfg.Mongo.Database)
	client, err := mdb.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		panic(fmt.Errorf("can't create mongo client: %s", err.Error()))
	}

	err = client.Connect(ctx)
	if err != nil {
		panic(fmt.Errorf("can't connect mongo db: %s", err.Error()))
	}
	return client.Database(cfg.Mongo.Database)
}
