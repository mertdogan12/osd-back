package mongo

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _client *mongo.Client
var _database *mongo.Database
var _ctx context.Context
var _cancel context.CancelFunc

func Connect() {
	mongoUrl := fmt.Sprintf("mongodb://%s:%s@%s:27017/",
		os.Getenv("MONGO_USERNAME"), os.Getenv("MONGO_PASSWORD"), os.Getenv("MONGO_HOST"))
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUrl))
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	if err = client.Connect(ctx); err != nil {
		panic(err)
	}

	_client = client
	_database = client.Database(os.Getenv("MONGO_DATABASE"))
	_ctx = ctx
	_cancel = cancel
}

func Disconnect() {
	defer _cancel()
	_client.Disconnect(_ctx)
}
