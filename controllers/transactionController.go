package controllers

import (
	"encoding/json"
	"mobee-test/libs/routers"
	"mobee-test/libs/utils"
	"mobee-test/models"
	"mobee-test/usecases"
	"net/http"
)

type transactionController struct {
	transactionUC usecases.TransactionUseCase
	r             routers.Resultset
}

type TransactionController interface {
	Deposit(w http.ResponseWriter, r *http.Request)
}

func NewTransactionController(transactionUC usecases.TransactionUseCase, r routers.Resultset) TransactionController {
	return &transactionController{
		transactionUC,
		r,
	}
}

func (tr *transactionController) Deposit(w http.ResponseWriter, r *http.Request) {
	var (
		request     models.DepositRequest
		successResp models.SuccessResponse
		errResp     models.ErrorResponse
	)

	binding := json.NewDecoder(r.Body).Decode(&request)
	if binding != nil {
		errResp.Status = http.StatusBadRequest
		errResp.Error = utils.BAD_REQUEST
		errResp.Message = http.StatusText(http.StatusBadRequest)

		tr.r.ResponsWithError(w, http.StatusBadRequest, utils.BAD_REQUEST)
		return
	}

	err := tr.transactionUC.Deposit(r.Context(), request)
	if err != nil {
		errResp.Status = http.StatusInternalServerError
		errResp.Error = utils.INTERNAL_SERVER_ERROR
		errResp.Message = http.StatusText(http.StatusInternalServerError)

		tr.r.ResponsWithError(w, http.StatusInternalServerError, utils.INTERNAL_SERVER_ERROR)
		return
	}

	successResp.Status = http.StatusOK
	successResp.Message = "Deposit successfully"
	tr.r.ResponsWithJSON(w, http.StatusOK, successResp)
}
