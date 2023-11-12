package event_test

import (
	"context"
	"errors"
	"testing"

	"github.com/lucky-pocket/luckyPocket-back/internal/global/event"
	"github.com/stretchr/testify/assert"
)

type psuedoEventDispatcher struct{}

func (d *psuedoEventDispatcher) Dispatch(_ context.Context, _ string, payload any) error {
	if payload != nil {
		return errors.New("haha")
	}
	return nil
}

func TestEventManager(t *testing.T) {
	topic := "topic"
	m := event.NewManager()

	m.Register(topic, &psuedoEventDispatcher{})

	err := m.Publish(context.Background(), topic, nil)
	if !assert.NoError(t, err) {
		return
	}

	err = m.Publish(context.Background(), topic, 1)
	if !assert.ErrorContains(t, err, "ha") {
		return
	}

	err = m.Publish(context.Background(), topic+"1", nil)
	if !assert.ErrorContains(t, err, "topic") {
		return
	}
}
