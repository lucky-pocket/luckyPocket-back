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

func (s *PocketUseCaseTestSuite) TestSendPocket() {
	userInfo := auth.Info{
		UserID: 1,
		Role:   constant.RoleMember,
	}

	testcases := []struct {
		desc   string
		input  *input.PocketInput
		on     func()
		assert func(err error)
	}{
		{
			desc:  "success (not public)",
			input: &input.PocketInput{ReceiverID: 2, Coins: 1, IsPublic: false},
			on: func() {
				s.mockPocketRepository.On("CountBySenderIdAndReceiverId", mock.Anything, mock.Anything, mock.Anything).Return(0, nil).Once()
				s.mockUserRepository.On("FindByID", mock.Anything, mock.Anything).Return(&domain.User{}, nil).Once()
				s.mockUserRepository.On("CountCoinsByUserID", mock.Anything, mock.Anything).Return(constant.CostSendPocket+1, nil).Once()
				s.mockPocketRepository.On("Create", mock.Anything, mock.Anything).Return(nil).Once()
				s.mockUserRepository.On("UpdateCoin", mock.Anything, mock.Anything, mock.Anything).Return(nil).Twice()
				s.mockEventManager.On("Publish", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
			},
			assert: func(err error) {
				s.Nil(err)
			},
		},
		{
			desc:  "success (public)",
			input: &input.PocketInput{ReceiverID: 2, Coins: 1, IsPublic: true},
			on: func() {
				s.mockPocketRepository.On("CountBySenderIdAndReceiverId", mock.Anything, mock.Anything, mock.Anything).Return(0, nil).Once()
				s.mockUserRepository.On("FindByID", mock.Anything, mock.Anything).Return(&domain.User{}, nil).Once()
				s.mockUserRepository.On("CountCoinsByUserID", mock.Anything, mock.Anything).Return(constant.CostSendPocket+1, nil).Once()
				s.mockPocketRepository.On("Create", mock.Anything, mock.Anything).Return(nil).Once()
				s.mockPocketRepository.On("CreateReveal", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
				s.mockUserRepository.On("UpdateCoin", mock.Anything, mock.Anything, mock.Anything).Return(nil).Twice()
				s.mockEventManager.On("Publish", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
			},
			assert: func(err error) {
				s.Nil(err)
			},
		},
		{
			desc:  "user not found",
			input: &input.PocketInput{ReceiverID: 2},
			on: func() {
				s.mockPocketRepository.On("CountBySenderIdAndReceiverId", mock.Anything, mock.Anything, mock.Anything).Return(0, nil).Once()
				s.mockUserRepository.On("FindByID", mock.Anything, mock.Anything).Return(nil, nil).Once()
			},
			assert: func(err error) {
				e, ok := err.(*status.Err)
				if s.True(ok) {
					s.Equal(http.StatusNotFound, e.Code)
				}
			},
		},
		{
			desc:  "not enough coins",
			input: &input.PocketInput{ReceiverID: 2, Coins: 0, IsPublic: false},
			on: func() {
				s.mockPocketRepository.On("CountBySenderIdAndReceiverId", mock.Anything, mock.Anything, mock.Anything).Return(0, nil).Once()
				s.mockUserRepository.On("FindByID", mock.Anything, mock.Anything).Return(&domain.User{}, nil).Once()
				s.mockUserRepository.On("CountCoinsByUserID", mock.Anything, mock.Anything).Return(constant.CostSendPocket-1, nil).Once()
			},
			assert: func(err error) {
				e, ok := err.(*status.Err)
				if s.True(ok) {
					s.Equal(http.StatusForbidden, e.Code)
				}
			},
		},
		{
			desc:  "self sending",
			input: &input.PocketInput{ReceiverID: 1},
			on:    func() {},
			assert: func(err error) {
				e, ok := err.(*status.Err)
				if s.True(ok) {
					s.Equal(http.StatusForbidden, e.Code)
				}
			},
		},
		{
			desc:  "limit send",
			input: &input.PocketInput{ReceiverID: 2},
			on: func() {
				s.mockPocketRepository.On("CountBySenderIdAndReceiverId", mock.Anything, mock.Anything, mock.Anything).Return(constant.SameSendLimit, nil).Once()
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

			err := s.uc.SendPocket(ctx, tc.input)

			tc.assert(err)

			s.mockPocketRepository.AssertExpectations(s.T())
			s.mockUserRepository.AssertExpectations(s.T())
		})
	}
}
