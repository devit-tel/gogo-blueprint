package staff

import (
    "context"

    domainCompany "github.com/devit-tel/gogo-blueprint/domain/company"
    domainStaff "github.com/devit-tel/gogo-blueprint/domain/staff"
)

func (suite *staffService) TestStaffService_CreateStaff() {
    expectedStaff := &domainStaff.Staff{Id: "xxx_1", CompanyId: "comp_1", Name: "Tester", Tel: "081-555-2222", CreatedAt: suite.now.Unix(), UpdatedAt: suite.now.Unix()}

    suite.xid.Freeze("xxx_1")
    suite.companyRepository.On("Get", context.Background(), "comp_1").Once().Return(&domainCompany.Company{Id: "comp_1", Name: "NextComp"}, nil)
    suite.staffRepository.On("Save", context.Background(), expectedStaff).Once().Return(nil)

    newStaff, err := suite.service.CreateStaff(context.Background(), &CreateStaffInput{Name: "Tester", CompanyId: "comp_1", Tel: "081-555-2222"})
    suite.companyRepository.AssertExpectations(suite.T())
    suite.staffRepository.AssertExpectations(suite.T())
    suite.NoError(err)
    suite.Equal(expectedStaff, newStaff)
}
