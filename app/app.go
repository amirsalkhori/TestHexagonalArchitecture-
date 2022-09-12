package app

import (
	"fmt"
	"goHexagonal/domain"
	"goHexagonal/service"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func Start() {
	fmt.Println("This is my log message")
	// mux := http.NewServeMux()
	customerHandler := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	router := mux.NewRouter()
	router.HandleFunc("/customers", customerHandler.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", customerHandler.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers", postCustomer).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe("localhost:8001", router))
}

func postCustomer(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "This is a post method")
}