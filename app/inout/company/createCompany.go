package company

type CreateCompanyInput struct {
	Name string `json:"name" binding:"required"`
}

type CreateCompanyOutput struct {
	Company *Company `json:"company"`
}
