package helper

import (
	"encoding/json"
	"net/http"
)

// Response represents a standard API response
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func WriteJSON(w http.ResponseWriter, statusCode int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(data)
}

func ResponseOKWithData(w http.ResponseWriter, message string, data interface{}) error {
	response := Response{
		Success: true,
		Message: message,
		Data:    data,
	}
	return WriteJSON(w, http.StatusOK, response)
}

func ResponseOK(w http.ResponseWriter, message string) error {
	response := Response{
		Success: true,
		Message: message,
	}
	return WriteJSON(w, http.StatusOK, response)
}

func ResponseCreated(w http.ResponseWriter, message string, data interface{}) error {
	response := Response{
		Success: true,
		Message: message,
		Data:    data,
	}
	return WriteJSON(w, http.StatusCreated, response)
}

func ResponseAccepted(w http.ResponseWriter, message string, data interface{}) error {
	response := Response{
		Success: true,
		Message: message,
		Data:    data,
	}
	return WriteJSON(w, http.StatusAccepted, response)
}

func ResponseNoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func ResponseBadRequest(w http.ResponseWriter, message string, err interface{}) error {
	response := Response{
		Success: false,
		Message: message,
		Error:   err,
	}
	return WriteJSON(w, http.StatusBadRequest, response)
}

// Unauthorized sends an unauthorized error with status 401
func ResponseUnauthorized(w http.ResponseWriter, message string) error {
	response := Response{
		Success: false,
		Message: message,
	}
	return WriteJSON(w, http.StatusUnauthorized, response)
}

// Forbidden sends a forbidden error with status 403
func ResponseForbidden(w http.ResponseWriter, message string) error {
	response := Response{
		Success: false,
		Message: message,
	}
	return WriteJSON(w, http.StatusForbidden, response)
}

// NotFound sends a not found error with status 404
func ResponseNotFound(w http.ResponseWriter, message string) error {
	response := Response{
		Success: false,
		Message: message,
	}
	return WriteJSON(w, http.StatusNotFound, response)
}

// MethodNotAllowed sends a method not allowed error with status 405
func ResponseMethodNotAllowed(w http.ResponseWriter, message string) error {
	response := Response{
		Success: false,
		Message: message,
	}
	return WriteJSON(w, http.StatusMethodNotAllowed, response)
}

// Conflict sends a conflict error with status 409
func ResponseConflict(w http.ResponseWriter, message string, err interface{}) error {
	response := Response{
		Success: false,
		Message: message,
		Error:   err,
	}
	return WriteJSON(w, http.StatusConflict, response)
}

// UnprocessableEntity sends an unprocessable entity error with status 422
func ResponseUnprocessableEntity(w http.ResponseWriter, message string, err interface{}) error {
	response := Response{
		Success: false,
		Message: message,
		Error:   err,
	}
	return WriteJSON(w, http.StatusUnprocessableEntity, response)
}

// InternalServerError sends an internal server error with status 500
func ResponseInternalServerError(w http.ResponseWriter, message string, err interface{}) error {
	response := Response{
		Success: false,
		Message: message,
		Error:   err,
	}
	return WriteJSON(w, http.StatusInternalServerError, response)
}

// ServiceUnavailable sends a service unavailable error with status 503
func ResponseServiceUnavailable(w http.ResponseWriter, message string) error {
	response := Response{
		Success: false,
		Message: message,
	}
	return WriteJSON(w, http.StatusServiceUnavailable, response)
}
