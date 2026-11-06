package repository

import (
	"context"
	"database/sql"
	"fmt"

	"training/models"
)

type PostgresDepartmentRepository struct {
	DB *sql.DB
}

func NewDepartmentRepository(
	db *sql.DB,
) DepartmentRepository {
	return &PostgresDepartmentRepository{
		DB: db,
	}
}
func (r *PostgresDepartmentRepository) GetByID(
	ctx context.Context,
	id int64,
	companyID int64,
) (*models.Department, error) {

	query := `
        SELECT
            id,
            company_id,
            code,
            name,
            description,
            is_active,
            created_at,
            updated_at
        FROM training_department
        WHERE id = $1
        AND company_id = $2
        AND deleted_at IS NULL
    `

	var department models.Department

	err := r.DB.QueryRowContext(
		ctx,
		query,
		id,
		companyID,
	).Scan(
		&department.ID,
		&department.CompanyID,
		&department.Code,
		&department.Name,
		&department.Description,
		&department.IsActive,
		&department.CreatedAt,
		&department.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &department, nil
}
func (r *PostgresDepartmentRepository) ListByCompanyID(
	ctx context.Context,
	companyID int64,
	filters QueryFilters,
) ([]models.Department, int64, error) {

	query := `
        SELECT
            id,
            company_id,
            code,
            name,
            description,
            is_active,
            created_at,
            updated_at
        FROM training_department
        WHERE company_id = $1
        AND deleted_at IS NULL
    `

	countQuery := `
        SELECT COUNT(*)
        FROM training_department
        WHERE company_id = $1
        AND deleted_at IS NULL
    `

	args := []interface{}{companyID}

	if filters.Search != "" {

		search := "%" + filters.Search + "%"

		query += ` AND (name ILIKE $2 OR code ILIKE $2)`
		countQuery += ` AND (name ILIKE $2 OR code ILIKE $2)`

		args = append(args, search)
	}

	var total int64

	err := r.DB.QueryRowContext(
		ctx,
		countQuery,
		args...,
	).Scan(&total)

	if err != nil {
		return nil, 0, err
	}

	allowedSort := map[string]bool{
		"id":         true,
		"name":       true,
		"code":       true,
		"created_at": true,
	}

	sortBy := filters.SortBy

	if !allowedSort[sortBy] {
		sortBy = "id"
	}

	sortOrder := "ASC"

	if filters.SortOrder == "desc" {
		sortOrder = "DESC"
	}

	query += fmt.Sprintf(
		" ORDER BY %s %s LIMIT $%d OFFSET $%d",
		sortBy,
		sortOrder,
		len(args)+1,
		len(args)+2,
	)

	offset := (filters.Page - 1) * filters.PageSize

	args = append(args, filters.PageSize, offset)

	rows, err := r.DB.QueryContext(
		ctx,
		query,
		args...,
	)

	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()

	departments := []models.Department{}

	for rows.Next() {

		var department models.Department

		err := rows.Scan(
			&department.ID,
			&department.CompanyID,
			&department.Code,
			&department.Name,
			&department.Description,
			&department.IsActive,
			&department.CreatedAt,
			&department.UpdatedAt,
		)

		if err != nil {
			return nil, 0, err
		}

		departments = append(departments, department)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return departments, total, nil
}
