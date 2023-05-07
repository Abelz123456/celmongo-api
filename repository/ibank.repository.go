package repository

import (
	"context"

	"github.com/celmonggo-api/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type IBankRepository interface {
	Save(ctx context.Context, c *mongo.Collection, bank entity.Bank) (*mongo.InsertOneResult, error)
	Update(ctx context.Context, c *mongo.Collection, bank entity.Bank) entity.Bank
	Delete(ctx context.Context, c *mongo.Collection, bank entity.Bank)
	FindById(ctx context.Context, c *mongo.Collection, oid string) (entity.Bank, error)
	Find(ctx context.Context, c *mongo.Collection, criteria string) (*mongo.Cursor, error)
}
