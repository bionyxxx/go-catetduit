package user

import (
	"catetduit/internal/helper"
	"net/http"

	"github.com/jmoiron/sqlx"
)

// Handler handles user-related HTTP requests
type Handler struct {
	db *sqlx.DB
}

// NewHandler creates a new user handler
func NewHandler(db *sqlx.DB) *Handler {
	return &Handler{
		db,
	}
}

// GetUser returns a static dummy user
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	// Static dummy user
	user := User{
		ID:    1,
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}

	err := helper.ResponseOKWithData(w, "Retrieval successful", user)

	if err != nil {
		panic(err.Error())
	}
}
