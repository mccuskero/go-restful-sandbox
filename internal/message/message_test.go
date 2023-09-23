package message


import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventCreation(t *testing.T) {
	eventLength := 20
	events := CreateMessages(eventLength)

	for i := 0; i < eventLength; i++ {
		fmt.Printf("%v\n", events[i])
	}

	assert.Equal(t, 20, len(events), "event count doesn't match expected")
}