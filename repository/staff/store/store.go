package store

import (
	"context"

	"github.com/devit-tel/goerror"
	domain "github.com/devit-tel/gogo-blueprint/domain/staff"
	repoStaff "github.com/devit-tel/gogo-blueprint/repository/staff"
	"github.com/devit-tel/jaegerstart"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

type Store struct {
	mongo          *mongo.Client
	dbName         string
	collectionName string
}

func (s *Store) collectionStaff() *mongo.Collection {
	return s.mongo.Database(s.dbName).Collection(s.collectionName)
}

func (s *Store) Get(ctx context.Context, staffId string) (*domain.Staff, goerror.Error) {
	if span := jaegerstart.StartNewSpan(ctx, "REPO_STAFF_Get"); span != nil {
		defer span.Finish()
	}

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
	if span := jaegerstart.StartNewSpan(ctx, "REPO_STAFF_Save"); span != nil {
		defer span.Finish()
	}

	_, err := s.collectionStaff().InsertOne(ctx, staff)
	if err != nil {
		return repoStaff.ErrUnableSaveStaff.WithInput(staff).WithCause(err)
	}

	return nil
}

func (s *Store) GetStaffsByCompany(ctx context.Context, companyId string, offset, limit int64) ([]*domain.Staff, goerror.Error) {
	if span := jaegerstart.StartNewSpan(ctx, "REPO_STAFF_GetStaffsByCompany"); span != nil {
		defer span.Finish()
	}

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
