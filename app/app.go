package app

import (
	"fmt"
	"goHexagonal/domain"
	"goHexagonal/logger"
	"goHexagonal/service"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)
func sanityCheck(){
	if os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == "" {
		log.Fatal("The enviorment not defined!")
	}
}
func Start() {
	sanityCheck()
	
	logger.Info("Starting app")
	// mux := http.NewServeMux()
	customerHandler := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	router := mux.NewRouter()
	router.HandleFunc("/customers", customerHandler.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", customerHandler.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers", postCustomer).Methods(http.MethodPost)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	// log.Fatal(http.ListenAndServe("localhost:8001", router))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}

func postCustomer(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "This is a post method")
}