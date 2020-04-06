package infrastructure

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateDatabase create a new database instance
func CreateDatabase() *mongo.Client {
	databaseURL := fmt.Sprintf("mongodb://%s:%s", os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"))

	client, err := mongo.NewClient(options.Client().ApplyURI(databaseURL))
	if err != nil {
		log.Fatal("Error on connecting to database!")
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Error on creating context for database!")
	}

	return client
}
