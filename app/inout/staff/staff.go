package staff

import (
    domainStaff "github.com/devit-tel/gogo-blueprint/domain/staff"
)

type Staff struct {
    Id        string `json:"id"`
    Name      string `json:"name"`
    CompanyId string `json:"companyId"`
    Tel       string `json:"tel"`
}

func ToStaffOutput(staff *domainStaff.Staff) *Staff {
    return &Staff{
        Id:        staff.Id,
        Name:      staff.Name,
        CompanyId: staff.CompanyId,
        Tel:       staff.Tel,
    }
}

func ToStaffsOutput(staffs []*domainStaff.Staff) []*Staff {
    outputStaffs := make([]*Staff, len(staffs))

    for index, staff := range staffs {
        outputStaffs[index] = ToStaffOutput(staff)
    }

    return outputStaffs
}
