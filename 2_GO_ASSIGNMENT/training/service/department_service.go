package domain

import (
	"context"
	"errors"
	"log"
	"strings"

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

func (s *DepartmentService) CreateDepartment(
	ctx context.Context,
	department *models.Department,
) (*models.Department, error) {
	if department == nil {
		return nil, errors.New("department required")
	}
	if department.CompanyID <= 0 {
		return nil, apperrors.ErrInvalidCompany
	}

	department.Code = strings.TrimSpace(department.Code)
	department.Name = strings.TrimSpace(department.Name)
	department.Description = strings.TrimSpace(department.Description)

	if department.Code == "" {
		return nil, errors.New("department code is required")
	}

	if department.Name == "" {
		return nil, errors.New("department is required")
	}

	created, err := s.Repository.Create(
		ctx,
		department,
	)

	if err != nil {
		return nil, apperrors.MapRepositoryError(err)
	}
	return created, nil

}

func (s *DepartmentService) UpdateDepartment(
	ctx context.Context,
	id int64,
	companyID int64,
	req models.UpdateDepartmentRequest,
) (*models.Department, error) {

	if id <= 0 {
		return nil, apperrors.ErrInvalidDepartmentID
	}

	if companyID <= 0 {
		return nil, apperrors.ErrInvalidCompany
	}

	req.Code = strings.TrimSpace(req.Code)
	req.Name = strings.TrimSpace(req.Name)
	req.Description = strings.TrimSpace(req.Description)

	if req.Code == "" {
		return nil, errors.New("department code is required")
	}

	if req.Name == "" {
		return nil, errors.New("department name is required")
	}
	department := &models.Department{
		ID:          id,
		CompanyID:   companyID,
		Code:        req.Code,
		Name:        req.Name,
		Description: req.Description,
	}
	updated, err := s.Repository.Update(ctx, department)

	if err != nil {
		return nil, apperrors.MapRepositoryError(err)
	}
	return updated, nil
}

func (s *DepartmentService) UpdateStatus(
	ctx context.Context,
	id int64,
	companyID int64,
	isActive bool,
) error {
	if id <= 0 {
		return apperrors.ErrInvalidDepartmentID
	}

	if companyID <= 0 {
		return apperrors.ErrInvalidCompany
	}

	err := s.Repository.UpdateStatus(
		ctx,
		id,
		companyID,
		isActive,
	)

	if err != nil {
		return apperrors.MapRepositoryError(err)
	}
	return nil
}

func (s *DepartmentService) DeleteDepartment(
	ctx context.Context,
	id int64,
	companyID int64,
) error {
	if id <= 0 {
		return apperrors.ErrInvalidDepartmentID
	}

	if companyID <= 0 {
		return apperrors.ErrInvalidCompany
	}

	log.Printf(
		"DeleteDepartment called: id=%d companyID=%d",
		id,
		companyID,
	)

	hasActiveUsers, err := s.Repository.HasActiveUser(
		ctx,
		id,
		companyID,
	)

	if err != nil {
		return apperrors.MapRepositoryError(err)
	}

	log.Printf(
		"Active users found: %t",
		hasActiveUsers,
	)

	if hasActiveUsers {
		log.Println("Department deletion blocked")
		return apperrors.ErrDepartmentHasActiveUsers
	}

	err = s.Repository.Delete(
		ctx,
		id,
		companyID,
	)

	if err != nil {
		return apperrors.MapRepositoryError(err)
	}
	return nil
}
