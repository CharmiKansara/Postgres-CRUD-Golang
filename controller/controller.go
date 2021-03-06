package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"postgres-crud/database"
	"postgres-crud/model"

	"github.com/gorilla/mux"
)

//Insert customer into database
func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	connection := database.GetDatabase()
	defer database.Closedatabase(connection)

	var customer model.Customer
	json.NewDecoder(r.Body).Decode(&customer)
	log.Println(customer)
	connection.Create(&customer)
	log.Println("Customer data inserted successfully into database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

//Get all customers from database
func GetAllCustomer(w http.ResponseWriter, r *http.Request) {
	connection := database.GetDatabase()
	defer database.Closedatabase(connection)

	var customers []model.Customer
	connection.Find(&customers)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func GetOneCustomer(w http.ResponseWriter, r *http.Request) {
	connection := database.GetDatabase()
	defer database.Closedatabase(connection)

	id := mux.Vars(r)["id"]
	var customer []model.Customer
	connection.First(&customer, id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

//Update customer in database using id
func UpdateOneCustomer(w http.ResponseWriter, r *http.Request) {
	connection := database.GetDatabase()
	defer database.Closedatabase(connection)

	id := mux.Vars(r)["id"]
	var customer model.Customer
	connection.First(&customer, id)
	json.NewDecoder(r.Body).Decode(&customer)
	connection.Save(&customer)
	log.Println("Customer data updated successfully into database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

//Delete customer in database using id
func DeleteOneCustomer(w http.ResponseWriter, r *http.Request) {
	connection := database.GetDatabase()
	defer database.Closedatabase(connection)

	id := mux.Vars(r)["id"]
	var customer model.Customer
	connection.Delete(&customer, id)

	w.Write([]byte("Customer deleted successfully - id: " + id))
}
