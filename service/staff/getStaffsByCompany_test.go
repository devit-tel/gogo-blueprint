package staff

import (
	"context"

	domainStaff "github.com/devit-tel/gogo-blueprint/domain/staff"
)

func (s *staffSuite) createStaffs() []*domainStaff.Staff {
	return []*domainStaff.Staff{
		{Id: "xxx_1", CompanyId: "comp_1", Name: "Tester1", Tel: "081-555-1111", CreatedAt: s.now.Unix(), UpdatedAt: s.now.Unix()},
		{Id: "xxx_2", CompanyId: "comp_2", Name: "Tester2", Tel: "081-555-2222", CreatedAt: s.now.Unix(), UpdatedAt: s.now.Unix()},
		{Id: "xxx_3", CompanyId: "comp_3", Name: "Tester3", Tel: "081-555-333", CreatedAt: s.now.Unix(), UpdatedAt: s.now.Unix()},
	}
}

func (s *staffSuite) TestStaffService_GetStaffsByCompany() {
	expectedStaffs := s.createStaffs()

	s.staffRepository.On("GetStaffsByCompany", context.Background(), "comp_1", int64(0), int64(30)).Once().Return(expectedStaffs, nil)

	staffs, err := s.service.GetStaffsByCompany(context.Background(), &GetStaffsByCompanyInput{CompanyId: "comp_1", Limit: 30, Offset: 0})
	s.staffRepository.AssertExpectations(s.T())
	s.NoError(err)
	s.Equal(expectedStaffs, staffs)
}
