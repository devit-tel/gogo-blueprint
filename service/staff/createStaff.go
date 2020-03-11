package staff

import (
	"context"

	"github.com/devit-tel/goerror"

	domainStaff "github.com/devit-tel/gogo-blueprint/domain/staff"
)

type CreateStaffInput struct {
	Name      string
	CompanyId string
	Tel       string
}

func (service *StaffService) CreateStaff(ctx context.Context, input *CreateStaffInput) (*domainStaff.Staff, goerror.Error) {
	_, err := service.companyRepository.Get(ctx, input.CompanyId)
	if err != nil {
		return nil, err
	}

	newStaff := domainStaff.Create(service.xid.Gen(), input.CompanyId, input.Name, input.Tel)

	if err := service.staffRepository.Save(ctx, newStaff); err != nil {
		return nil, err
	}

	return newStaff, nil
}
