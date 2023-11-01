package ticket_test

import (
	"context"
	"time"
)

func (s *TicketRepositoryTestSuite) TestIncrease() {
	cmd := s.client.Set(context.Background(), "ticket:3", 2, time.Hour)
	if !s.NoError(cmd.Err()) {
		return
	}

	s.Run("found", func() {
		err := s.r.Increase(context.Background(), 3)
		if !s.NoError(err) {
			return
		}

		cmd := s.client.Get(context.Background(), "ticket:3")
		if !s.NoError(cmd.Err()) {
			return
		}

		val, _ := cmd.Int()
		s.Equal(3, val)
	})

	s.Run("not found", func() {
		err := s.r.Increase(context.Background(), 1)
		if !s.NoError(err) {
			return
		}

		cmd := s.client.Get(context.Background(), "ticket:1")
		if !s.NoError(cmd.Err()) {
			return
		}

		val, _ := cmd.Int()
		s.Equal(1, val)
	})
}
