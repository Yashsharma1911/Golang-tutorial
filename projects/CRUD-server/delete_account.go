package main

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

func deleteAccount(collection *mongo.Collection, ctx context.Context, data map[string]interface{}) (map[string]interface{}, error) {
	// Delete the document in the collection that matches the provided product ID
	_, err := collection.DeleteOne(ctx, bson.M{"product_id": data["product_id"]})

	if err != nil {
		return nil, err
	}

	// Prepare the response indicating successful deletion
	res := map[string]interface{}{
		"data": "Document deleted successfully.",
	}

	return res, nil // Return the response and nil error indicating successful deletion
}
