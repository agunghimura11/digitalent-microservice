package handler

import (
	"github.com/gorilla/context"
	"encoding/json"
	"digitalent-microservice/menu-service/database"
	"digitalent-microservice/menu-service/utils"
	"fmt"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
)

type Menu struct {
	Db *gorm.DB
}

// AddMenuHandler handle add menu
func (menu *Menu) AddMenu(w http.ResponseWriter, r *http.Request) {
	log.Printf("hello")
	if r.Method != "POST" {
		utils.WrapAPIError(w, r, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		utils.WrapAPIError(w, r, "can't read body", http.StatusBadRequest)
		return
	}

	username := context.Get(r, "user")
	// untuk mendapatkan value dari service auth-servie berupa username user login, informasi didapatkan dari token user

	var dataMenu database.Menu
	err = json.Unmarshal(body, &dataMenu)
	if err != nil {
		utils.WrapAPIError(w, r, "error unmarshal : "+err.Error(), http.StatusInternalServerError)
		return
	}

	dataMenu.Username = fmt.Sprintf("%v", username)

	err = dataMenu.Insert(menu.Db)
	if err != nil {
		utils.WrapAPIError(w, r, "insert menu error : "+err.Error(), http.StatusInternalServerError)
	}
	utils.WrapAPISuccess(w, r, "success", 200)
}

func (menu *Menu) GetAllMenu(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		utils.WrapAPIError(w, r, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	menuDb := database.Menu{}

	menus, err := menuDb.GetAll(menu.Db)
	if err != nil {
		utils.WrapAPIError(w, r, "failed get menu:"+err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WrapAPIData(w, r, menus, 200, "success")
}