package controllers

import (
	"net/http"

	"github.com/jtobin321/go-jwt-api/api/responses"
)

// Home is welcome entry point for API
func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To this Awesome API")
}
