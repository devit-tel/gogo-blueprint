package store

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/devit-tel/goerror"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	domain "github.com/devit-tel/gogo-blueprint/domain/company"
	repoCompany "github.com/devit-tel/gogo-blueprint/repository/company"
)

type Store struct {
	db             *mongo.Client
	dbName         string
	collectionName string
}

func New(mongoEndpoint, dbName, collectionName string) *Store {
	clientOptions := options.Client().ApplyURI(mongoEndpoint)

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	db, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	return &Store{
		dbName:         dbName,
		collectionName: collectionName,
		db:             db,
	}
}

func (s *Store) collectionCompany() *mongo.Collection {
	return s.db.Database(s.dbName).Collection(s.collectionName)
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
