package configs

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DBCollection name constants
const (
	TodosCollection = "todos"
)

// MongoDB connection variables
var (
	Client     *mongo.Client
	TodoDB     *mongo.Database
	Collection *mongo.Collection
)

// getEnv gets an environment variable or returns the default value if not set
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// MongoDB connection config variables
var (
	// MongoDB connection string from environment variable or default
	mongoURI = getEnv("MONGO_URI", "mongodb://localhost:27017")
	// Database name from environment variable or default
	dbName = getEnv("MONGO_DB_NAME", "todoDB")
)

// ConnectDB connects to MongoDB
func ConnectDB() {
	// Set client options
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	Client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = Client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Connected to MongoDB at %s using database %s!\n", mongoURI, dbName)

	// Get database and collection handles
	TodoDB = Client.Database(dbName)
	Collection = TodoDB.Collection(TodosCollection)
}

// GetCollection returns the specified MongoDB collection
func GetCollection(collectionName string) *mongo.Collection {
	return TodoDB.Collection(collectionName)
}

// DisconnectDB closes the MongoDB connection
func DisconnectDB() {
	if Client == nil {
		return
	}
	
	err := Client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection to MongoDB closed.")
}