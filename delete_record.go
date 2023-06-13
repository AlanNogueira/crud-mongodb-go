package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func deleteRecord(collection *mongo.Collection, ctx context.Context, data map[string]interface{}) (map[string]interface{}, error) {
	_, err := collection.DeleteOne(ctx, bson.M{"product_id": data["product_id"]})

	if err != nil {
		return nil, err
	}

	res := map[string]interface{}{
		"data": "Document deleted socessfully",
	}

	return res, nil
}
