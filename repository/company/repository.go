package company

import (
	"context"

	"github.com/devit-tel/goerror"

	"github.com/devit-tel/gogo-blueprint/domain/company"
)

var (
	ErrCompanyNotFound   = goerror.DefineNotFound("CompanyNotFound", "company not found")
	ErrUnableGetCompany  = goerror.DefineNotFound("UnableGetCompany", "unable to get company")
	ErrUnableSaveCompany = goerror.DefineNotFound("UnableSaveCompany", "unable to save company")
)

//go:generate mockery -name=Repository
type Repository interface {
	Save(ctx context.Context, company *company.Company) goerror.Error
	Get(ctx context.Context, companyId string) (*company.Company, goerror.Error)
}
