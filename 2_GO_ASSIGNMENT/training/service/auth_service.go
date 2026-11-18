package domain

import (
	"context"
	"database/sql"
	"errors"
	apperrors "training/errors"
	"training/models"
	"training/repository"
	"training/shared/password"
)

type AuthService struct {
	Repository repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) *AuthService {
	return &AuthService{
		Repository: repo,
	}
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
