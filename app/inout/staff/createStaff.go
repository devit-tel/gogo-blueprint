package staff

type CreateStaffInput struct {
	Name      string `json:"name" binding:"required"`
	CompanyId string `json:"companyId" binding:"required"`
	Tel       string `json:"tel"`
}

type CreateStaffOutput struct {
	Staff *Staff `json:"staff"`
}
