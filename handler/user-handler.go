package handler

import (
	"api-google-keep/core/entity"
	"api-google-keep/core/module"
	"api-google-keep/libs/response"
	"encoding/json"
	"net/http"
)

type userHandler struct {
	userUsecase module.UserUsecase
}

type UserHandler interface{
	RegisterUser(w http.ResponseWriter, r *http.Request)
}

func NewUserHandler(userUsecase module.UserUsecase) UserHandler {
	return &userHandler{userUsecase}
}

func (c *userHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user entity.UserCreate

	err :=  json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, nil, err.Error())
	}

	var noDataResponse response.NoDataResponse
	noDataResponse.SetData(nil)
	noDataResponse.SetMessage("success create user")
	response.Success(w, http.StatusCreated, noDataResponse)
}