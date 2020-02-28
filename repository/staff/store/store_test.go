// +build integration

package store

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	domain "github.com/devit-tel/gogo-blueprint/domain/staff"
	repoStaff "github.com/devit-tel/gogo-blueprint/repository/staff"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/devit-tel/gogo-blueprint/config"
)

func setup() *Store {
	appConfig := config.Get()

	s := New(appConfig.MongoDBEndpoint, appConfig.MongoDBName, appConfig.MongoDBStaffTableName)
	_, err := s.collectionStaff().DeleteMany(context.Background(), bson.D{})
	if err != nil {
		panic(err)
	}

	return s
}

func (s *Store) withInitStaffs() *Store {
	expectedStaffs := []*domain.Staff{
		{Id: "staff_1", CompanyId: "company_1", Name: "tester_1", Tel: "088-999-8888", CreatedAt: 50, UpdatedAt: 100},
		{Id: "staff_2", CompanyId: "company_1", Name: "tester_2", Tel: "081-333-4444", CreatedAt: 50, UpdatedAt: 100},
		{Id: "staff_3", CompanyId: "company_2", Name: "tester_3", Tel: "056-666-7777", CreatedAt: 50, UpdatedAt: 100},
		{Id: "staff_4", CompanyId: "company_3", Name: "tester_4", Tel: "082-111-6666", CreatedAt: 50, UpdatedAt: 100},
	}

	for _, staff := range expectedStaffs {
		err := s.Save(context.Background(), staff)
		if err != nil {
			panic(err)
		}
	}
	return s
}

func TestStore_SaveAndGet(t *testing.T) {
	s := setup()

	expectedStaff := &domain.Staff{Id: "staff_1", CompanyId: "company_1", Name: "tester", Tel: "088-999-8888", CreatedAt: 50, UpdatedAt: 100}

	err := s.Save(context.Background(), expectedStaff)
	require.NoError(t, err)

	staff, err := s.Get(context.Background(), expectedStaff.Id)
	require.NoError(t, err)
	require.Equal(t, expectedStaff, staff)
}

func TestStore_GetNotFound(t *testing.T) {
	s := setup()

	company, err := s.Get(context.Background(), "notfound_id")
	require.Nil(t, company)
	require.Error(t, err)
	require.True(t, repoStaff.ErrStaffNotFound.IsCodeEqual(err))
}

func TestStore_GetStaffsByCompany(t *testing.T) {
	s := setup().withInitStaffs()

	staffs, err := s.GetStaffsByCompany(context.Background(), "company_1", 0, 10)
	require.NoError(t, err)
	require.Len(t, staffs, 2)
	require.Equal(t, []*domain.Staff{
		{Id: "staff_1", CompanyId: "company_1", Name: "tester_1", Tel: "088-999-8888", CreatedAt: 50, UpdatedAt: 100},
		{Id: "staff_2", CompanyId: "company_1", Name: "tester_2", Tel: "081-333-4444", CreatedAt: 50, UpdatedAt: 100},
	}, staffs)

	staffs, err = s.GetStaffsByCompany(context.Background(), "company_1", 1, 1)
	require.NoError(t, err)
	require.Len(t, staffs, 1)
	require.Equal(t, []*domain.Staff{
		{Id: "staff_2", CompanyId: "company_1", Name: "tester_2", Tel: "081-333-4444", CreatedAt: 50, UpdatedAt: 100},
	}, staffs)
}

func TestStore_GetStaffsByCompany_EmptyResult(t *testing.T) {
	s := setup()

	staffs, err := s.GetStaffsByCompany(context.Background(), "company_notfound", 0, 10)
	require.NoError(t, err)
	require.Len(t, staffs, 0)
	require.Len(t, []*domain.Staff{}, 0)
}
