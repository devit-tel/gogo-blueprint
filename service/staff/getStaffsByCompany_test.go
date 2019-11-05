package staff

import (
    "context"

    domainStaff "github.com/devit-tel/gogo-blueprint/domain/staff"
)

func (suite *staffService) createStaffs() []*domainStaff.Staff {
    return []*domainStaff.Staff{
        {Id: "xxx_1", CompanyId: "comp_1", Name: "Tester1", Tel: "081-555-1111", CreatedAt: suite.now.Unix(), UpdatedAt: suite.now.Unix()},
        {Id: "xxx_2", CompanyId: "comp_2", Name: "Tester2", Tel: "081-555-2222", CreatedAt: suite.now.Unix(), UpdatedAt: suite.now.Unix()},
        {Id: "xxx_3", CompanyId: "comp_3", Name: "Tester3", Tel: "081-555-333", CreatedAt: suite.now.Unix(), UpdatedAt: suite.now.Unix()},
    }
}

func (suite *staffService) TestStaffService_GetStaffsByCompany() {
    expectedStaffs := suite.createStaffs()

    suite.staffRepository.On("GetStaffsByCompany", context.Background(), "comp_1", int64(0), int64(30)).Once().Return(expectedStaffs, nil)

    staffs, err := suite.service.GetStaffsByCompany(context.Background(), &GetStaffsByCompanyInput{CompanyId: "comp_1", Limit: 30, Offset: 0})
    suite.staffRepository.AssertExpectations(suite.T())
    suite.NoError(err)
    suite.Equal(expectedStaffs, staffs)
}
