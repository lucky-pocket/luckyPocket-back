package notice_test

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
)

func (s *NoticeRepositoryTestSuite) TestExistsByUserID() {
	userID, pocketID := s.createUserAndPocket()

	err := s.client.Notice.Create().
		SetChecked(false).
		SetPocketID(pocketID).
		SetType(constant.NoticeTypeReceived).
		SetUserID(userID).Exec(context.Background())

	if !s.NoError(err) {
		return
	}

	s.Run("found", func() {
		exists, err := s.r.ExistsByUserID(context.Background(), userID)

		if s.NoError(err) {
			s.True(exists)
		}
	})
	s.Run("not found", func() {
		exists, err := s.r.ExistsByUserID(context.Background(), userID+1)

		if s.NoError(err) {
			s.False(exists)
		}
	})
}
