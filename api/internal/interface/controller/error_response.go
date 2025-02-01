package handler

import (
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v4"
)

type Error struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

type ValidationErrorResponse struct {
	Message string  `json:"message"`
	Errors  []Error `json:"errors"`
}

func respondValidationError(errors validation.Errors) error {
	var params []Error

	for key, val := range errors {
		params = append(params, Error{Name: key, Message: val.Error()})
	}

	return echo.NewHTTPError(
		http.StatusBadRequest,
		ValidationErrorResponse{
			Message: "validation error",
			Errors:  params,
		},
	)
}
