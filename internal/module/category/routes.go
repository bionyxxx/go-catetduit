package category

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func RegisterRoutes(r chi.Router, validator *validator.Validate, categoryService *Service) {
	handler := NewHandler(categoryService, validator)

	r.Get("/categories", handler.GetCategoriesByUser)
	r.Get("/categories/{id}", handler.GetCategory)
	r.Post("/categories", handler.CreateCategory)
}
