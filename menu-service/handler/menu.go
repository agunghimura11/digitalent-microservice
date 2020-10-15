package handler

import (
	"io/ioutil"
	"encoding/json"
	"net/http"
	"gorm.io/gorm"
	"digitalent-microservice/menu-service/utils"

	"digitalent-microservice/menu-service/database"
)

type MenuHandler struct {
	db *gorm.DB
}

func  (handler *MenuHandler) AddMenu(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		utils.WrapAPIError(w, r, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		utils.WrapAPIError(w,r,err.Error(), http.StatusInternalServerError)
		return
	}

	var menu database.Menu
	err = json.Unmarshal(body, &menu)
	if err != nil {
		utils.WrapAPIError(w,r,err.Error(), http.StatusInternalServerError)
		return
	}

	err = menu.Insert(handler.db)
	if err != nil {
		utils.WrapAPIError(w,r,err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WrapAPISuccess(w, r ,"success", http.StatusOK)

}

func (handler *MenuHandler) GetMenu(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		utils.WrapAPIError(w, r, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	menuDb := database.Menu{}

	menus, err := menuDb.GetAll(handler.db)
	if err != nil {
		utils.WrapAPIError(w, r, "failed get menu:"+err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WrapAPIData(w, r, menus, 200, "success")
}