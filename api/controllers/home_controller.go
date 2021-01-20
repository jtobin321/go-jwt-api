package controllers

import (
	"net/http"

	"github.com/jtobin321/go-jwt-api/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To this Awesome API")
}
