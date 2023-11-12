package ticket_test

import (
	"context"
	"time"
)

func (s *TicketRepositoryTestSuite) TestGetRefillAt() {
	cmd := s.client.Set(context.Background(), "ticket:3", 1, time.Hour)
	if !s.NoError(cmd.Err()) {
		return
	}

	s.Run("found", func() {
		_, err := s.r.GetRefillAt(context.Background(), 3)
		if !s.NoError(err) {
			return
		}
	})

	s.Run("not found", func() {
		refillAt, err := s.r.GetRefillAt(context.Background(), 1)
		if !s.NoError(err) {
			return
		}

		s.Equal(
			time.Now().Truncate(time.Minute),
			refillAt.Truncate(time.Minute),
		)
	})
}
