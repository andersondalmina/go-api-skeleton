package models

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserRepository repository for users
type UserRepository struct {
	database *mongo.Client
}

// NewUserRepository creates a new repository
func NewUserRepository(database *mongo.Client) *UserRepository {
	return &UserRepository{
		database: database,
	}
}

// CreateUser save user into database
func (r *UserRepository) CreateUser(u User) (User, error) {
	collection := r.database.Database("api").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	bson := bson.M{
		"name":     u.Name,
		"email":    u.Email,
		"password": u.Password,
	}
	res, err := collection.InsertOne(ctx, bson)

	u.ID = res.InsertedID.(primitive.ObjectID).String()

	return u, err
}

// GetUserByEmail search for an user by the given email
func (r *UserRepository) GetUserByEmail(email string) (User, error) {
	collection := r.database.Database("api").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	filter := bson.M{
		"email": email,
	}

	var u User
	err := collection.FindOne(ctx, filter).Decode(&u)
	if err != nil {
		log.Fatal(err)
	}

	return u, err
}
