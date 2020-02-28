package staff

import (
	"context"

	"github.com/devit-tel/goerror"

	domainStaff "github.com/devit-tel/gogo-blueprint/domain/staff"
)

type UpdateStaffInput struct {
	StaffId string
	Name    string
	Tel     string
}

func (service *StaffService) UpdateStaff(ctx context.Context, input *UpdateStaffInput) (*domainStaff.Staff, goerror.Error) {
	staff, err := service.staffRepository.Get(ctx, input.StaffId)
	if err != nil {
		return nil, err
	}

	staff.Update(input.Name, input.Tel)

	if err := service.staffRepository.Save(ctx, staff); err != nil {
		return nil, err
	}

	return staff, nil
}
