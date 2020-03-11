package company

import (
	"testing"

	"github.com/devit-tel/goxid"
	"github.com/stretchr/testify/suite"

	"github.com/devit-tel/gogo-blueprint/repository/company/mocks"
)

type companySuite struct {
	suite.Suite
	companyRepository *mocks.Repository
	xid               *goxid.ID
	service           Service
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(companySuite))
}

func (s *companySuite) SetupTest() {
	s.xid = goxid.New()
	s.companyRepository = &mocks.Repository{}
	s.service = New(s.xid, s.companyRepository)
}
