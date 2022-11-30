package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var (
	// ErrNoDBURI is returned when the DB_URI environment variable is not set.
	ErrNoDBURI = errors.New("DB_URI not set")

	// ErrLoadEnv is returned when the .env file cannot be loaded.
	ErrLoadEnv = errors.New("error loading .env file")
)

type DB struct {
	client *mongo.Client
}

type Link struct {
	Before string
	After  string
}

func (db *DB) getURI() (uri string) {
	if err := godotenv.Load(); err != nil {
		log.Fatal(ErrLoadEnv)
	}
	uri = os.Getenv("DB_URI")
	if uri == "" {
		log.Fatal(ErrNoDBURI)
	}
	return uri
}

func (db *DB) Connect(ctx context.Context) (err error) {
	opt := options.Client().ApplyURI(db.getURI())
	if err = opt.Validate(); err != nil {
		return err
	}
	db.client, err = mongo.Connect(ctx, opt)
	return err
}

func (db *DB) Disconnect(ctx context.Context) error {
	return db.client.Disconnect(ctx)
}

func (db *DB) Ping(ctx context.Context) (err error) {
	if err = db.client.Ping(ctx, nil); err != nil {
		return err
	}
	fmt.Println("Ping to MongoDB successful")
	return nil
}

func (db *DB) GetLink(ctx context.Context, path string) (after string, err error) {
	dbName := os.Getenv("DB_NAME")
	collectionName := os.Getenv("DB_COLLECTION_NAME")
	filter := bson.D{{"before", path}}
	result := db.client.Database(dbName).Collection(collectionName).FindOne(ctx, filter)
	if err = result.Decode(&after); err != nil {
		return "", err
	}
	return after, nil
}

func (db *DB) AddLink(ctx context.Context, link Link) (err error) {
	dbName := os.Getenv("DB_NAME")
	collectionName := os.Getenv("DB_COLLECTION_NAME")
	_, err = db.client.Database(dbName).Collection(collectionName).InsertOne(ctx, link)
	return err
}
