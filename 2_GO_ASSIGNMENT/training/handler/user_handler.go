package handler

import (
	"errors"
	"net/http"
	"strconv"
	apperrors "training/errors"
	"training/middleware"
	domain "training/service"
)

type UserHandler struct {
	Service *domain.AuthService
}

func NewUserHandler(
	service *domain.AuthService,
) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

func (h *UserHandler) GetUserByID(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Get User ID from URL
	id, err := strconv.ParseInt(
		r.PathValue("id"),
		10,
		64,
	)

	if err != nil || id <= 0 {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "invalid user id",
		})
		return
	}
	claims, ok := middleware.GetClaims(r.Context())

	if !ok {
		writeJSON(w, http.StatusUnauthorized, map[string]string{
			"error": "unauthorized",
		})
		return
	}

	user, err := h.Service.GetCurrentUser(
		r.Context(),
		id,
		claims.CompanyID,
	)
	if err != nil {
		if errors.Is(err, apperrors.ErrInvalidCredentials) {
			writeJSON(w, http.StatusNotFound, map[string]string{
				"error": "user not found",
			})
			return
		}
		writeJSON(w, http.StatusInternalServerError, map[string]string{
			"error": "internal server error",
		})
		return
	}
	writeJSON(w, http.StatusOK, user)
}

//baki ahe ajun implementation user konta kontya company cha data update delete kru shakto
