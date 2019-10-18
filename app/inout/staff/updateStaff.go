package staff

type UpdateStaffInput struct {
	Id   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Tel  string `json:"tel" binding:"required"`
}

type UpdateStaffOutput struct {
	Staff *Staff `json:"staff"`
}
