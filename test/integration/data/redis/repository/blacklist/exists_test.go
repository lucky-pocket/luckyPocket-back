package blacklist_test

import (
	"context"
	"time"
)

func (s *BlacklistRepositoryTestSuite) TestExists() {
	cmd := s.client.Set(context.Background(), "blacklist:hi", nil, time.Hour)

	if s.NoError(cmd.Err()) {
		exists, err := s.r.Exists(context.Background(), "hi")
		if s.NoError(err) {
			s.True(exists)
		}
	}
}
