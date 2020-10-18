package handler

import (
	"digitalent-microservice/auth-service/database"
	"digitalent-microservice/auth-service/utils"
	"encoding/json"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
)

type  AuthDB struct {
	Db *gorm.DB
}

func ValidateAuth(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		utils.WrapAPIError(w, r, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	authToken := r.Header.Get("Authorization")
	if authToken == "" {
		utils.WrapAPIError(w, r, "Invalid auth", http.StatusForbidden)
		return
	}

	//res, err := database.ValidateAuth(authToken, db.Db); if err != nil {
	//	utils.WrapAPIError(w,r,"error "+err.Error(), http.StatusBadRequest)
	//	return
	//}

	//utils.WrapAPIData(w,r,database.Auth{
	//	Username: res.Username,
	//	Token: res.Token
	//}, http.StatusOK, "success")

	if authToken != "asdfghjk" {
		utils.WrapAPIError(w, r, "Invalid auth", http.StatusForbidden)
		return
	}

	utils.WrapAPISuccess(w, r, "success", 200)
}

func (db *AuthDB) SignUp  (w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		utils.WrapAPIError(w,r,http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		utils.WrapAPIError(w,r, "cant read", http.StatusBadRequest)
		return
	}

	var signup database.Auth

	err = json.Unmarshal(body, &signup)
	if err != nil {
		utils.WrapAPIError(w,r, "error unmarshal"+err.Error(), http.StatusInternalServerError)
	}

	signup.Token = utils.IdGenerator()

	err = signup.SignUp(db.Db); if err != nil {
		utils.WrapAPIError(w,r,"error "+err.Error(), http.StatusBadRequest)
		return
	}

	utils.WrapAPISuccess(w,r,"success", http.StatusOK)
	return
	// SIGNUP
}

func (db *AuthDB) Login  (w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		utils.WrapAPIError(w,r,http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		utils.WrapAPIError(w,r, "cant read", http.StatusBadRequest)
		return
	}

	var login database.Auth

	err = json.Unmarshal(body, &login)
	if err != nil {
		utils.WrapAPIError(w,r, "error unmarshal"+err.Error(), http.StatusInternalServerError)
	}

	login.Token = utils.IdGenerator()

	res, err := login.Login(db.Db); if err != nil {
		utils.WrapAPIError(w,r,"error "+err.Error(), http.StatusBadRequest)
		return
	}

	utils.WrapAPIData(w,r, database.Auth{
		Username: res.Username,
		Token: res.Token,
	}, http.StatusOK,"success")

	// LOGIN
}