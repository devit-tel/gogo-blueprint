package staff

import (
	"context"

	"github.com/devit-tel/goerror"

	domainStaff "github.com/devit-tel/gogo-blueprint/domain/staff"
)

type GetStaffsByCompanyInput struct {
	CompanyId string
	Offset    int64
	Limit     int64
}

func (service *StaffService) GetStaffsByCompany(ctx context.Context, input *GetStaffsByCompanyInput) ([]*domainStaff.Staff, goerror.Error) {
	staffs, err := service.staffRepository.GetStaffsByCompany(ctx, input.CompanyId, input.Offset, input.Limit)
	if err != nil {
		return nil, err
	}

	return staffs, nil
}
