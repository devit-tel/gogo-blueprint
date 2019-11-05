package company

import (
    domainCompany "github.com/devit-tel/gogo-blueprint/domain/company"
)

type Company struct {
    Id   string `json:"id"`
    Name string `json:"name"`
}

func ToCompanyOutput(company *domainCompany.Company) *Company {
    return &Company{
        Id:   company.Id,
        Name: company.Name,
    }
}
