package customersRoutes

import (
	"github.com/gorilla/mux"
	customersHelpers "github.com/weslleyrsr/go_crm/customers/helpers"
	"fmt"
)

func LoadRoutes(path string, router *mux.Router) {
	router.HandleFunc(path + "/customers", customersHelpers.GetCostumers).Methods("GET")
	router.HandleFunc(path + "/customers", customersHelpers.AddCustomers).Methods("POST")
	
	router.HandleFunc(path + "/customer", customersHelpers.AddCustomer).Methods("POST")
	router.HandleFunc(path + "/customer/{id}", customersHelpers.GetCostumer).Methods("GET")
	router.HandleFunc(path + "/customer/{id}", customersHelpers.DeleteCustomer).Methods("DELETE")
	router.HandleFunc(path + "/customer/{id}", customersHelpers.UpdateCostumer).Methods("PUT")

	fmt.Println("Routes loaded")
}