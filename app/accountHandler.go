package app

import (
	"encoding/json"
	"goHexagonal/dto"
	"goHexagonal/service"
	"net/http"

	"github.com/gorilla/mux"
)

type AccountHandler struct {
	service service.AccountService
}

func(ah *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	customerId := vars["customer_id"]

	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, r, http.StatusBadRequest, err.Error())
	}else{
		request.CustomerId = customerId
		account, appError := ah.service.NewAccount(request)
		if appError != nil {
			writeResponse(w, r, appError.Code, err.Error())
		}else{
			writeResponse(w, r, http.StatusCreated, account)
		}
	}	
}