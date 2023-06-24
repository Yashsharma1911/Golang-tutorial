package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

func updateAccount(collection *mongo.Collection, ctx context.Context, data map[string]interface{}) (map[string]interface{}, error) {
	// Define the filter to identify the document to be updated based on the provided product ID
	filter := bson.M{"product_id": data["product_id"]}

	// Define the fields to be updated using the provided data
	fields := bson.M{"$set": data}

	// Perform the update operation on the collection using the filter and fields
	_, err := collection.UpdateOne(ctx, filter, fields)

	if err != nil {
		return nil, err
	}

	// Prepare the response indicating successful update
	res := map[string]interface{}{
		"data": "Document updated successfully.",
	}

	return res, nil // Return the response and nil error indicating successful update
}
