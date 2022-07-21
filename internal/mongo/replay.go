package mongo

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveReplay(uploader int, player int, uuid uuid.UUID) {
	if _database == nil {
		panic(errors.New("You are not connected (use mongo.Connect() to Connect))"))
	}

	coll := _database.Collection("groups")

	filter := bson.D{primitive.E{Key: "id", Value: uploader}}
	update := bson.D{primitive.E{Key: "$push",
		Value: bson.D{primitive.E{Key: "uploadedReplays", Value: uuid.String()}}}}

	coll.UpdateOne(context.TODO(), filter, update)
}
