package app

import (
	"testing"

	mockCompnay "github.com/devit-tel/gogo-blueprint/service/company/mocks"
	mockStaff "github.com/devit-tel/gogo-blueprint/service/staff/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type AppTestSuite struct {
	suite.Suite
	staffService   *mockStaff.Service
	companyService *mockCompnay.Service

	app    *App
	router *gin.Engine
}

func (suite *AppTestSuite) SetupTest() {
	suite.staffService = &mockStaff.Service{}
	suite.companyService = &mockCompnay.Service{}

	app := New(suite.staffService, suite.companyService)

	gin.SetMode("release")
	g := gin.New()

	app.RegisterRoute(g)

	suite.app = app
	suite.router = g
}

func TestAppSuite(t *testing.T) {
	suite.Run(t, new(AppTestSuite))
}
