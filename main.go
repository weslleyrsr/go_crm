package main

import (
	"github.com/gorilla/mux"
	customersRoutes "github.com/weslleyrsr/go_crm/customers/routes"
	customersHelpers "github.com/weslleyrsr/go_crm/customers/helpers"
	"fmt"
	"net/http"
)

func apiDocs(res http.ResponseWriter, req *http.Request) {
    http.ServeFile(res, req, "README.html")  
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", apiDocs)

	customersHelpers.InitMock()
	customersRoutes.LoadRoutes("", router)

	fmt.Println("Starting server at port 3000")
	http.ListenAndServe(":3000", router)
}