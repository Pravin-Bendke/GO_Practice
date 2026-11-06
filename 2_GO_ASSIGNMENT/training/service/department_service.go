package domain

import (
	"context"

	apperrors "training/errors"
	"training/models"
	"training/repository"
)

type DepartmentService struct {
	Repository repository.DepartmentRepository
}

func NewDepartmentService(
	repo repository.DepartmentRepository,
) *DepartmentService {

	return &DepartmentService{
		Repository: repo,
	}
}

func (s *DepartmentService) GetDepartment(
	ctx context.Context,
	id int64,
	companyID int64,
) (*models.Department, error) {

	if id <= 0 {
		return nil, apperrors.ErrInvalidDepartmentID
	}

	if companyID <= 0 {
		return nil, apperrors.ErrInvalidCompany
	}

	department, err := s.Repository.GetByID(
		ctx,
		id,
		companyID,
	)

	if err != nil {
		return nil, apperrors.MapRepositoryError(err)
	}

	return department, nil
}

func (s *DepartmentService) ListDepartments(
	ctx context.Context,
	companyID int64,
	filters repository.QueryFilters,
) ([]models.Department, int64, error) {

	if companyID <= 0 {
		return nil, 0, apperrors.ErrInvalidCompany
	}

	if filters.Page < 1 ||
		filters.PageSize < 1 ||
		filters.PageSize > 100 {
		return nil, 0, apperrors.ErrInvalidPagination
	}

	departments, total, err := s.Repository.ListByCompanyID(
		ctx,
		companyID,
		filters,
	)

	if err != nil {
		return nil, 0, apperrors.MapRepositoryError(err)
	}

	return departments, total, nil
}
