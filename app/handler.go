package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"goHexagonal/service"
	"net/http"

	"github.com/gorilla/mux"
)

type CustomerHandler struct {
	service service.CustomerService
}


func(customerHandler *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request){
	customers, _ := customerHandler.service.GetAllCustomers()
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	}else{
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func(ch *CustomerHandler) getCustomer(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["customer_id"]
	customers, err := ch.service.GetById(id)
	if err != nil {
		w.WriteHeader(err.Code)
		fmt.Fprintf(w, err.Message)
	}
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	}else{
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}