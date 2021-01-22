package views

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/jtobin321/go-jwt-api/api/auth"
	"github.com/jtobin321/go-jwt-api/api/models"
	"github.com/jtobin321/go-jwt-api/api/utils"
	"golang.org/x/crypto/bcrypt"
)

// Login is used to login user
func Login(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := signIn(user.Email, user.Password, db)
	if err != nil {
		formattedError := utils.FormatError(err.Error())
		utils.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	utils.JSON(w, http.StatusOK, token)
}

func signIn(email, password string, db *gorm.DB) (string, error) {

	var err error

	user := models.User{}

	err = db.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(user.ID)
}
