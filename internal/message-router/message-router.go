package messagerouter

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mccuskero/go-restful-sandbox/internal/message"
	"github.com/mccuskero/go-restful-sandbox/internal/mongodb"
)


type MessageRouter struct {
	dbConn       *mongodb.MongoDbConnection
	router    *gin.Engine
	testMode     bool
	testMessages []*message.Message	
}

func initializeTestMessages() []*message.Message {
	return message.CreateMessages(100)
}

func NewMessageRouter(dbConn *mongodb.MongoDbConnection, testMode bool) *MessageRouter {

	router := gin.Default()
	testMessages := initializeTestMessages()

	messageRouter := &MessageRouter{
		dbConn: dbConn,
		router: router,
		testMessages: testMessages,
		testMode: testMode,
	}
	
	return messageRouter
} 

func (mr* MessageRouter) InitializeRoutes() {
	mr.router.GET("/messages", mr.getMessages)
	mr.router.POST("/postMessage", mr.postMessage)
}

func (mr* MessageRouter) Run() error {

	if err := mr.router.Run(); err != nil {
		return errors.New("Could not start message router: " + err.Error())
	}

	return nil
} 


func (mr* MessageRouter) getMessages(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, mr.testMessages)
}

func (mr* MessageRouter) postMessage(c *gin.Context) {
	var newMessage message.Message

	// bind the received JSON to Message
	if err := c.BindJSON(&newMessage); err != nil {
		// return error code to user
		return
	}

	// add new Messages to messages
	mr.testMessages = append(mr.testMessages, &newMessage)
	c.IndentedJSON(http.StatusCreated, newMessage)
}



