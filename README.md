# Dependencies
- github.com/segmentio/ksuid
- github.com/gorilla/mux
- github.com/goombaio/namegenerator

# Local setup
- Install dependencies if needed
- Run `go run main.go` to start the server

# Endpoints
- GET `/`
    > Home endpoint returns an HTML page with the api docs

- GET `/customers`
    > Return all customers

- GET `/customer/<id>`
    > Return specific customer by ID

- POST `/customer`
    ```json
    {
        "name": string,
        "role": string,
        "email": string,
        "phone": string,
        "contacted": boolean
    }
    ```
    > Create a new customer

- POST `/customers`
    ```json
    [
        {
            "name": string,
            "role": string,
            "email": string,
            "phone": string,
            "contacted": boolean
        }
    ]
    ```
    > Create new customers in batch

- PUT `/customer/<id>`
    ```json
    {
        "name": string,
        "role": string,
        "email": string,
        "phone": string,
        "contacted": boolean
    }
    ```
    > Update existing customer or create a new one