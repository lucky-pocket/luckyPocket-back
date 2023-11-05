package event

import (
	"context"
	"errors"
	"sync"
)

// Dispatcher dispatches an event to specific destination.
type Dispatcher interface {
	Dispatch(ctx context.Context, topic string, payload any) error
}

type Manager interface {
	// Register registers dispatcher to specific topic.
	Register(topic string, dst Dispatcher)

	// Publish publishes event to given topic's dispatchers.
	Publish(ctx context.Context, topic string, payload any) error
}

type manager struct {
	topics map[string][]Dispatcher
}

func NewManager() Manager {
	return &manager{
		topics: map[string][]Dispatcher{},
	}
}

func (m *manager) Register(topic string, dst Dispatcher) {
	m.topics[topic] = append(m.topics[topic], dst)
}

func (m *manager) Publish(ctx context.Context, topic string, payload any) (err error) {
	dispatchers, ok := m.topics[topic]
	if !ok {
		return errors.New("topic not found")
	}

	var errs = make(chan error, len(dispatchers))
	go func() {
		var wg sync.WaitGroup
		defer close(errs)

		wg.Add(len(dispatchers))

		for _, d := range dispatchers {
			go func(d Dispatcher) {
				defer wg.Done()
				if e := d.Dispatch(ctx, topic, payload); e != nil {
					errs <- e
				}
			}(d)
		}

		wg.Wait()
	}()

	for e := range errs {
		if e != nil {
			err = errors.Join(err, e)
		}
	}

	return
}
