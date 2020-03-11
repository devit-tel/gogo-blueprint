package app

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"

	mockCompnay "github.com/devit-tel/gogo-blueprint/service/company/mocks"
	mockStaff "github.com/devit-tel/gogo-blueprint/service/staff/mocks"
)

type AppTestSuite struct {
	suite.Suite
	staffService   *mockStaff.Service
	companyService *mockCompnay.Service

	app    *App
	router *gin.Engine
}

func (s *AppTestSuite) SetupTest() {
	s.staffService = &mockStaff.Service{}
	s.companyService = &mockCompnay.Service{}

	app := New(s.staffService, s.companyService)

	gin.SetMode("release")
	g := gin.New()

	app.RegisterRoute(g)

	s.app = app
	s.router = g
}

func TestAppSuite(t *testing.T) {
	suite.Run(t, new(AppTestSuite))
}
