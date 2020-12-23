package broker

import (
	"sync"

	"github.com/c2nc/snippets/errors"
	"github.com/c2nc/snippets/pubsub"
)

type brokerImpl struct {
	EventQueueSize int
	pubCh          chan pubsub.Event
	subCh          chan pubsub.Subscription
	unsubCh        chan pubsub.Subscription
	stopCh         chan struct{}

	mu      sync.RWMutex
	stopped bool
}

// NewPub - broker constructor
func New(qsz int) pubsub.Broker {
	return &brokerImpl{
		EventQueueSize: qsz,
		pubCh:          make(chan pubsub.Event, 1),
		subCh:          make(chan pubsub.Subscription, 1),
		unsubCh:        make(chan pubsub.Subscription, 1),
		stopCh:         make(chan struct{}),
	}
}

// Start - start broker
func (b *brokerImpl) Start() {
	subs := map[string]map[chan pubsub.Event]struct{}{}
	for {
		select {
		case _, ok := <-b.stopCh:
			if !ok {
				return
			}
		case sub := <-b.subCh:
			name := sub.GetName()
			if _, ok := subs[name]; !ok {
				subs[name] = make(map[chan pubsub.Event]struct{})
			}
			subs[name][sub.Events()] = struct{}{}
		case sub := <-b.unsubCh:
			delete(subs, sub.GetName())
		case event := <-b.pubCh:
			for subscriber := range subs[event.GetName()] {
				subscriber <- event
			}
		}
	}
}

// Subscribe - subscribe for events with events group name
func (b *brokerImpl) Subscribe(name string) (pubsub.Subscription, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	if !b.stopped {
		sub := NewSubscription(name, b.EventQueueSize)
		b.subCh <- sub
		return sub, nil
	}
	return nil, onStoppedBrokerError("Subscribe")
}

// Publish - publish a new event
func (b *brokerImpl) Publish(ev pubsub.Event) error {
	b.mu.RLock()
	defer b.mu.RUnlock()
	if !b.stopped {
		b.pubCh <- ev
		return nil
	}
	return onStoppedBrokerError("Publish")
}

// Unsubscribe - cancel a specify subscription
func (b *brokerImpl) Unsubscribe(sub pubsub.Subscription) error {
	b.mu.RLock()
	defer b.mu.RUnlock()
	if !b.stopped {
		b.unsubCh <- sub
		return nil
	}
	return onStoppedBrokerError("Unsubscribe")
}

// Stop - stop broker
func (b *brokerImpl) Stop() error {
	b.mu.Lock()
	if b.stopped {
		b.mu.Unlock()
		return onStoppedBrokerError("Stop")
	}
	b.stopped = true
	close(b.stopCh)
	b.mu.Unlock()
	return nil
}

func onStoppedBrokerError(fn string) error {
	return errors.NewfConst("try %s on stopped broker", fn)
}
