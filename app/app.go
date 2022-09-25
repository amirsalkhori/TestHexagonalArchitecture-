package app

import (
	"fmt"
	"goHexagonal/domain"
	"goHexagonal/logger"
	"goHexagonal/service"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
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
	router := mux.NewRouter()
	dbClient := getDbClient()

	cusotmerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)

	customerHandler := CustomerHandler{service.NewCustomerService(cusotmerRepositoryDb)}
	accountHandler := AccountHandler{service.NewAccountService(accountRepositoryDb)}

	router.HandleFunc("/customers", customerHandler.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", customerHandler.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", accountHandler.CreateAccount).Methods(http.MethodPost)
	router.HandleFunc("/customers", postCustomer).Methods(http.MethodPost)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	// log.Fatal(http.ListenAndServe("localhost:8001", router))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}

func postCustomer(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "This is a post method")
}

func getDbClient() *sqlx.DB{
		// db_user := os.Getenv("DB_USER")
	// db_pass := os.Getenv("DB_PASSWD")
	// db_address := os.Getenv("DB_ADDR")
	// db_port := os.Getenv("DB_PORT")
	// db_name := os.Getenv("DB_NAME")

	// dbConfig := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db_user, db_pass,db_address, db_port, db_name)
	// client, err := sqlx.Open("mysql", dbConfig)

	client, err := sqlx.Open("mysql", "rooti:changeMe@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)		
	return client
}