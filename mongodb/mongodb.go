package mongodb

import (
	"context"
	"log"

	"github.com/escaletech/go-escale/messages"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

func GetDatabase(mongoURL string) (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(mongoURL)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	cs, err := connstring.Parse(mongoURL)
	if err != nil {
		log.Fatal(err)
	}

	return client.Database(cs.Database), nil
}

func EnsureIndex(collection *mongo.Collection, indexName string, keys bson.M, unique bool) error {
	if indexExists(collection, indexName) {
		return nil
	}

	opt := options.Index()
	opt.SetUnique(unique)
	opt.SetName(indexName)
	index := mongo.IndexModel{Keys: keys, Options: opt}

	_, err := collection.Indexes().CreateOne(context.Background(), index)

	return err
}

func indexExists(collection *mongo.Collection, indexName string) bool {
	cursor, err := collection.Indexes().List(context.Background())

	if err != nil {
		log.Println(messages.CouldNotGetIndexes, err)
		return false
	}

	var results []bson.M
	if err = cursor.All(context.Background(), &results); err != nil {
		log.Println(messages.CouldNotGetIndexes, err)
		return false
	}

	for _, index := range results {
		if index["name"] == indexName {
			return true
		}
	}

	return false
}

func IsErrorDuplicateKey(err error) bool {
	mongoError := err.(mongo.WriteException)
	errCode := mongoError.WriteErrors[0].Code

	return errCode == 11000
}

func IsErrorNotFound(err error) bool {
	return err == mongo.ErrNoDocuments
}
