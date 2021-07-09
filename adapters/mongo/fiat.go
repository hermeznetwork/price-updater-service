package mongo

import (
	"context"

	"github.com/hermeznetwork/price-updater-service/core/domain"
	"github.com/hermeznetwork/price-updater-service/core/ports"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mdb "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FiatPrices struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Currency     string             `bson:"currency" json:"currency"`
	BaseCurrency string             `bson:"base_currency" json:"base_currency"`
	Price        float64            `bson:"price" json:"price"`
}

type FiatPriceRepository struct {
	collection *mdb.Collection
}

func NewFiatPriceRepository(db *mdb.Database) ports.FiatPriceRepository {
	return &FiatPriceRepository{
		collection: db.Collection("fiat_prices"),
	}
}

func (f *FiatPriceRepository) GetFiatPrices(ctx context.Context, currency string) ([]domain.FiatPrice, error) {
	var prices []domain.FiatPrice
	cur, err := f.collection.Find(ctx, bson.M{"currency": currency})
	if err != nil {
		return prices, err
	}

	for cur.Next(ctx) {
		fpDB := FiatPrices{}
		if err := cur.Decode(&fpDB); err != nil {
			return prices, err
		}
		prices = append(prices, f.dbToDomainFiatPrice(fpDB))
	}
	return prices, nil
}

func (f *FiatPriceRepository) CreateFiatPrice(ctx context.Context, base_currency, currency string, price float64) error {
	searchBy := bson.M{
		"currency": currency,
	}
	fiatPriceSet := bson.M{
		"$set": bson.M{
			"currency":      currency,
			"base_currency": base_currency,
			"price":         price,
		},
	}

	_, err := f.collection.UpdateOne(ctx, searchBy, fiatPriceSet, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}
	return nil
}

func (f *FiatPriceRepository) UpdateFiatPrice(ctx context.Context, base_currency, currency string, price float64) error {
	searchBy := bson.M{
		"currency": currency,
	}
	fiatPriceSet := bson.M{
		"$set": bson.M{
			"currency":      currency,
			"base_currency": base_currency,
			"price":         price,
		},
	}

	_, err := f.collection.UpdateOne(ctx, searchBy, fiatPriceSet, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}
	return nil
}

func (f *FiatPriceRepository) dbToDomainFiatPrice(fp FiatPrices) domain.FiatPrice {
	return domain.FiatPrice{
		Currency:     fp.Currency,
		BaseCurrency: fp.BaseCurrency,
		Price:        fp.Price,
	}
}
