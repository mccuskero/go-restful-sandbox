package mongodb

import (
	"context"
	"fmt"
	"time"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/mccuskero/go-restful-sandbox/pkg/log"

)

// TODO: create mongodbdriver struct, and NewMongoDBDriver func
// pass in log

type MongoDbConnection struct {
	ctx context.Context
	opts *options.ClientOptions
	uri string
	log *log.NormalLogger
	client *mongo.Client
	timeout time.Duration
	dbConnCancelFunc context.CancelFunc
}

func NewMongoDbConnection(uri string, timeout time.Duration, log *log.NormalLogger) *MongoDbConnection {

	if uri == "" {
		uri = "mongodb://localhost:27017"
	}
	// if the operations does not finish within 10 seconds, then it will timeout
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)


	return &MongoDbConnection{
		ctx: ctx,
		opts: opts,
		uri: uri,
		log: log,
		timeout: timeout,
		dbConnCancelFunc: cancel,
	}
}

func (db *MongoDbConnection) Connect() error {

	client, err := mongo.Connect(db.ctx, db.opts)

	if err != nil {
		db.log.Info(err.Error())
		return errors.New("Could not connect to db: " + db.uri)
	}

	db.client = client

	return nil
}

func (db* MongoDbConnection) Ping() error {
	if db.client == nil {
		return errors.New("db client has not be initialized")
	}

	err := db.client.Ping(db.ctx, readpref.Primary())
	if err != nil {
		db.log.Error(err.Error())
	}

	return nil
} 

func (db* MongoDbConnection) ListDatabaseNames() {
	databases, err := db.client.ListDatabaseNames(db.ctx, bson.M{})
	if err != nil {
		db.log.Fatal(err.Error())
	}
	fmt.Println(databases)
}

func (db* MongoDbConnection) Disconnect() error {

	if db.client == nil {
		return errors.New("db client has not be initialized")
	}

	err := db.client.Disconnect(db.ctx)
	if err != nil {
		return errors.New("db client could not disconnect: " + err.Error())
	}

	return nil
}
