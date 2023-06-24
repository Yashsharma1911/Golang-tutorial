package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

func getAccounts(collection *mongo.Collection, ctx context.Context) (map[string]interface{}, error) {
	// Retrieve all documents from the collection
	cur, err := collection.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	defer cur.Close(ctx)

	var products []bson.M

	// Iterate through the cursor and decode each document into a bson.M
	for cur.Next(ctx) {
		var product bson.M
		if err = cur.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	// Prepare the response with the retrieved products
	res := map[string]interface{}{
		"data": products,
	}

	return res, nil // Return the response and nil error indicating successful retrieval
}
