package controllers

import (
	"encoding/json"
	"mobee-test/libs/routers"
	"mobee-test/libs/utils"
	"mobee-test/models"
	"mobee-test/usecases"
	"net/http"
)

type userController struct {
	userUC usecases.UserUseCase
	r      routers.Resultset
}

type UserController interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
}

func NewUserController(userUc usecases.UserUseCase, r routers.Resultset) UserController {
	return &userController{
		userUc,
		r,
	}
}

func (u *userController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var (
		request     models.Users
		successResp models.SuccessResponse
		errResp     models.ErrorResponse
	)

	binding := json.NewDecoder(r.Body).Decode(&request)
	if binding != nil {
		errResp.Status = http.StatusBadRequest
		errResp.Error = utils.BAD_REQUEST
		errResp.Message = http.StatusText(http.StatusBadRequest)

		u.r.ResponsWithError(w, http.StatusBadRequest, utils.BAD_REQUEST)
		return
	}

	err := u.userUC.CreateUser(r.Context(), request)
	if err != nil {
		errResp.Status = http.StatusInternalServerError
		errResp.Error = utils.INTERNAL_SERVER_ERROR
		errResp.Message = http.StatusText(http.StatusInternalServerError)

		u.r.ResponsWithError(w, http.StatusInternalServerError, utils.INTERNAL_SERVER_ERROR)
		return
	}

	successResp.Status = http.StatusCreated
	successResp.Message = "User created successfully"
	u.r.ResponsWithJSON(w, http.StatusCreated, successResp)
}
