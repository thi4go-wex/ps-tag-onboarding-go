package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DBURI            = "mongodb://localhost:27017"
	DBName           = "user-service-test"
	DBUserCollection = "user"
)

type MongoDB struct {
	db *mongo.Database
}

func NewMongoDB() (*MongoDB, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(DBURI))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	return &MongoDB{
		db: client.Database(DBName),
	}, nil
}

func (m *MongoDB) Get(table, id string) (interface{}, error) {
	var result interface{}
	err := m.db.Collection(table).FindOne(context.Background(), bson.M{"_id": id}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (m *MongoDB) Set(table string, data interface{}) error {
	_, err := m.db.Collection(table).UpdateOne(context.Background(), bson.M{}, bson.M{"$set": data}, options.Update().SetUpsert(true))
	return err
}
