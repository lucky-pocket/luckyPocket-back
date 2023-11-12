package batch_test

import (
	"testing"
	"time"

	"github.com/lucky-pocket/luckyPocket-back/internal/global/batch"
	"github.com/stretchr/testify/assert"
)

type pseudoProcessor struct{}

func (p *pseudoProcessor) Do() {
	time.Sleep(500 * time.Millisecond)
}

func TestScheduler(t *testing.T) {
	s := batch.NewScheduler(time.Local)

	err := s.Register(time.Second, &pseudoProcessor{})
	if !assert.NoError(t, err) {
		return
	}

	job := s.Jobs[0]
	s.Start()
	defer s.Stop()

	start := time.Now()
	time.Sleep(time.Until(job.NextRun()))

	for job.IsRunning() {
	}

	assert.True(t, time.Since(start) > 500*time.Millisecond)
}
