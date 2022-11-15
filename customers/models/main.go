package customersModels

import (
	"github.com/goombaio/namegenerator"
	"time"
)

type Customer struct {
	Name string `json:"name"`
	Role string `json:"role"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Contacted bool `json:"contacted"`
}

func GetMockedCustomers(quantity int) []Customer {
	var list []Customer

	for i := 1; i <= quantity; i++ {
		seed := time.Now().UTC().UnixNano() + int64(i)
    	nameGenerator := namegenerator.NewNameGenerator(seed)
    	name := nameGenerator.Generate()

		customer := Customer{
			name,
			"student",
			"random@email.com",
			"+55 (00) 99999-9999",
			false,
		}

		list = append(list, customer)
	}

	return list
}