package batch

import (
	"time"

	"github.com/go-co-op/gocron"
	"github.com/pkg/errors"
)

// Processor processes batch work.
type Processor interface{ Do() }

// Scheduler schedules tasks.
// You have to call Start to start all jobs.
// And you should always call Stop before program exits.
type Scheduler struct {
	s    *gocron.Scheduler
	Jobs []*gocron.Job
}

func NewScheduler(loc *time.Location) *Scheduler {
	return &Scheduler{
		s:    gocron.NewScheduler(loc),
		Jobs: make([]*gocron.Job, 0),
	}
}

func (s *Scheduler) Start() { s.s.StartAsync() }

func (s *Scheduler) Stop() { s.s.Stop() }

// Register registers task with given duration.
func (s *Scheduler) Register(duration time.Duration, p Processor) error {
	job, err := s.s.Every(duration).
		StartAt(
			time.Now().Truncate(time.Minute).Add(time.Minute),
		).WaitForSchedule().Do(p.Do)
	if err != nil {
		return errors.Wrap(err, "job registration failed")
	}

	s.Jobs = append(s.Jobs, job)

	return nil
}
