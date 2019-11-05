package staff

import (
    "context"

    domainStaff "github.com/devit-tel/gogo-blueprint/domain/staff"
)

func (suite *staffService) TestStaffService_UpdateStaff() {
    oldStaff := &domainStaff.Staff{Id: "xxx_1", CompanyId: "comp_1", Name: "Tester", Tel: "081-555-2222", CreatedAt: suite.now.Unix(), UpdatedAt: suite.now.Unix()}
    expectedStaff := &domainStaff.Staff{Id: "xxx_1", CompanyId: "comp_1", Name: "Tester2", Tel: "081-555-3333", CreatedAt: suite.now.Unix(), UpdatedAt: suite.now.Unix()}

    suite.staffRepository.On("Get", context.Background(), "xxx_1").Once().Return(oldStaff, nil)
    suite.staffRepository.On("Save", context.Background(), expectedStaff).Once().Return(nil)

    newStaff, err := suite.service.UpdateStaff(context.Background(), &UpdateStaffInput{StaffId: "xxx_1", Name: "Tester2", Tel: "081-555-3333"})
    suite.staffRepository.AssertExpectations(suite.T())
    suite.NoError(err)
    suite.Equal(expectedStaff, newStaff)
}
