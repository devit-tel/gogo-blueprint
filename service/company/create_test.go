package company

import (
	"context"

	domainCompany "github.com/devit-tel/gogo-blueprint/domain/company"
)

func (suite *companyService) TestCompanyService_CreateCompany() {
	expectedCompany := &domainCompany.Company{Id: "xxx_1", Name: "NextComp"}

	suite.xid.Freeze("xxx_1")
	suite.companyRepository.On("Save", context.Background(), expectedCompany).Once().Return(nil)

	newCompany, err := suite.service.CreateCompany(context.Background(), &CreateCompanyInput{Name: "NextComp"})
	suite.companyRepository.AssertExpectations(suite.T())
	suite.NoError(err)
	suite.Equal(expectedCompany, newCompany)
}
