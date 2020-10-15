package handler

import (
	"net/http"
	"digitalent-microservice/menu-service/utils"
)

func AddMenu(w http.ResponseWriter, r *http.Request) {
	utils.WrapAPISuccess(w, r ,"success", http.StatusOK)

}