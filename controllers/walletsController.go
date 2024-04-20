package controllers

import (
	"encoding/json"
	"mobee-test/libs/routers"
	"mobee-test/libs/utils"
	"mobee-test/models"
	"mobee-test/usecases"
	"net/http"
)

type walletController struct {
	walletUC usecases.WalletUseCase
	r        routers.Resultset
}

type WalletController interface {
	CreateWallet(w http.ResponseWriter, r *http.Request)
}

func NewWalletController(walletUC usecases.WalletUseCase, r routers.Resultset) WalletController {
	return &walletController{
		walletUC,
		r,
	}
}

func (wr *walletController) CreateWallet(w http.ResponseWriter, r *http.Request) {
	var (
		request     models.Wallet
		successResp models.SuccessResponse
		errResp     models.ErrorResponse
	)

	binding := json.NewDecoder(r.Body).Decode(&request)
	if binding != nil {
		errResp.Status = http.StatusBadRequest
		errResp.Error = utils.BAD_REQUEST
		errResp.Message = http.StatusText(http.StatusBadRequest)

		wr.r.ResponsWithError(w, http.StatusBadRequest, utils.BAD_REQUEST)
		return
	}

	err := wr.walletUC.CreateWallet(r.Context(), request)
	if err != nil {
		errResp.Status = http.StatusInternalServerError
		errResp.Error = utils.INTERNAL_SERVER_ERROR
		errResp.Message = http.StatusText(http.StatusInternalServerError)

		wr.r.ResponsWithError(w, http.StatusInternalServerError, utils.INTERNAL_SERVER_ERROR)
		return
	}

	successResp.Status = http.StatusCreated
	successResp.Message = "Wallet created successfully"
	wr.r.ResponsWithJSON(w, http.StatusCreated, successResp)
}
