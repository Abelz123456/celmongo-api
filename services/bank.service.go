package services

import (
	"context"
	"time"

	"github.com/celmonggo-api/common"
	"github.com/celmonggo-api/dto"
	"github.com/celmonggo-api/entity"
	"github.com/celmonggo-api/mapping"
	"github.com/celmonggo-api/repository"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

type BankService struct {
	BankRepository repository.IBankRepository
	database       *mongo.Database
	Validate       *validator.Validate
}

// Create implements IBankService
func (u *BankService) Create(ctx context.Context, request dto.CreateBankDto) mapping.BankVm {
	repo := repository.NewBankRepository()
	c := u.database.Collection("bank")
	bank := entity.Bank{
		BankCode:     &request.BankCode,
		BankName:     &request.BankName,
		UserInserted: &request.UserInserted,
		InsertedDate: time.Now(),
		LastUserId:   new(string),
		LastUpdate:   time.Now(),
	}
	_, err := repo.Save(ctx, c, bank)
	common.PanicIfError(err)
	return mapping.ToBankResponse(bank)
}

// Find implements IBankService
func (u *BankService) Find(ctx context.Context, criteria string) ([]*mapping.BankVm, error) {
	c := u.database.Collection("bank")
	var banks []*mapping.BankVm
	repo := repository.NewBankRepository()
	cursor, err := repo.Find(ctx, c, criteria)

	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var bank mapping.BankVm
		err := cursor.Decode(&bank)
		if err != nil {
			return nil, err
		}

		banks = append(banks, &bank)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(ctx)

	return banks, nil
}

// Find implements IBankService
// func (*BankService) Find(ctx context.Context, collection *mongo.Collection, criteria string) (*mongo.Cursor, error) {

// 	cursor, err := collection.Find(ctx, bson.D{{}})
// 	return cursor, err
// }

func NewBankService(bankRepository repository.IBankRepository, database *mongo.Database, validate *validator.Validate) IBankService {
	return &BankService{
		BankRepository: bankRepository,
		database:       database,
		Validate:       validate,
	}
}
