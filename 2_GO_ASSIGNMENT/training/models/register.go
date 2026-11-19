package models

type RegisterRequest struct {
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	DepartmentCode string `json:"department_code"`
	CompanyID      int64  `json:"company_id"`
}
