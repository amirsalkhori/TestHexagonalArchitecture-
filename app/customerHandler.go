package app

import (
	"encoding/json"
	"encoding/xml"
	"goHexagonal/service"
	"net/http"

	"github.com/gorilla/mux"
)

type CustomerHandler struct {
	service service.CustomerService
}


func(customerHandler *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request){
	customers, err := customerHandler.service.GetAllCustomers()
	if err != nil {
		writeResponse(w, r, err.Code, err.AsMessage())
	}
	writeResponse(w, r, http.StatusOK, customers)
}

func(ch *CustomerHandler) getCustomer(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["customer_id"]
	customer, err := ch.service.GetById(id)
	if err != nil {
		writeResponse(w, r, err.Code, err.AsMessage())
	}
	writeResponse(w, r, http.StatusOK, customer)
}

func writeResponse(w http.ResponseWriter, r *http.Request, statusCode int, data interface{} ){
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		w.WriteHeader(statusCode)
		err := xml.NewEncoder(w).Encode(data)
		if err != nil {
			panic(err)
		}
	}else{
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			panic(err)
		}
	}
}