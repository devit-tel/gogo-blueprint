package staff

type GetStaffsByCompanyInput struct {
	CompanyId string `json:"companyId" form:"companyId" binding:"required"`
	Limit     int64  `json:"limit,default=20" form:"limit"`
	Offset    int64  `json:"offset" form:"offset"`
}

type GetStaffsByCompanyOutput struct {
	Staffs []*Staff `json:"staffs"`
}
