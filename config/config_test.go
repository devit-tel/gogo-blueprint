package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	appConfig := Get()
	require.Equal(t, "mongodb://localhost:27017", appConfig.MongoDBEndpoint)
	require.Equal(t, "company_test", appConfig.MongoDBCompanyTableName)
	require.Equal(t, "staff_test", appConfig.MongoDBStaffTableName)
}
