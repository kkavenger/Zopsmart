# Gofr Customer Management API

This project implements a simple RESTful API for managing customer data using the Gofr web framework and MySQL database.

## Functionalities

1. **Add Customer:** Add a new customer to the database by sending a POST request.

2. **Get Customers:** Retrieve a list of all customers from the database by sending a GET request.

3. **Update Customer:** Update a customer's name in the database by sending a PUT request with the customer ID and new name.

4. **Delete Customer:** Delete a customer from the database by sending a DELETE request with the customer ID.

## Prerequisites

- [Go](https://golang.org/dl/) installed on your machine.
- [Docker](https://www.docker.com/get-started) for running the MySQL database.

## Setup

1. Clone the repository:
   git clone https://github.com/your-username/gofr-customer-api.git
   cd gofr-customer-API
2.Start the MySQL Docker container:
    docker run --name mysql-container -e MYSQL_ROOT_PASSWORD=your_password -e MYSQL_DATABASE=customer_db -p 3311:3306 -d mysql:8.0.30
    Replace your_password with your desired MySQL root password.

3.Run the Go application:
    go run main.go
    The API will be accessible at http://localhost:8000.

## API Endpoints
  GET /greet: Greet endpoint (sample).
  POST /customer/{name}: Add a new customer.
  GET /customer: Retrieve all customers.
  PUT /customer/{id}/{name}: Update a customer's name.
  DELETE /customer/{id}: Delete a customer.

##Example Curl Commands:
Add Customer:
    curl --location --request POST 'http://localhost:8000/customer/JohnDoe'
    
Get Customers:
    curl --location --request GET 'http://localhost:8000/customer'
    
Update Customer:
    curl --location --request PUT 'http://localhost:8000/customer/1/NewName'
    
Delete Customer:
curl --location --request DELETE 'http://localhost:8000/customer/1'
