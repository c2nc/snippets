package broker

import (
	"github.com/c2nc/snippets/pubsub"
)

// Subscription - events subscription
type Subscription struct {
	name string
	chEv chan pubsub.Event
}

// NewSubscription - events subscription constructor
func NewSubscription(name string, sz int) *Subscription {
	sub := &Subscription{
		name: name,
		chEv: make(chan pubsub.Event, sz),
	}
	return sub
}

// GetName - return subscription name
func (s *Subscription) GetName() string {
	return s.name
}

// Events - events iterator
func (s *Subscription) Events() chan pubsub.Event {
	return s.chEv
}


