package store

import (
	"context"

	"github.com/devit-tel/goerror"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	domain "github.com/devit-tel/gogo-blueprint/domain/company"
	repoCompany "github.com/devit-tel/gogo-blueprint/repository/company"
)

type Store struct {
	mongo          *mongo.Client
	dbName         string
	collectionName string
}

func New(mongoEndpoint, dbName, collectionName string) *Store {
	clientOptions := options.Client().ApplyURI(mongoEndpoint)

	db, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}

	return &Store{
		dbName:         dbName,
		collectionName: collectionName,
		mongo:          db,
	}
}

func (s *Store) collectionCompany() *mongo.Collection {
	return s.mongo.Database(s.dbName).Collection(s.collectionName)
}

func (s *Store) Get(ctx context.Context, companyId string) (*domain.Company, goerror.Error) {
	company := &domain.Company{}
	if err := s.collectionCompany().FindOne(ctx, bson.D{{"_id", companyId}}).Decode(company); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, repoCompany.ErrCompanyNotFound.WithInput(companyId)
		}

		return nil, repoCompany.ErrUnableGetCompany.WithInput(companyId).WithCause(err)
	}

	return company, nil
}

func (s *Store) Save(ctx context.Context, company *domain.Company) goerror.Error {
	_, err := s.collectionCompany().InsertOne(ctx, company)
	if err != nil {
		return repoCompany.ErrUnableSaveCompany.WithInput(company).WithCause(err)
	}

	return nil
}
