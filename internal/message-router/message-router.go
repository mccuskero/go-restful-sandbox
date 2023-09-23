package messagerouter

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mccuskero/go-restful-sandbox/internal/message"
)


var testMessages []*message.Message

func initializeTestMessages() []*message.Message {
	return message.CreateMessages(100)
}

func InitializeAndRun() {
	// TODO: create a MessageRouter struct, and NewMessageRouter, passing in logger
	// make an Initialize, and Run as method sets

	router := gin.Default()
	testMessages = initializeTestMessages()

	router.GET("/messages", getMessages)
	router.POST("/postMessage", postMessage)

	router.Run()
}

func getMessages(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, testMessages)
}

func postMessage(c *gin.Context) {
	var newMessage message.Message

	// bind the received JSON to Message
	if err := c.BindJSON(&newMessage); err != nil {
		// return error code to user
		return
	}

	// add new Messages to messages
	testMessages = append(testMessages, &newMessage)
	c.IndentedJSON(http.StatusCreated, newMessage)
}



