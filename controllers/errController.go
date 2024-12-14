package controllers

import (
	"groupie_tracker/models"
	"net/http"
)

func ErrorController(w http.ResponseWriter, r *http.Request, statusCode int) {
	errorType := models.Error{
		StatusCode: statusCode,
		Message: http.StatusText(statusCode),
	}
	w.WriteHeader(statusCode)
	ParseController(w, r, "error", errorType)
}
