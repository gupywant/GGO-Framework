// controllers/auth_controller.go
package controllers

import (
	"GGO/config"
	"GGO/utils"
    "GGO/models"


	"encoding/json"
	"net/http"
    "crypto/md5"
    "encoding/hex"
)

type AuthController struct {}

func (a *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	body := utils.GetJSONBody(r)

	username, _ := body["username"].(string)
	password, _ := body["password"].(string)

	user, err := models.FindUserByUsername(config.GGODB, username)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "Invalid username or password"})
		return
	}

	hashedPassword := md5.Sum([]byte(password))
	hashedPasswordString := hex.EncodeToString(hashedPassword[:])

	if user.Password != hashedPasswordString {
		json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "Invalid username or password"})
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "Login successful"})
}

func (a *AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "Logout successful"})
}
