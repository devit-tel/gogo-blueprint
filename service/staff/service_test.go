package staff

import (
	"testing"
	"time"

	"github.com/devit-tel/gotime"
	"github.com/devit-tel/goxid"
	"github.com/stretchr/testify/suite"

	mockCompany "github.com/devit-tel/gogo-blueprint/repository/company/mocks"
	mockStaff "github.com/devit-tel/gogo-blueprint/repository/staff/mocks"
)

type staffService struct {
	suite.Suite
	companyRepository *mockCompany.Repository
	staffRepository   *mockStaff.Repository
	xid               *goxid.ID
	service           Service
	now               time.Time
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(staffService))
}

func (suite *staffService) SetupTest() {
	suite.xid = goxid.New()
	suite.staffRepository = &mockStaff.Repository{}
	suite.companyRepository = &mockCompany.Repository{}

	now := time.Now()
	gotime.Freeze(now)
	suite.now = now

	suite.service = New(suite.xid, suite.staffRepository, suite.companyRepository)
}
