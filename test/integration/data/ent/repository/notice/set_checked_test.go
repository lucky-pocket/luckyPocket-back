package notice_test

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
)

func (s *NoticeRepositoryTestSuite) TestSetChecked() {
	userID, pocketID := s.createUserAndPocket()

	notice, err := s.client.Notice.Create().
		SetChecked(false).
		SetPocketID(pocketID).
		SetType(constant.NoticeTypeReceived).
		SetUserID(userID).Save(context.Background())

	if !s.NoError(err) {
		return
	}

	err = s.r.SetChecked(context.Background(), notice.ID, true)
	if !s.NoError(err) {
		return
	}

	found, err := s.client.Notice.Get(context.Background(), notice.ID)
	if !s.NoError(err) {
		return
	}

	s.True(found.Checked)
}
