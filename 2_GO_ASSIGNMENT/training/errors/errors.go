package errors

import (
	"database/sql"
	"errors"
)

var (
	ErrDepartmentNotFound = errors.New(
		"department not found",
	)

	ErrInvalidCompany = errors.New(
		"invalid company id",
	)

	ErrInvalidDepartmentID = errors.New(
		"invalid department id",
	)

	ErrInvalidPagination = errors.New(
		"invalid pagination parameters",
	)

	ErrDatabase = errors.New(
		"database error",
	)

	//Authentication error
	ErrInvalidCredentials = errors.New(
		"invalid email or password",
	)

	ErrEmailAlreadyExists = errors.New(
		"email already exists",
	)

	ErrInvalidRegistration = errors.New(
		"invalid registration data",
	)
	ErrDepartmentHasActiveUsers = errors.New(
		"cannot delete department because active users are assigned",
	)
)

func MapRepositoryError(err error) error {

	if errors.Is(err, sql.ErrNoRows) {
		return ErrDepartmentNotFound
	}

	return ErrDatabase
}
