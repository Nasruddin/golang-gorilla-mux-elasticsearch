package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/nasruddin/golang/elastic/domain/users"
	"github.com/nasruddin/golang/elastic/services"
	"github.com/nasruddin/golang/elastic/utils/http_utils"
	"io/ioutil"
	"net/http"
)

var (
	UserControllerInterface userControllerInterface = &userController{}
)

type userControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type userController struct {

}


func (c *userController) Create(w http.ResponseWriter, r *http.Request)  {
	requestBody, err := ioutil.ReadAll(r.Body)

	defer r.Body.Close()

	var userRequest users.User
	if err := json.Unmarshal(requestBody, &userRequest); err != nil {
		http_utils.RespondError(w, 500, err)
	}
	result, err := services.UserServiceInterface.Create(userRequest)

	if err != nil {
		// TODO handle the error
		http_utils.RespondJson(w, 500, err)
	}

	http_utils.RespondJson(w, http.StatusCreated, result)
}

func (c *userController) Get(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	userId := vars["id"]
	user, err := services.UserServiceInterface.Get(userId)

	if err != nil {
		http_utils.RespondError(w, 500, err)
	}

	http_utils.RespondJson(w, http.StatusOK, user)
}

