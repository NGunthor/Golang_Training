package main

import (
	n "github.com/solympe/Golang_Training/pkg/pObserver/news"
	s "github.com/solympe/Golang_Training/pkg/pObserver/subscription"
)

func main() {

	newsPortal := n.NewPortal()
	subscriber := s.NewSubscriber("John")
	subscriber2 := s.NewSubscriber("Jake")

	newsPortal.AddSubscriber(subscriber)
	newsPortal.AddSubscriber(subscriber2)
	newsPortal.ShowSubscribers()

	newsPortal.DeleteSubscriber(subscriber)
	newsPortal.ShowSubscribers()

	newsPortal.Notify("Something")
}