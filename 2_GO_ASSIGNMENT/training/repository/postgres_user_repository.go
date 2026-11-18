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
