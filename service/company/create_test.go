package company

import (
	"context"

	domainCompany "github.com/devit-tel/gogo-blueprint/domain/company"
)

func (s *companySuite) TestCompanyService_CreateCompany() {
	expectedCompany := &domainCompany.Company{Id: "xxx_1", Name: "NextComp"}

	s.xid.Freeze("xxx_1")
	s.companyRepository.On("Save", context.Background(), expectedCompany).Once().Return(nil)

	newCompany, err := s.service.CreateCompany(context.Background(), &CreateCompanyInput{Name: "NextComp"})
	s.companyRepository.AssertExpectations(s.T())
	s.NoError(err)
	s.Equal(expectedCompany, newCompany)
}
