package repository

import (
	"context"
	"training/models"
)

// Interface म्हणजे कोणत्या methods उपलब्ध असाव्यात याचा contract.
type UserRepository interface {
	GetByEmail(
		ctx context.Context, //Context वापरून Database Query ला cancellation आणि timeout support देता येतो.
		email string,
	) (*models.User, error)
	GetByID(
		ctx context.Context,
		id int64,
		companyID int64,
	) (*models.User, error)
	CreateWithDepartment(
		ctx context.Context,
		user *models.User,
		departmentCode string,
	) error
}
