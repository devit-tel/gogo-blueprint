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

type staffSuite struct {
	suite.Suite
	companyRepository *mockCompany.Repository
	staffRepository   *mockStaff.Repository
	xid               *goxid.ID
	service           Service
	now               time.Time
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(staffSuite))
}

func (s *staffSuite) SetupTest() {
	s.xid = goxid.New()
	s.staffRepository = &mockStaff.Repository{}
	s.companyRepository = &mockCompany.Repository{}

	now := time.Now()
	gotime.Freeze(now)
	s.now = now

	s.service = New(s.xid, s.staffRepository, s.companyRepository)
}
