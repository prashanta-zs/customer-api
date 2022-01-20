package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Customers struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Phone   int    `json:"phone"`
	Address string `json:"address"`
}

//GetAllCustomers using GET Method
func GetAllCustomers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application-Json")
	Db, err := sql.Open("mysql", "root:maji815355@tcp(127.0.0.1:3306)/assignment")
	if err != nil {
		log.Fatal(err)
	}

	defer Db.Close()

	var c []Customers

	res, err := Db.Query("SELECT * FROM CUSTOMER")

	defer res.Close()

	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {

		var customer Customers
		err := res.Scan(&customer.Id, &customer.Name, &customer.Phone, &customer.Address)
		if err != nil {
			log.Fatal(err)
		}
		c = append(c, customer)
		//w.WriteHeader(http.StatusBadRequest)
	}

	CustJson, _ := json.Marshal(c)
	w.Write(CustJson)
}

//GetCustomerById using GET method
func GetCustomerById(w http.ResponseWriter, r *http.Request) {

	Db, err := sql.Open("mysql", "root:maji815355@tcp(127.0.0.1:3306)/assignment")

	if err != nil {
		log.Fatal(err)
	}

	defer Db.Close()

	params := mux.Vars(r)
	v := params["id"]
	id, _ := strconv.Atoi(v)

	res, err := Db.Query("SELECT * FROM CUSTOMER WHERE id = ?", id)
	defer res.Close()

	if err != nil {
		log.Fatal(err)
	}

	var customer Customers

	for res.Next() {

		err := res.Scan(&customer.Id, &customer.Name, &customer.Phone, &customer.Address)
		if err != nil {
			log.Fatal(err)
		}
		//w.WriteHeader(http.StatusBadRequest)

		CustJson, _ := json.Marshal(customer)
		w.Write(CustJson)
	}

}

//AddCustomer using POST method
func AddCustomer(w http.ResponseWriter, r *http.Request) {

	Db, err := sql.Open("mysql", "root:maji815355@tcp(127.0.0.1:3306)/assignment")

	if err != nil {
		log.Fatal(err)
	}

	defer Db.Close()

	var person Customers

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(requestBody, &person)
	if err != nil {
		log.Fatal(err)
	}

	_, err = Db.Exec("INSERT INTO CUSTOMER(Name, Phone, Address) VALUES (? ,?, ?)", person.Name, person.Phone, person.Address)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(requestBody)
}

//UpdateCustomer using PUT method
func UpdateCustomer(w http.ResponseWriter, r *http.Request) {

	Db, err := sql.Open("mysql", "root:maji815355@tcp(127.0.0.1:3306)/assignment")

	if err != nil {
		log.Fatal(err)
	}

	defer Db.Close()

	param := mux.Vars(r)
	v := param["id"]
	id, _ := strconv.Atoi(v)

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var c Customers
	json.Unmarshal(requestBody, &c)

	_, err = Db.Exec("UPDATE CUSTOMER SET Name = ?, Phone = ?, Address = ? WHERE id = ?", c.Name, c.Phone, c.Address, id)
	if err != nil {
		log.Fatal(err)
	}
	c.Id = id
	w.Write([]byte(fmt.Sprint(c)))

}

//DeleteCustomerDetails using DELETE method
func DeleteCustomerDetails(w http.ResponseWriter, r *http.Request) {

	Db, err := sql.Open("mysql", "root:maji815355@tcp(127.0.0.1:3306)/assignment")
	if err != nil {
		log.Fatal(err)
	}

	defer Db.Close()

	params := mux.Vars(r)
	v := params["id"]
	id, _ := strconv.Atoi(v)

	_, err = Db.Exec("DELETE FROM CUSTOMER WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
}
