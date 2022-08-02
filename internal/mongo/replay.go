package mongo

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func SaveReplay(uploader int, player int, replayHash string) (*mongo.UpdateResult, *mongo.UpdateResult, error) {
	if _database == nil {
		return nil, nil, errors.New("You are not connected (use mongo.Connect() to Connect))")
	}

	coll := _database.Collection("users")

	// uploader
	filter := bson.D{primitive.E{Key: "id", Value: uploader}}
	update := bson.D{primitive.E{Key: "$push",
		Value: bson.D{primitive.E{Key: "uploadedReplays", Value: replayHash}}}}

	uploaderRes, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, nil, err
	}

	// player
	filter = bson.D{primitive.E{Key: "id", Value: player}}
	update = bson.D{primitive.E{Key: "$push",
		Value: bson.D{primitive.E{Key: "replays", Value: replayHash}}}}

	playerRes, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, nil, err
	}

	// Creates the player and saves uuid again
	if playerRes.ModifiedCount == 0 {
		// TODO create player
		_, playerRes, err = SaveReplay(0, player, replayHash)
		if err != nil {
			return nil, nil, err
		}
	}

	return uploaderRes, playerRes, nil
}
