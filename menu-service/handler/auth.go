package handler

import (
	"digitalent-microservice/menu-service/entity"
	"encoding/json"
	"digitalent-microservice/menu-service/config"
	"digitalent-microservice/menu-service/utils"
	"github.com/gorilla/context"
	"io/ioutil"
	"net/http"
)

type AuthMiddleware struct {
	AuthService config.AuthService
}

func (auth *AuthMiddleware) ValidateAuth(nextHandler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request, err := http.NewRequest("POST", auth.AuthService.Host+"/auth/validate", nil)
		if err != nil {
			utils.WrapAPIError(w, r, "failed to create request : "+err.Error(), http.StatusInternalServerError)
			return
		}

		request.Header = r.Header
		authResponse, err := http.DefaultClient.Do(request)
		if err != nil {
			utils.WrapAPIError(w, r, "validate auth failed : "+err.Error(), http.StatusInternalServerError)
			return
		}
		defer authResponse.Body.Close()

		body, err := ioutil.ReadAll(authResponse.Body)
		if err != nil {
			utils.WrapAPIError(w, r, err.Error(), http.StatusInternalServerError)
			return
		}

		//var authResult map[string]interface{}
		var authResult entity.AuthResponse
		err = json.Unmarshal(body, &authResult)

		if authResponse.StatusCode != 200 {
			utils.WrapAPIError(w, r, authResult.ErrorDetails, authResponse.StatusCode)
			return
		}

		context.Set(r, "user", authResult.Data.Username)
		// mengirimkan data ke menu handler
		nextHandler(w, r)
	}
}