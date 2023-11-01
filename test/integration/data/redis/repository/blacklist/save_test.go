package blacklist_test

import (
	"context"
	"time"
)

func (s *BlacklistRepositoryTestSuite) TestSave() {
	err := s.r.Save(context.Background(), "hi", time.Now().Add(time.Hour))
	s.NoError(err)

	cmd := s.client.Get(context.Background(), "blacklist:hi")

	s.NoError(cmd.Err())
}
