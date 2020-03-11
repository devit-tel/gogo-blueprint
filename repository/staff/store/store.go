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

	domain "github.com/devit-tel/gogo-blueprint/domain/staff"
	repoStaff "github.com/devit-tel/gogo-blueprint/repository/staff"
)

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

type Store struct {
	db             *mongo.Client
	dbName         string
	collectionName string
}

func (s *Store) collectionStaff() *mongo.Collection {
	return s.db.Database(s.dbName).Collection(s.collectionName)
}

func (s *Store) Get(ctx context.Context, staffId string) (*domain.Staff, goerror.Error) {
	staff := &domain.Staff{}
	if err := s.collectionStaff().FindOne(ctx, bson.D{{"_id", staffId}}).Decode(staff); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, repoStaff.ErrStaffNotFound.WithInput(staffId)
		}

		return nil, repoStaff.ErrUnableGetStaff.WithInput(staffId).WithCause(err)
	}

	return staff, nil
}

func (s *Store) Save(ctx context.Context, staff *domain.Staff) goerror.Error {
	_, err := s.collectionStaff().InsertOne(ctx, staff)
	if err != nil {
		return repoStaff.ErrUnableSaveStaff.WithInput(staff).WithCause(err)
	}

	return nil
}

func (s *Store) GetStaffsByCompany(ctx context.Context, companyId string, offset, limit int64) ([]*domain.Staff, goerror.Error) {
	cursor, err := s.collectionStaff().Find(ctx, bson.M{"companyId": companyId}, options.Find().SetLimit(limit).SetSkip(offset))
	if err != nil {
		return nil, repoStaff.ErrUnableGetStaffs.WithInput(companyId).WithCause(err)
	}
	defer cursor.Close(ctx)

	staffs := make([]*domain.Staff, 0)
	for cursor.Next(ctx) {
		staff := &domain.Staff{}
		if err := cursor.Decode(staff); err != nil {
			return nil, repoStaff.ErrUnableGetStaffs.WithInput(cursor.Current).WithCause(err)
		}

		staffs = append(staffs, staff)
	}

	return staffs, nil
}
