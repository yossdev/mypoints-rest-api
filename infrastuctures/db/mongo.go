package db

import (
	"context"
	"github.com/yossdev/mypoints-rest-api/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type MongoDB interface {
	DB() *mongo.Client
}

type mongoDB struct {
	client *mongo.Client
}

func NewMongoClient() MongoDB {
	var client *mongo.Client
	// Set client options
	//clientOptions := options.Client().ApplyURI(configs.Get().MongodbAddress) //for local connection
	clientOptions := options.Client().ApplyURI("mongodb+srv://" + configs.Get().MongodbUsername + ":" + configs.Get().MongodbPassword + "@cluster0.atngo.mongodb.net/" + configs.Get().MongodbName + "?retryWrites=true&w=majority")

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}

	return &mongoDB{
		client: client,
	}
}

func (c mongoDB) DB() *mongo.Client {
	return c.client
}
