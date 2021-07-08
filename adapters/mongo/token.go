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

type Token struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	TokenID int                `bson:"token_id" json:"tokenID"`
	Price   float64            `bson:"usd" json:"usd"`
	Symbol  string             `bson:"symbol" json:"symbol"`
	Address string             `bson:"eth_addr" json:"eth_add"`
}

type TokenRepository struct {
	collection *mdb.Collection
}

func NewTokenRepository(db *mdb.Database) ports.TokenRepository {
	return &TokenRepository{
		collection: db.Collection("tokens"),
	}
}

func (t *TokenRepository) GetTokens(ctx context.Context) ([]domain.Token, error) {
	var tokens []domain.Token
	cur, err := t.collection.Find(ctx, bson.M{})
	if err != nil {
		return tokens, err
	}

	for cur.Next(ctx) {
		tkDB := Token{}
		if err := cur.Decode(&tkDB); err != nil {
			return tokens, err
		}
		tokens = append(tokens, t.dbToDomainToken(tkDB))
	}
	return tokens, nil
}

func (t *TokenRepository) UpdateTokenPrice(ctx context.Context, tokenID int, value float64) error {
	searchBy := bson.M{
		"token_id": tokenID,
	}
	newPriceValue := bson.M{
		"$set": bson.M{
			"usd": value,
		},
	}
	t.collection.UpdateOne(ctx, searchBy, newPriceValue, options.Update())
	return nil
}

func (t *TokenRepository) dbToDomainToken(token Token) domain.Token {
	return domain.Token{
		TokenID: token.TokenID,
		Price:   token.Price,
		Symbol:  token.Symbol,
		Address: token.Address,
	}
}
