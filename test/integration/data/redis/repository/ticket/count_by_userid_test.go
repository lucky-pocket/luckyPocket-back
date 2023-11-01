package ticket_test

import (
	"context"
	"time"
)

func (s *TicketRepositoryTestSuite) TestCountByUserID() {
	cmd := s.client.Set(context.Background(), "ticket:3", 2, time.Hour)
	if !s.NoError(cmd.Err()) {
		return
	}

	s.Run("found", func() {
		count, err := s.r.CountByUserID(context.Background(), 3)
		if !s.NoError(err) {
			return
		}
		s.Equal(2, count)
	})

	s.Run("not found", func() {
		count, err := s.r.CountByUserID(context.Background(), 1)
		if !s.NoError(err) {
			return
		}

		s.Equal(0, count)
	})
}
