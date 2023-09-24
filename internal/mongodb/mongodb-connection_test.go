package mongodb

import (
	"testing"
	"time"

	"github.com/mccuskero/go-restful-sandbox/pkg/log"
)

func TestMongoConnection(t *testing.T) {
	logger := log.NewNormalLogger()

	conn := NewMongoDbConnection("mongodb://localhost:27017", 10*time.Second, logger)

	err := conn.Connect()

	if err != nil {
		t.Error("Could not connect to db")	
	}

	if err := conn.Ping(); err != nil {
		t.Error("Could not ping db")
	}

	if err := conn.Disconnect(); err != nil {
		t.Error("Could not disconnect from db")
	}
}