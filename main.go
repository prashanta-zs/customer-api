package main

import (
	"database/sql"
	"net/http"

	"github.com/prashant/customerApi/handler"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var Db *sql.DB
var err error

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/customer", handler.GetAllCustomers).Methods(http.MethodGet)
	r.HandleFunc("/customer/{id}", handler.GetCustomerById).Methods(http.MethodGet)
	r.HandleFunc("/customer/insert", handler.AddCustomer).Methods(http.MethodPost)
	r.HandleFunc("/customer/update/{id}", handler.UpdateCustomer).Methods(http.MethodPut)
	r.HandleFunc("/customer/delete/{id}", handler.DeleteCustomerDetails).Methods(http.MethodDelete)
	http.ListenAndServe(":8000", r)

}

//package main
//
//import (
//	"database/sql"
//	"fmt"
//	"log"
//
//	_ "github.com/go-sql-driver/mysql"
//)
//
//type Customers struct {
//	Id      int
//	Name    string
//	Phone   int
//	Address string
//}
//
//func main() {
//
//	db, err := sql.Open("mysql", "root:maji815355@tcp(127.0.0.1:3306)/assignment")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	defer db.Close()
//
//	res, err := db.Query("SELECT * FROM CUSTOMER")
//	//	defer res.Close()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for res.Next() {
//
//		var customer Customers
//		err := res.Scan(&customer.Id, &customer.Name, &customer.Phone, &customer.Address)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		fmt.Printf("%v\n", customer)
//	}
//}
