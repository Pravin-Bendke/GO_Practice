package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

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
func (r *PostgresDepartmentRepository) Create(
	ctx context.Context,
	department *models.Department,
) (*models.Department, error) {
	query := `
	INSERT INTO training_department (
	company_id,
	code,
	name,
	description
	)
	VALUES ($1,$2,$3,$4)
	RETURNING
	id,
	company_id,
	code,
	name,
	COALESCE(description,''),
	is_active,
	created_at,
	updated_at 
	`
	var created models.Department

	err := r.DB.QueryRowContext(
		ctx,
		query,
		department.CompanyID,
		department.Code,
		department.Name,
		department.Description,
	).Scan(
		&created.ID,
		&created.CompanyID,
		&created.Code,
		&created.Name,
		&created.Description,
		&created.IsActive,
		&created.CreatedAt,
		&created.UpdatedAt,
	)
	if err != nil {
		log.Printf("Department repository Create failed : %v", err)
		return nil, err
	}
	return &created, nil
}

func (r *PostgresDepartmentRepository) Update(
	ctx context.Context,
	department *models.Department,
) (*models.Department, error) {
	query :=
		`UPDATE training_department
	SET 
	code =$1,
	name =$2,
	description =$3,
	updated_at = CURRENT_TIMESTAMP
	WHERE id = $4
	AND company_id = $5
	AND deleted_at IS NULL
	RETURNING
	id,
	company_id,
	code,
	name,
	COALESCE(description,''),
	is_active,
	created_at,
	updated_at
	`

	var updated models.Department

	err := r.DB.QueryRowContext(
		ctx,
		query,
		department.Code,
		department.Name,
		department.Description,
		department.ID,
		department.CompanyID,
	).Scan(
		&updated.ID,
		&updated.CompanyID,
		&updated.Code,
		&updated.Name,
		&updated.Description,
		&updated.IsActive,
		&updated.CreatedAt,
		&updated.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &updated, nil
}

// Delete
func (r *PostgresDepartmentRepository) Delete(
	ctx context.Context,
	id int64,
	companyID int64,
) error {
	query := `
	UPDATE training_department
	SET deleted_at = CURRENT_TIMESTAMP
	WHERE id = $1
	AND company_id =$2
	AND deleted_at IS NULL
	`
	result, err := r.DB.ExecContext(
		ctx,
		query,
		id,
		companyID,
	)
	if err != nil {
		log.Printf("Department repository Delete failed: %v: ", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
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

		query += ` AND (name ILIKE $2 OR code ILIKE $2)`      //we use for find ENG - Engineering
		countQuery += ` AND (name ILIKE $2 OR code ILIKE $2)` //we are use for after search found total

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

	//this is allow sorting fields.
	allowedSort := map[string]bool{
		"id":         true,
		"name":       true,
		"code":       true,
		"created_at": true,
	}

	// this filters option we already define in ListByCompanyID because we are applying pagination sorting on list not on single record
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

	//QueryContext We can use to get multiple rows data.
	//QueryRowContext We can use to get single row data.
	rows, err := r.DB.QueryContext(
		ctx,
		query,
		args...,
	)

	if err != nil {
		return nil, 0, err
	}
	//we need to close row after get data/row
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
		//Adding in the department structure
		departments = append(departments, department)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return departments, total, nil
}
