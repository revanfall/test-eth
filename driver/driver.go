package driver

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DB struct {
	Mongo *mongo.Client
}

var dbConn = &DB{}

func ConnectMongoDB(uri string) (*DB, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	//defer func() {
	//	if err = client.Disconnect(context.TODO()); err != nil {
	//		panic(err)
	//	}
	//}()

	err = testDb(client)
	if err != nil {
		return nil, err
	}

	dbConn.Mongo = client
	return dbConn, nil
}

func testDb(client *mongo.Client) error {
	err := client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		return err
	}
	return nil
}
