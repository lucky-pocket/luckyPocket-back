package usecase_test

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth"
	"github.com/stretchr/testify/mock"
)

func (s *NoticeUseCaseTestSuite) TestCheckAllNotices() {
	userInfo := auth.Info{
		UserID: 1,
		Role:   constant.RoleMember,
	}

	testcases := []struct {
		desc   string
		on     func()
		assert func(err error)
	}{
		{
			desc: "success",
			on: func() {
				s.mockNoticeRepository.On("SetCheckedByUserID", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
			},
			assert: func(err error) {
				s.Nil(err)
			},
		},
	}

	ctx := auth.Inject(context.Background(), userInfo)
	for _, tc := range testcases {
		s.Run(tc.desc, func() {
			tc.on()

			err := s.uc.CheckAllNotices(ctx)

			tc.assert(err)

			s.mockNoticeRepository.AssertExpectations(s.T())
		})
	}
}
