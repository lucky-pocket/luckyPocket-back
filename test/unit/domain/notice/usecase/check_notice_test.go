package usecase_test

import (
	"context"
	"net/http"
	"time"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/error/status"
	"github.com/stretchr/testify/mock"
)

func (s *NoticeUseCaseTestSuite) TestCheckNotice() {
	userInfo := auth.Info{
		UserID: 1,
		Role:   constant.RoleMember,
	}

	myNotice := domain.Notice{
		NoticeID:  1,
		UserID:    userInfo.UserID,
		Type:      constant.NoticeTypeReceived,
		Checked:   true,
		CreatedAt: time.Now(),
	}

	otherNotice := domain.Notice{
		NoticeID:  1,
		UserID:    userInfo.UserID + 1,
		Type:      constant.NoticeTypeReceived,
		Checked:   true,
		CreatedAt: time.Now(),
	}

	testcases := []struct {
		desc   string
		on     func()
		assert func(err error)
	}{
		{
			desc: "success",
			on: func() {
				s.mockNoticeRepository.On("FindByID", mock.Anything, mock.Anything).Return(&myNotice, nil).Once()
				s.mockNoticeRepository.On("SetChecked", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
			},
			assert: func(err error) {
				s.Nil(err)
			},
		},
		{
			desc: "not found",
			on: func() {
				s.mockNoticeRepository.On("FindByID", mock.Anything, mock.Anything).Return(nil, nil).Once()
			},
			assert: func(err error) {
				e, ok := err.(*status.Err)
				if s.True(ok) {
					s.Equal(http.StatusNotFound, e.Code)
				}
			},
		},
		{
			desc: "not mine",
			on: func() {
				s.mockNoticeRepository.On("FindByID", mock.Anything, mock.Anything).Return(&otherNotice, nil).Once()
			},
			assert: func(err error) {
				e, ok := err.(*status.Err)
				if s.True(ok) {
					s.Equal(http.StatusForbidden, e.Code)
				}
			},
		},
	}

	ctx := auth.Inject(context.Background(), userInfo)
	for _, tc := range testcases {
		s.Run(tc.desc, func() {
			tc.on()

			err := s.uc.CheckNotice(ctx, 1)

			tc.assert(err)

			s.mockNoticeRepository.AssertExpectations(s.T())
		})
	}
}
