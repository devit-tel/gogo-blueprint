package company

import (
	"context"

	"github.com/devit-tel/goerror"

	domainCompany "github.com/devit-tel/gogo-blueprint/domain/company"
)

type CreateCompanyInput struct {
	Name string
}

func (service *CompanyService) CreateCompany(ctx context.Context, input *CreateCompanyInput) (*domainCompany.Company, goerror.Error) {
	newCompany := domainCompany.Create(service.xid.Gen(), input.Name)

	if err := service.companyRepository.Save(ctx, newCompany); err != nil {
		return nil, err
	}

	return newCompany, nil
}
