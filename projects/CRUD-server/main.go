package main

import (
	"context"

	"net/http"

	"encoding/json"

	_ "log"

	"fmt"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbUser = "mongo_db_admin"

	dbPassword = "EXAMPLE_PASSWORD"

	dbName = "accounts_db"
)

func main() {
	http.HandleFunc("/api/v1/info", accountsHandler)

	http.ListenAndServe(":8080", nil)
}

func accountsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := map[string]interface{}{}

	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+dbUser+":"+dbPassword+"@localhost:27017"))

	if err != nil {
		fmt.Println(err.Error())
	}

	collection := client.Database(dbName).Collection("accounts")

	data := map[string]interface{}{}

	err = json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		fmt.Println(err.Error())
	}

	switch r.Method {
	case "GET":
		response, err = getAccounts(collection, ctx)
	case "POST":
		response, err = createAccount(collection, ctx, data)
	case "PUT":
		response, err = updateAccount(collection, ctx, data)
	case "DELETE":
		response, err = deleteAccount(collection, ctx, data)
	}

	if err != nil {
		response = map[string]interface{}{"error": err.Error()}
	}

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")

	if err := enc.Encode(response); err != nil {
		fmt.Println(err.Error())
	}
}
