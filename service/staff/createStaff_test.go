package staff

import (
	"context"

	domainCompany "github.com/devit-tel/gogo-blueprint/domain/company"
	domainStaff "github.com/devit-tel/gogo-blueprint/domain/staff"
)

func (s *staffSuite) TestStaffService_CreateStaff() {
	expectedStaff := &domainStaff.Staff{Id: "xxx_1", CompanyId: "comp_1", Name: "Tester", Tel: "081-555-2222", CreatedAt: s.now.Unix(), UpdatedAt: s.now.Unix()}

	s.xid.Freeze("xxx_1")
	s.companyRepository.On("Get", context.Background(), "comp_1").Once().Return(&domainCompany.Company{Id: "comp_1", Name: "NextComp"}, nil)
	s.staffRepository.On("Save", context.Background(), expectedStaff).Once().Return(nil)

	newStaff, err := s.service.CreateStaff(context.Background(), &CreateStaffInput{Name: "Tester", CompanyId: "comp_1", Tel: "081-555-2222"})
	s.companyRepository.AssertExpectations(s.T())
	s.staffRepository.AssertExpectations(s.T())
	s.NoError(err)
	s.Equal(expectedStaff, newStaff)
}
