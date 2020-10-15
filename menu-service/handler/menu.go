package handler

import (
	"net/http"
	"digitalent-microservice/utils"
)

func AddMenu(w http.ResponseWriter, r *http.Request) {
	utils.WrapAPISuccess(w, r ,"success", http.StatusOk)

}