package main

import (
	"database/sql"
	"encoding/json"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gofr.dev/pkg/gofr"
)

type Customer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	// Initialize database connection
	db, err := sql.Open("mysql", "kkavenger:mishra@tcp(localhost:3306)/test_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a gofr application
	app := gofr.New()

	// HTTP GET Route for Greeting
	app.GET("/greet", func(ctx *gofr.Context) (interface{}, error) {
		// Get the value using the redis instance
		value, err := ctx.Redis.Get(ctx.Context, "greeting").Result()
		return value, err
	})

	// HTTP POST Route for Adding Customer
	app.POST("/customer/{name}", func(ctx *gofr.Context) (interface{}, error) {
		name := ctx.PathParam("name")

		// Inserting a customer row in the database using SQL
		_, err := db.ExecContext(ctx, "INSERT INTO customers (name) VALUES (?)", name)
		if err != nil {
			return nil, err
		}

		return "Customer added successfully", nil
	})

	// HTTP GET Route for Retrieving Customers
	app.GET("/customer", func(ctx *gofr.Context) (interface{}, error) {
		var customers []Customer

		// Getting customers from the database using SQL
		rows, err := db.QueryContext(ctx, "SELECT * FROM customers")
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var customer Customer
			if err := rows.Scan(&customer.ID, &customer.Name); err != nil {
				return nil, err
			}

			customers = append(customers, customer)
		}

		// Return customers as JSON
		jsonData, err := json.Marshal(customers)
		if err != nil {
			return nil, err
		}

		return string(jsonData), nil
	})

	// Starts the server, it will listen on the default port 8000.
	// It can be overridden through the configs
	app.Start()
}
