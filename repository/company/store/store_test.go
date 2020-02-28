// +build integration

package store

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	domain "github.com/devit-tel/gogo-blueprint/domain/company"
	repoCompany "github.com/devit-tel/gogo-blueprint/repository/company"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/devit-tel/gogo-blueprint/config"
)

func setup() *Store {
	appConfig := config.Get()

	s := New(appConfig.MongoDBEndpoint, appConfig.MongoDBName, appConfig.MongoDBCompanyTableName)
	_, err := s.collectionCompany().DeleteMany(context.Background(), bson.D{})
	if err != nil {
		panic(err)
	}

	return s
}

func TestStore_SaveAndGet(t *testing.T) {
	s := setup()

	expectedCompany := &domain.Company{Id: "comp_1", Name: "comp_tester"}

	err := s.Save(context.Background(), expectedCompany)
	require.NoError(t, err)

	company, err := s.Get(context.Background(), expectedCompany.Id)
	require.NoError(t, err)
	require.Equal(t, expectedCompany, company)
}

func TestStore_GetNotFound(t *testing.T) {
	s := setup()

	company, err := s.Get(context.Background(), "notfound_id")
	require.Nil(t, company)
	require.Error(t, err)
	require.True(t, repoCompany.ErrCompanyNotFound.IsCodeEqual(err))
}
