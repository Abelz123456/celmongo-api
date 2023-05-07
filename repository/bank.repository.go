package repository

import (
	"context"

	"github.com/celmonggo-api/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/google/uuid"
)

type BankRepository struct {
}

// Delete implements IBankRepository
func (*BankRepository) Delete(ctx context.Context, c *mongo.Collection, bank entity.Bank) {
	panic("unimplemented")
}

// Find implements IBankRepository
func (*BankRepository) Find(ctx context.Context, c *mongo.Collection, criteria string) (*mongo.Cursor, error) {
	cursor, err := c.Find(ctx, bson.D{{}})
	return cursor, err
}

// FindById implements IBankRepository
func (*BankRepository) FindById(ctx context.Context, c *mongo.Collection, oid string) (entity.Bank, error) {
	panic("unimplemented")
}

// Update implements IBankRepository
func (*BankRepository) Update(ctx context.Context, c *mongo.Collection, bank entity.Bank) entity.Bank {
	panic("unimplemented")
}

func NewBankRepository() IBankRepository {
	return &BankRepository{}
}

func (repository *BankRepository) Save(ctx context.Context, c *mongo.Collection, bank entity.Bank) (*mongo.InsertOneResult, error) {

	id := uuid.New().String()

	bank.Oid = id
	result, err := c.InsertOne(ctx, bank)
	return result, err

}
