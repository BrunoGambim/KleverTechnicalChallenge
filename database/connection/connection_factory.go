package connection

import (
	"context"
	"log"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientInstance *mongo.Client

var clientInstanceError error

var mongoOnce sync.Once

func GetMongoClient(context context.Context) (*mongo.Client, error) {
	mongoOnce.Do(func() {
		databaseConnection := os.Getenv("DATABASE_CONNECTION")
		clientOptions := options.Client().ApplyURI(databaseConnection)
		client, err := mongo.Connect(context, clientOptions)
		if err != nil {
			clientInstanceError = err
		}

		err = client.Ping(context, nil)
		if err != nil {
			clientInstanceError = err
		}
		clientInstance = client
		log.Print("Connection created")
	})
	return clientInstance, clientInstanceError
}
