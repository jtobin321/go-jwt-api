package views

import (
	"net/http"

	"github.com/jtobin321/go-jwt-api/api/utils"
)

// Home is welcome entry point for API
func Home(w http.ResponseWriter, r *http.Request) {
	utils.JSON(w, http.StatusOK, "Welcome To this Awesome API")
}
