package routes

import (
	"log"
	"net/http"
	"postgres-crud/controller"

	"github.com/gorilla/mux"
)

var router *mux.Router

func CreateRouter() {
	//Initialize router
	router = mux.NewRouter()
}

//Routes
func InitializeRoutes() {
	router.HandleFunc("/customer", controller.GetAllCustomer).Methods("GET")
	router.HandleFunc("/customer", controller.CreateCustomer).Methods("POST")
	router.HandleFunc("/customer/{id}", controller.GetOneCustomer).Methods("GET")
	router.HandleFunc("/customer/{id}", controller.UpdateOneCustomer).Methods("PUT")
	router.HandleFunc("/customer/{id}", controller.DeleteOneCustomer).Methods("DELETE")
}

func StartServer() {
	log.Println("Server started at 9000")
	http.ListenAndServe(":9000", router)
}
