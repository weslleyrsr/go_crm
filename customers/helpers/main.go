package customersHelpers

import (
	"encoding/json"
	"net/http"
	"fmt"
	customersModels "github.com/weslleyrsr/go_crm/customers/models"
	"github.com/gorilla/mux"
	"github.com/segmentio/ksuid"
	"strconv"
	"io/ioutil"
)

var customers = make(map[string]customersModels.Customer)

func InitMock() {
	for index, customer := range customersModels.GetMockedCustomers(3) {
		saveCustomer(customer, strconv.Itoa(index))
	}

	fmt.Println("Mock data loaded")
}

func GetCostumers(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(customers)
}

func GetCostumer(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(req)
    id := vars["id"]

	if _, ok := customers[id]; ok {
		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(customers[id])
	} else {
		res.WriteHeader(http.StatusNotFound)
		json.NewEncoder(res).Encode("Not found")
	}
}

func AddCustomer(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var customer customersModels.Customer

	reqBody, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(reqBody, &customer)

	// err := json.NewDecoder(req.Body).Decode(&customer)
    // if err != nil {
    //     http.Error(res, err.Error(), http.StatusBadRequest)
    //     return
    // }

	saveCustomer(customer, "")

	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(customers)
}

func AddCustomers(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var customersList []customersModels.Customer

	reqBody, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(reqBody, &customersList)

	// err := json.NewDecoder(req.Body).Decode(&customersList)
    // if err != nil {
    //     http.Error(res, err.Error(), http.StatusBadRequest)
    //     return
    // }

	saveCustomers(customersList)

	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(customers)
}

func UpdateCostumer(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(req)
    id := vars["id"]

	var customer customersModels.Customer

	reqBody, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(reqBody, &customer)

	// err := json.NewDecoder(req.Body).Decode(&customer)
    // if err != nil {
    //     http.Error(res, err.Error(), http.StatusBadRequest)
    //     return
    // }

	customers[id] = customer

	json.NewEncoder(res).Encode(customers)
}

func DeleteCustomer(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(req)
	id := vars["id"]

	if _, ok := customers[id]; ok {
		delete(customers, id)
		res.WriteHeader(http.StatusOK)
	} else {
		res.WriteHeader(http.StatusNotFound)
	}

	json.NewEncoder(res).Encode(customers)
}

func saveCustomer(customer customersModels.Customer, id string) {
	if(id == "") {
		id = ksuid.New().String()
	}

	customers[id] = customer
}

func saveCustomers(customers []customersModels.Customer) {
	for _, customer := range customers {
		saveCustomer(customer, "")
	}
}