package staff

import (
	"context"

	domainStaff "github.com/devit-tel/gogo-blueprint/domain/staff"
)

func (s *staffSuite) TestStaffService_UpdateStaff() {
	oldStaff := &domainStaff.Staff{Id: "xxx_1", CompanyId: "comp_1", Name: "Tester", Tel: "081-555-2222", CreatedAt: s.now.Unix(), UpdatedAt: s.now.Unix()}
	expectedStaff := &domainStaff.Staff{Id: "xxx_1", CompanyId: "comp_1", Name: "Tester2", Tel: "081-555-3333", CreatedAt: s.now.Unix(), UpdatedAt: s.now.Unix()}

	s.staffRepository.On("Get", context.Background(), "xxx_1").Once().Return(oldStaff, nil)
	s.staffRepository.On("Save", context.Background(), expectedStaff).Once().Return(nil)

	newStaff, err := s.service.UpdateStaff(context.Background(), &UpdateStaffInput{StaffId: "xxx_1", Name: "Tester2", Tel: "081-555-3333"})
	s.staffRepository.AssertExpectations(s.T())
	s.NoError(err)
	s.Equal(expectedStaff, newStaff)
}
