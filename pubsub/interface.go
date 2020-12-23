package pubsub

import (
	"github.com/c2nc/snippets/cmp"
)

// Broker - events broker
type Broker interface {
	// Start - start broker
	Start()
	// Subscribe - subscribe for events with events group name
	Subscribe(name string) (sub Subscription, err error)
	// Publish - publish a new event
	Publish(ev Event) error
	// Unsubscribe - cancel a specify subscription
	Unsubscribe(sub Subscription) error
	// Stop - stop broker
	Stop() error
}

// Subscription - events subscription
type Subscription interface {
	// GetName - return subscription name
	GetName() string
	// Events - events iterator
	Events() chan Event
}

// Event - something event
type Event interface {
	cmp.Comparable
	GetName() string
}
