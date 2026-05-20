package mongodb

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoStore struct {
	Client          *mongo.Client
	NotesCollection *mongo.Collection
	UsersCollection *mongo.Collection
}

func Connect() (*MongoStore, error) {

	fmt.Println("connection url", os.Getenv("DB_CONNECTION_URL"))

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(
		options.Client().ApplyURI(os.Getenv("DB_CONNECTION_URL")),
	)

	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		return nil, err
	}

	notesCollection := client.Database(
		os.Getenv("DB_NAME"),
	).Collection(os.Getenv("NOTE_COLLECTION_NAME"))

	usersCollection := client.Database(
		os.Getenv("DB_NAME"),
	).Collection(os.Getenv("USER_COLLECTION_NAME"))

	store := &MongoStore{
		Client:          client,
		NotesCollection: notesCollection,
		UsersCollection: usersCollection,
	}

	return store, nil
}
