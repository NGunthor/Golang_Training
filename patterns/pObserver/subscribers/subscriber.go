package subscribers

import (
	"fmt"

	sn "github.com/solympe/Golang_Training/patterns/pObserver/subscription"
)

// subscribe represents subscriber interface
type subscriber struct {
	name string
}

// GetName return subscriber`s name
func (s *subscriber) GetName() string {
	return s.name
}

// GetNotify returns recieved message
func (s *subscriber) GetNotify(message string) {
	fmt.Println("I am " + s.name + " and i received " + message)
}

// NewSubscriber returns new copy of subscriber
func NewSubscriber(name string) sn.Subscriber {
	return &subscriber{name: name}
}