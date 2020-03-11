package company

import (
	"context"

	"github.com/devit-tel/goerror"
	"github.com/devit-tel/goxid"

	domainCompany "github.com/devit-tel/gogo-blueprint/domain/company"
	"github.com/devit-tel/gogo-blueprint/repository/company"
)

//go:generate mockery -name=Service
type Service interface {
	CreateCompany(ctx context.Context, input *CreateCompanyInput) (*domainCompany.Company, goerror.Error)
}

type CompanyService struct {
	companyRepository company.Repository
	xid               *goxid.ID
}

func New(xid *goxid.ID, c company.Repository) *CompanyService {
	return &CompanyService{
		companyRepository: c,
		xid:               xid,
	}
}
