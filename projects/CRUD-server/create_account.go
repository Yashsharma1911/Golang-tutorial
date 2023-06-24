package main

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func createAccount(collection *mongo.Collection, ctx context.Context, data map[string]interface{}) (map[string]interface{}, error) {
	// Insert the provided data into the specified collection
	req, err := collection.InsertOne(ctx, data)

	if err != nil {
		return nil, err
	}

	insertedId := req.InsertedID // Get the inserted ID from the request response

	// Prepare the response with the inserted ID
	res := map[string]interface{}{
		"data": map[string]interface{}{
			"insertedId": insertedId,
		},
	}

	return res, nil // Return the response and nil error indicating successful insertion
}
