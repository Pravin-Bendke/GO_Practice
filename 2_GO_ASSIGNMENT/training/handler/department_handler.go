package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	apperrors "training/errors"
	"training/models"
	"training/repository"
	reqctx "training/request"
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

func getRequestIdentity(r *http.Request) (
	requestID string,
	userID int64,
	companyID int64,
	ok bool,
) {
	ctx := r.Context()

	requestID = reqctx.GetRequestID(ctx)
	userID = reqctx.GetUserID(ctx)
	companyID = reqctx.GetCompanyID(ctx)

	if userID <= 0 || companyID <= 0 {
		return requestID, userID, companyID, false
	}

	return requestID, userID, companyID, true
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

	case errors.Is(err, apperrors.ErrDepartmentHasActiveUsers):

		writeJSON(w, http.StatusConflict, map[string]string{
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

	requestID, userID, companyID, ok := getRequestIdentity(r)

	if !ok {
		writeJSON(w, http.StatusUnauthorized, map[string]string{
			"error": "authentication required",
		})
		return
	}

	log.Printf(
		"GET DEPARTMENT: RequestID=%s UserID=%d CompanyID=%d",
		requestID,
		userID,
		companyID,
	)

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

	requestID, userID, companyID, ok := getRequestIdentity(r)

	if !ok {
		writeJSON(w, http.StatusUnauthorized, map[string]string{
			"error": "authentication required",
		})
		return
	}

	log.Printf(
		"LIST DEPARTMENTS: RequestID=%s UserID=%d CompanyID=%d",
		requestID,
		userID,
		companyID,
	)

	//default page and page record size
	page := 1
	pageSize := 10

	var err error

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

func (h *DepartmentHandler) CreateDepartment(
	w http.ResponseWriter,
	r *http.Request,
) {
	var request models.CreateDepartmentRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
		return
	}
	requestID, userID, companyID, ok := getRequestIdentity(r)

	if !ok {
		writeJSON(w, http.StatusUnauthorized, map[string]string{
			"error": "authentication required",
		})
		return
	}

	log.Printf(
		"CREATE DEPARTMENT: RequestID=%s UserID=%d CompanyID=%d",
		requestID,
		userID,
		companyID,
	)

	//department model create

	department := &models.Department{
		CompanyID: companyID,
		Code:      request.Code,
		Name:      request.Name,

		Description: request.Description,
	}

	created, err := h.Service.CreateDepartment(
		r.Context(),
		department,
	)
	if err != nil {
		log.Printf("Created Department failed:%v", err)
		handleDepartmentError(w, err)
		return
	}

	writeJSON(
		w,
		http.StatusCreated,
		created,
	)
}

func (h *DepartmentHandler) UpdateDepartment(w http.ResponseWriter, r *http.Request) {
	idstr := r.PathValue("id")

	id, err := strconv.ParseInt(idstr, 10, 64)

	if err != nil || id <= 0 {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "invalid department id",
		})
		return
	}

	var req models.UpdateDepartmentRequest

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
		return
	}

	req.Code = strings.TrimSpace(req.Code)
	req.Name = strings.TrimSpace(req.Name)
	req.Description = strings.TrimSpace(req.Description)

	if req.Code == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "department code is required",
		})
		return
	}

	if req.Name == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "department name is required",
		})
		return
	}

	requestID, userID, companyID, ok := getRequestIdentity(r)

	if !ok {
		writeJSON(w, http.StatusUnauthorized, map[string]string{
			"error": "authentication required",
		})
		return
	}

	log.Printf(
		"UPDATE DEPARTMENT: RequestID=%s DepartmentID=%d UserID=%d CompanyID=%d",
		requestID,
		id,
		userID,
		companyID,
	)

	if companyID <= 0 {
		handleDepartmentError(w, apperrors.ErrInvalidCompany)
		return
	}

	updated, err := h.Service.UpdateDepartment(
		r.Context(),
		id,
		companyID,
		req,
	)

	if err != nil {
		log.Printf("UPDATE SERVICE ERROR: %v", err)
		handleDepartmentError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, updated)
}

func (h *DepartmentHandler) UpdateDepartmentStatus(
	w http.ResponseWriter,
	r *http.Request,
) {
	//Get dept ID from URL

	id, err := strconv.ParseInt(
		r.PathValue("id"),
		10, 64,
	)

	if err != nil || id <= 0 {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "invalid dept ID",
		})
		return
	}
	var req struct {
		IsActive *bool `json:"is_active"`
	}

	err = json.NewDecoder(r.Body).Decode(&req)

	if err != nil || req.IsActive == nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "is_active is required and must be boolean",
		})
		return
	}

	requestID, userID, companyID, ok := getRequestIdentity(r)

	if !ok {
		writeJSON(w, http.StatusUnauthorized, map[string]string{
			"error": "authentication required",
		})
		return
	}

	log.Printf(
		"UPDATE DEPARTMENT STATUS: RequestID=%s DepartmentID=%d UserID=%d CompanyID=%d",
		requestID,
		id,
		userID,
		companyID,
	)

	err = h.Service.UpdateStatus(
		r.Context(),
		id,
		companyID,
		*req.IsActive,
	)
	if err != nil {
		handleDepartmentError(w, err)
		return
	}

	// Success Response
	message := "department deactivated successfully"

	if *req.IsActive {
		message = "department activated successfully"
	}

	response := map[string]interface{}{
		"message":   message,
		"is_active": *req.IsActive,
	}

	writeJSON(
		w,
		http.StatusOK,
		response,
	)

}

func (h *DepartmentHandler) DeleteDepartment(
	w http.ResponseWriter,
	r *http.Request,
) {
	// Get Department ID from URL
	id, err := strconv.ParseInt(
		r.PathValue("id"),
		10,
		64,
	)
	if err != nil || id <= 0 {
		http.Error(
			w,
			`{"error":"invalid department id"}`,
			http.StatusBadRequest,
		)
		return
	}
	// Get JWT Claims

	requestID, userID, companyID, ok := getRequestIdentity(r)

	if !ok {
		writeJSON(w, http.StatusUnauthorized, map[string]string{
			"error": "authentication required",
		})
		return
	}

	log.Printf(
		"DELETE DEPARTMENT: RequestID=%s DepartmentID=%d UserID=%d CompanyID=%d",
		requestID,
		id,
		userID,
		companyID,
	)
	if companyID <= 0 {
		handleDepartmentError(
			w,
			apperrors.ErrInvalidCompany,
		)
		return
	}
	log.Println("DELETE DEBUG: Calling DeleteDepartment service")
	//call service
	err = h.Service.DeleteDepartment(
		r.Context(),
		id,
		companyID,
	)
	if err != nil {
		log.Printf("DELETE SERVICE ERROR: %v", err)
		handleDepartmentError(w, err)
		return
	}

	//success Response
	writeJSON(
		w,
		http.StatusOK,
		map[string]string{
			"message": "department deleted successfully",
		},
	)

}
