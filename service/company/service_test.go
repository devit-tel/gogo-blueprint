package company

import (
    "testing"

    "github.com/devit-tel/gogo-blueprint/repository/company/mocks"
    "github.com/devit-tel/goxid"
    "github.com/stretchr/testify/suite"
)

type companyService struct {
    suite.Suite
    companyRepository *mocks.Repository
    xid               *goxid.ID
    service           Service
}

func TestRunSuite(t *testing.T) {
    suite.Run(t, new(companyService))
}

func (suite *companyService) SetupTest() {
    suite.xid = goxid.New()
    suite.companyRepository = &mocks.Repository{}
    suite.service = New(suite.xid, suite.companyRepository)
}
