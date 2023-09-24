package mongodb

import (
	"testing"

	"github.com/mccuskero/go-restful-sandbox/pkg/log"
)

func TestMongoConnection(t *testing.T) {
	logger := log.NewNormalLogger()

	conn := NewMongoDbConnection("mongodb://localhost:27017", logger)

	err := conn.Connect()

	if err != nil {
		t.Error("Could not connect to db")	
	}

	conn.Ping()

	conn.Disconnect()
}