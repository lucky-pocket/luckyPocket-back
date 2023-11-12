package usecase_test

import (
	"context"
	"time"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output/mapper"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth"
	"github.com/stretchr/testify/mock"
)

func (s *NoticeUseCaseTestSuite) TestGetNotice() {
	userInfo := auth.Info{
		UserID: 1,
		Role:   constant.RoleMember,
	}

	notices := []*domain.Notice{
		{
			NoticeID:  1,
			User:      &domain.User{UserID: 1},
			Pocket:    &domain.Pocket{PocketID: 1},
			Type:      constant.NoticeTypeReceived,
			Checked:   true,
			CreatedAt: time.Now(),
		},
	}

	testcases := []struct {
		desc   string
		on     func()
		assert func(output *output.NoticeListOutput, err error)
	}{
		{
			desc: "success",
			on: func() {
				s.mockNoticeRepository.On("FindAllByUserID", mock.Anything, mock.Anything).Return(notices, nil).Once()
			},
			assert: func(output *output.NoticeListOutput, err error) {
				if s.Nil(err) {
					s.Equal(mapper.ToNoticeListOutPut(notices), output)
				}
			},
		},
	}

	ctx := auth.Inject(context.Background(), userInfo)
	for _, tc := range testcases {
		s.Run(tc.desc, func() {
			tc.on()

			output, err := s.uc.GetNotice(ctx)

			tc.assert(output, err)

			s.mockNoticeRepository.AssertExpectations(s.T())
		})
	}
}
