package usecase_test

import (
	"context"
	"net/http"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/input"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/error/status"
	"github.com/stretchr/testify/mock"
)

func (s *PocketUseCaseTestSuite) TestRevealSender() {
	userInfo := auth.Info{
		UserID: 2,
		Role:   constant.RoleMember,
	}

	myPocket := domain.Pocket{
		PocketID: 123,
		Receiver: &domain.User{
			UserID: userInfo.UserID,
		},
		ReceiverID: 2,
		SenderID:   3,
	}

	otherPocket := domain.Pocket{
		PocketID: 1531,
		Receiver: &domain.User{
			UserID: userInfo.UserID + 1,
		},
		ReceiverID: 3,
		SenderID:   1,
	}

	testcases := []struct {
		desc   string
		input  *input.PocketIDInput
		on     func()
		assert func(err error)
	}{
		{
			desc:  "success",
			input: &input.PocketIDInput{},
			on: func() {
				s.mockPocketRepository.On("FindByID", mock.Anything, mock.Anything).Return(&myPocket, nil).Once()
				s.mockUserRepository.On("CountCoinsByUserID", mock.Anything, mock.Anything).Return(constant.CostRevealSender+1, nil).Once()
				s.mockPocketRepository.On("RevealExists", mock.Anything, mock.Anything, mock.Anything).Return(false, nil).Once()
				s.mockPocketRepository.On("CreateReveal", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
				s.mockUserRepository.On("UpdateCoin", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
				s.mockEventManager.On("Publish", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
			},
			assert: func(err error) {
				s.Nil(err)
			},
		},
		{
			desc:  "pocket not found",
			input: &input.PocketIDInput{},
			on: func() {
				s.mockPocketRepository.On("FindByID", mock.Anything, mock.Anything).Return(nil, nil).Once()
			},
			assert: func(err error) {
				e, ok := err.(*status.Err)
				if s.True(ok) {
					s.Equal(http.StatusNotFound, e.Code)
				}
			},
		},
		{
			desc:  "no permission to the pocket",
			input: &input.PocketIDInput{},
			on: func() {
				s.mockPocketRepository.On("FindByID", mock.Anything, mock.Anything).Return(&otherPocket, nil).Once()
			},
			assert: func(err error) {
				e, ok := err.(*status.Err)
				if s.True(ok) {
					s.Equal(http.StatusForbidden, e.Code)
				}
			},
		},
		{
			desc:  "already revealed",
			input: &input.PocketIDInput{},
			on: func() {
				s.mockPocketRepository.On("FindByID", mock.Anything, mock.Anything).Return(&myPocket, nil).Once()
				s.mockUserRepository.On("CountCoinsByUserID", mock.Anything, mock.Anything).Return(constant.CostRevealSender+1, nil).Once()
				s.mockPocketRepository.On("RevealExists", mock.Anything, mock.Anything, mock.Anything).Return(true, nil).Once()
			},
			assert: func(err error) {
				e, ok := err.(*status.Err)
				if s.True(ok) {
					s.Equal(http.StatusConflict, e.Code)
				}
			},
		},
		{
			desc:  "not enough coins",
			input: &input.PocketIDInput{},
			on: func() {
				s.mockPocketRepository.On("FindByID", mock.Anything, mock.Anything).Return(&myPocket, nil).Once()
				s.mockUserRepository.On("CountCoinsByUserID", mock.Anything, mock.Anything).Return(constant.CostRevealSender-1, nil).Once()
				s.mockPocketRepository.On("RevealExists", mock.Anything, mock.Anything, mock.Anything).Return(false, nil).Once()
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

			err := s.uc.RevealSender(ctx, tc.input)

			tc.assert(err)

			s.mockPocketRepository.AssertExpectations(s.T())
			s.mockUserRepository.AssertExpectations(s.T())
			s.mockEventManager.AssertExpectations(s.T())
		})
	}
}
