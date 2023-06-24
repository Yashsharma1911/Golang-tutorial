package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

const (
	dbUser     = "mongo_db_admin"   // MongoDB database user
	dbPassword = "EXAMPLE_PASSWORD" // MongoDB database password
	dbName     = "account_db"      // MongoDB database name
)

func main() {
	http.HandleFunc("/api/v1/info", accountsHandler) // Register "/api/v1/info" endpoint handler
	http.ListenAndServe(":8080", nil)                // Start the HTTP server on port 8080
}

func accountsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Set response content type to JSON

	response := map[string]interface{}{} // Response map to store the response data

	ctx := context.Background() // Create a background context

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+dbUser+":"+dbPassword+"@localhost:27017"))
	if err != nil {
		fmt.Println(err.Error())
	}

	collection := client.Database(dbName).Collection("accounts") // Get the "accounts" collection from the database

	data := map[string]interface{}{} // Map to store the JSON data from the request body

	err = json.NewDecoder(r.Body).Decode(&data) // Decode the JSON data from the request body into the data map
	if err != nil {
		fmt.Println(err.Error())
	}

	switch r.Method {
	case "GET":
		response, err = getAccounts(collection, ctx) // Call the function to handle GET request
	case "POST":
		response, err = createAccount(collection, ctx, data) // Call the function to handle POST request
	case "PUT":
		response, err = updateAccount(collection, ctx, data) // Call the function to handle PUT request
	case "DELETE":
		response, err = deleteAccount(collection, ctx, data) // Call the function to handle DELETE request
	}

	if err != nil {
		response = map[string]interface{}{"error": err.Error()} // Set an error message in the response if there's an error
	}

	enc := json.NewEncoder(w) // Create a JSON encoder for the response writer
	enc.SetIndent("", "  ")   // Set indentation for the encoded JSON

	if err := enc.Encode(response); err != nil {
		fmt.Println(err.Error())
	}
}
