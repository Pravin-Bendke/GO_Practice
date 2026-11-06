package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	apperrors "training/errors"
	"training/repository"
	domain "training/service"
)

type DepartmentHandler struct {
	Service *domain.DepartmentService
}

func NewDepartmentHandler(
	service *domain.DepartmentService,
) *DepartmentHandler {

	return &DepartmentHandler{
		Service: service,
	}
}
func writeJSON(
	w http.ResponseWriter,
	status int,
	data interface{},
) {

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	w.WriteHeader(status)

	json.NewEncoder(w).Encode(data)
}

func handleDepartmentError(
	w http.ResponseWriter,
	err error,
) {

	switch {

	case errors.Is(err, apperrors.ErrDepartmentNotFound):

		writeJSON(w, http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})

	case errors.Is(err, apperrors.ErrInvalidCompany),
		errors.Is(err, apperrors.ErrInvalidDepartmentID),
		errors.Is(err, apperrors.ErrInvalidPagination):

		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})

	default:

		writeJSON(w, http.StatusInternalServerError, map[string]string{
			"error": "internal server error",
		})
	}
}

func (h *DepartmentHandler) GetDepartment(
	w http.ResponseWriter,
	r *http.Request,
) {

	id, err := strconv.ParseInt(
		r.PathValue("id"),
		10,
		64,
	)

	if err != nil {
		handleDepartmentError(
			w,
			apperrors.ErrInvalidDepartmentID,
		)
		return
	}

	companyID, err := strconv.ParseInt(
		r.URL.Query().Get("company_id"),
		10,
		64,
	)

	if err != nil {
		handleDepartmentError(
			w,
			apperrors.ErrInvalidCompany,
		)
		return
	}

	department, err := h.Service.GetDepartment(
		r.Context(),
		id,
		companyID,
	)

	if err != nil {
		handleDepartmentError(w, err)
		return
	}

	writeJSON(
		w,
		http.StatusOK,
		department,
	)
}

func (h *DepartmentHandler) ListDepartments(
	w http.ResponseWriter,
	r *http.Request,
) {

	query := r.URL.Query()

	companyID, err := strconv.ParseInt(
		query.Get("company_id"),
		10,
		64,
	)

	if err != nil {
		handleDepartmentError(
			w,
			apperrors.ErrInvalidCompany,
		)
		return
	}

	page := 1
	pageSize := 20

	if query.Get("page") != "" {

		page, err = strconv.Atoi(query.Get("page"))

		if err != nil {
			handleDepartmentError(
				w,
				apperrors.ErrInvalidPagination,
			)
			return
		}
	}

	if query.Get("page_size") != "" {

		pageSize, err = strconv.Atoi(query.Get("page_size"))

		if err != nil {
			handleDepartmentError(
				w,
				apperrors.ErrInvalidPagination,
			)
			return
		}
	}

	filters := repository.QueryFilters{
		Page:      page,
		PageSize:  pageSize,
		Search:    query.Get("search"),
		SortBy:    query.Get("sort_by"),
		SortOrder: query.Get("sort_order"),
	}

	departments, total, err := h.Service.ListDepartments(
		r.Context(),
		companyID,
		filters,
	)

	if err != nil {
		handleDepartmentError(w, err)
		return
	}

	response := map[string]interface{}{

		"data": departments,

		"pagination": map[string]interface{}{
			"page":        page,
			"page_size":   pageSize,
			"total":       total,
			"total_pages": (total + int64(pageSize) - 1) / int64(pageSize),
		},
	}

	writeJSON(
		w,
		http.StatusOK,
		response,
	)
}
