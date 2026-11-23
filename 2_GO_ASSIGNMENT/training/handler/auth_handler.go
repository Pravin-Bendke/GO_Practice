package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"training/config"
	apperrors "training/errors"
	"training/middleware"
	"training/models"
	domain "training/service"
	jwthelper "training/shared/jwt"
	shared "training/shared/validator"
)

type AuthHandler struct {
	Service *domain.AuthService
	Config  config.Config
}

func NewAuthHandler(
	service *domain.AuthService,
	cfg config.Config,
) *AuthHandler {

	return &AuthHandler{
		Service: service,
		Config:  cfg,
	}
}

//login req

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// login res
type LoginResponse struct {
	Message     string    `json:"message"`
	AccessToken string    `json:"access_token"`
	TokenType   string    `json:"token_type"`
	ExpiresIn   int64     `json:"expires_in"`
	User        LoginUser `json:"user"`
}

// Register Response
type RegisterResponse struct {
	Message string `json:"message"`
}

// Login user
type LoginUser struct {
	ID        int64  `json:"id"`
	CompanyID int64  `json:"company_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func (h *AuthHandler) Login(
	w http.ResponseWriter,
	r *http.Request,
) {

	//decode json
	var request LoginRequest

	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&request)

	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
		return
	}
	request.Email = strings.TrimSpace(request.Email)

	if err := shared.ValidateEmail(request.Email); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	user, err := h.Service.Authenticate(
		r.Context(),
		request.Email,
		request.Password,
	)

	if err != nil {
		if errors.Is(err, apperrors.ErrInvalidCredentials) {
			writeJSON(w, http.StatusUnauthorized, map[string]string{
				"error": "invalid email or password",
			})
			return
		}
		writeJSON(w, http.StatusInternalServerError, map[string]string{
			"error": "internal server error",
		})
		return
	}
	token, err := jwthelper.GenerateToken(
		h.Config,
		user.ID,
		user.CompanyID,
		user.Email,
		user.Role,
	)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{
			"error": "could not generate access token",
		})
		return
	}

	responseUser := LoginUser{
		ID:        user.ID,
		CompanyID: user.CompanyID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}

	response := LoginResponse{
		Message:     "login successful",
		AccessToken: token,
		TokenType:   "Bearer",
		ExpiresIn:   int64(h.Config.JWTExpiry.Seconds()),
		User:        responseUser,
	}

	writeJSON(
		w,
		http.StatusOK,
		response,
	)

}

func (h *AuthHandler) Me(
	w http.ResponseWriter,
	r *http.Request,
) {

	claims, ok := middleware.GetClaims(r.Context())

	if !ok {
		writeJSON(w, http.StatusUnauthorized, map[string]string{
			"error": "unauthorized",
		})
		return
	}

	if claims.UserID <= 0 || claims.CompanyID <= 0 {
		writeJSON(w, http.StatusUnauthorized, map[string]string{
			"error": "invalid token claim",
		})
		return
	}

	user, err := h.Service.GetCurrentUser(
		r.Context(),
		claims.UserID,
		claims.CompanyID,
	)

	if err != nil {
		if errors.Is(err, apperrors.ErrInvalidCredentials) {
			writeJSON(w, http.StatusUnauthorized, map[string]string{
				"error": "user not found or inactive",
			})
			return
		}
		writeJSON(w, http.StatusInternalServerError, map[string]string{
			"error": "internal server server",
		})
		return
	}
	writeJSON(w, http.StatusOK, user)

}

func (h *AuthHandler) Logout(
	w http.ResponseWriter,
	r *http.Request,
) {
	writeJSON(w, http.StatusOK, map[string]string{
		"message": "logout successful",
		"info":    "please remove access token from client storage",
	})
}

func (h *AuthHandler) RegisterRequest(
	w http.ResponseWriter,
	r *http.Request,
) {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	var request models.RegisterRequest

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&request)

	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
		return
	}

	request.FirstName = strings.TrimSpace(request.FirstName)
	request.LastName = strings.TrimSpace(request.LastName)
	request.Email = strings.TrimSpace(request.Email)
	request.DepartmentCode = strings.TrimSpace(request.DepartmentCode)

	//Call Registration Service
	err = h.Service.Register(
		r.Context(),
		request,
	)

	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrEmailAlreadyExists):
			writeJSON(w, http.StatusConflict, map[string]string{
				"error": "email already exists",
			})
		case errors.Is(err, apperrors.ErrInvalidRegistration):
			writeJSON(w, http.StatusBadRequest, map[string]string{
				"error": "invalid registration data",
			})
		default:
			writeJSON(w, http.StatusInternalServerError, map[string]string{
				"error": "internal server error",
			})
		}
		return
	}
	writeJSON(w, http.StatusCreated, RegisterResponse{
		Message: "registration successful",
	})
}
