package main

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func createAccount(collection *mongo.Collection, ctx context.Context, data map[string]interface{}) (map[string]interface{}, error) {

	req, err := collection.InsertOne(ctx, data)

	if err != nil {

		return nil, err

	}

	insertedId := req.InsertedID

	res := map[string]interface{}{

		"data": map[string]interface{}{

			"insertedId": insertedId,
		},
	}

	return res, nil

}
