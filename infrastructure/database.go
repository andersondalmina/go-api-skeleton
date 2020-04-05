package infrastructure

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateDatabase create a new database instance
func CreateDatabase() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://database:27017"))
	if err != nil {
		log.Fatal("Error on creating database!")
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Error on creating context for database!")
	}

	return client
}
