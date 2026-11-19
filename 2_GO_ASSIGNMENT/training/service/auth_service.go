package domain

import (
	"context"
	"database/sql"
	"errors"
	apperrors "training/errors"
	"training/models"
	"training/repository"
	"training/shared/password"
	shared "training/shared/validator"
)

type AuthService struct {
	Repository repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) *AuthService {
	return &AuthService{
		Repository: repo,
	}
}

// User Registration
func (s *AuthService) Register(
	ctx context.Context,
	req models.RegisterRequest,
) error {

	// LOGIC FOR VALIDATION REGISTER FIELDS :
	if err := shared.ValidateFirstName(req.FirstName); err != nil {
		return err
	}

	if err := shared.ValidateLastName(req.LastName); err != nil {
		return err
	}

	if err := shared.ValidateEmail(req.Email); err != nil {
		return err
	}

	if err := shared.ValidatePassword(req.Password); err != nil {
		return err
	}

	existingUser, err := s.Repository.GetByEmail(ctx, req.Email)

	if err == nil && existingUser != nil {
		return apperrors.ErrEmailAlreadyExists
	}

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return apperrors.ErrDatabase
	}

	hashedPassword, err := password.HashPassword(req.Password)
	if err != nil {
		return apperrors.ErrDatabase
	}

	user := models.User{
		CompanyID: req.CompanyID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  hashedPassword,
	}

	err = s.Repository.CreateWithDepartment(
		ctx,
		&user,
		req.DepartmentCode,
	)

	if err != nil {
		return err
	}
	return nil
}

func (s *AuthService) Authenticate(
	ctx context.Context,
	email string,
	plainPassword string,
) (*models.User, error) {

	user, err := s.Repository.GetByEmail(
		ctx,
		email,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.ErrInvalidCredentials
		}
		return nil, apperrors.ErrDatabase
	}
	if !user.IsActive {
		return nil, apperrors.ErrInvalidCredentials
	}
	isValid := password.CheckPassword(
		user.Password,
		plainPassword,
	)
	if !isValid {
		return nil, apperrors.ErrInvalidCredentials
	}
	return user, nil
}

func (s *AuthService) GetCurrentUser(
	ctx context.Context,
	UserID int64,
	companyID int64,
) (*models.User, error) {
	user, err := s.Repository.GetByID(
		ctx,
		UserID,
		companyID,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.ErrInvalidCredentials
		}
		return nil, apperrors.ErrDatabase
	}
	return user, nil
}
