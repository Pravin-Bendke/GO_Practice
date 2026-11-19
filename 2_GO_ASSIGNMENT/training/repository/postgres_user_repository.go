package repository

import (
	"context"
	"database/sql"
	"training/models"
)

type PostgresUserRepository struct {
	DB *sql.DB
}

func NewUserRepository(
	db *sql.DB,
) UserRepository {
	return &PostgresUserRepository{
		DB: db,
	}
}

func (r *PostgresUserRepository) GetByEmail(
	ctx context.Context,
	email string,
) (*models.User, error) {
	query := `SELECT 
	id,
	company_id,
	first_name,
	last_name,
	email,
	password,
	is_active,
	created_at,
	updated_at,
	deleted_at
	FROM training_user
	WHERE email = $1
	AND deleted_at IS NULL
	`
	var user models.User

	err := r.DB.QueryRowContext(
		ctx,
		query,
		email,
	).Scan(
		&user.ID,
		&user.CompanyID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *PostgresUserRepository) GetByID(
	ctx context.Context,
	id int64,
	companyID int64,
) (*models.User, error) {
	query := `
	SELECT 
	id,
	company_id,
	first_name,
	last_name,
	email,
	is_active,
	created_at,
	updated_at,
	deleted_at
	FROM training_user
	WHERE id = $1
	AND company_id =$2
	AND is_active = TRUE
	AND deleted_at IS NULL
	`
	var user models.User

	err := r.DB.QueryRowContext(
		ctx,
		query,
		id,
		companyID,
	).Scan(
		&user.ID,
		&user.CompanyID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *PostgresUserRepository) CreateWithDepartment(
	ctx context.Context,
	user *models.User,
	departmentCode string,
) error {
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var departmentID int64

	// find active department
	if departmentCode != "" {
		query := `
		SELECT id 
		FROM training_department
		WHERE code = $1
		AND company_id = $2
		AND is_active = TRUE
		AND deleted_at IS NULL
		`
		err = tx.QueryRowContext(
			ctx,
			query,
			departmentCode,
			user.CompanyID,
		).Scan(&departmentID)
		if err != nil {
			return err
		}
	}

	//Create user
	userQuery := `
	INSERT INTO training_user (
	company_id,
	first_name,
	last_name,
	email,
	password,
	is_active,
	created_at,
	updated_at
	)
	VALUES ($1,$2,$3,$4,$5,TRUE,CURRENT_TIMESTAMP,CURRENT_TIMESTAMP)
	RETURNING id
	`

	err = tx.QueryRowContext(
		ctx,
		userQuery,
		user.CompanyID,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
	).Scan(&user.ID)

	if err != nil {
		return err
	}

	//Create user_department mapping

	if departmentCode != "" {
		mappingQuery := `
		INSERT INTO training_user_department (
		user_id,
		department_id
		)
		VALUES ($1,$2)
		`
		_, err = tx.ExecContext(
			ctx,
			mappingQuery,
			user.ID,
			departmentID,
		)

		if err != nil {
			return err
		}
	}
	return tx.Commit()
}
