package repository

import (
	"context"
	"training/models"
)

type QueryFilters struct {
	Page      int
	PageSize  int
	Search    string
	SortBy    string
	SortOrder string
}

type DepartmentRepository interface {
	Create(
		ctx context.Context,
		department *models.Department,
	) (*models.Department, error)

	Update(
		ctx context.Context,
		department *models.Department,
	) (*models.Department, error)

	Delete(
		ctx context.Context,
		id int64,
		companyID int64,
	) error

	GetByID(
		ctx context.Context,
		id int64,
		companyID int64,
	) (*models.Department, error)

	ListByCompanyID(
		ctx context.Context,
		companyID int64,
		filters QueryFilters,
	) ([]models.Department, int64, error)
}
