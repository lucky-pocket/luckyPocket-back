package usecase_test

import (
	"context"
	"net/http"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/input"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output/mapper"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/error/status"
	"github.com/stretchr/testify/mock"
)

func (s *PocketUseCaseTestSuite) TestGetPocketDetail() {
	receiver := domain.User{
		UserID: 1,
		Name:   "receiver",
	}

	pocket := domain.Pocket{
		PocketID: 123,
		Receiver: &receiver,
		Sender:   &receiver,
	}

	testcases := []struct {
		desc     string
		input    *input.PocketIDInput
		userInfo *auth.Info
		on       func()
		assert   func(output *output.PocketOutput, err error)
	}{
		{
			desc:     "success (without sender)",
			input:    &input.PocketIDInput{},
			userInfo: &auth.Info{UserID: receiver.UserID},
			on: func() {
				s.mockPocketRepository.On("FindByID", mock.Anything, mock.Anything).Return(&pocket, nil).Once()
				s.mockPocketRepository.On("RevealExists", mock.Anything, mock.Anything, mock.Anything).Return(false, nil).Once()
			},
			assert: func(output *output.PocketOutput, err error) {
				if s.Nil(err) {
					s.Equal(mapper.ToPocketOutput(&pocket, nil), output)
				}
			},
		},
		{
			desc:     "success (with sender)",
			input:    &input.PocketIDInput{},
			userInfo: &auth.Info{UserID: receiver.UserID},
			on: func() {
				s.mockPocketRepository.On("FindByID", mock.Anything, mock.Anything).Return(&pocket, nil).Once()
				s.mockPocketRepository.On("RevealExists", mock.Anything, mock.Anything, mock.Anything).Return(true, nil).Once()
				s.mockUserRepository.On("FindByID", mock.Anything, mock.Anything).Return(&receiver, nil).Once()
			},
			assert: func(output *output.PocketOutput, err error) {
				if s.Nil(err) {
					s.Equal(mapper.ToPocketOutput(&pocket, &receiver), output)
				}
			},
		},
		{
			desc:     "pocket not found",
			input:    &input.PocketIDInput{},
			userInfo: &auth.Info{UserID: receiver.UserID},
			on: func() {
				s.mockPocketRepository.On("FindByID", mock.Anything, mock.Anything).Return(nil, nil).Once()
			},
			assert: func(output *output.PocketOutput, err error) {
				e, ok := err.(*status.Err)
				if s.True(ok) {
					s.Equal(http.StatusNotFound, e.Code)
				}
			},
		},
		{
			desc:     "no permission",
			input:    &input.PocketIDInput{},
			userInfo: &auth.Info{UserID: receiver.UserID + 1},
			on: func() {
				s.mockPocketRepository.On("FindByID", mock.Anything, mock.Anything).Return(&pocket, nil).Once()
			},
			assert: func(output *output.PocketOutput, err error) {
				e, ok := err.(*status.Err)
				if s.True(ok) {
					s.Equal(http.StatusForbidden, e.Code)
				}
			},
		},
	}

	for _, tc := range testcases {
		s.Run(tc.desc, func() {
			tc.on()

			ctx := context.Background()
			if tc.userInfo != nil {
				ctx = auth.Inject(ctx, *tc.userInfo)
			}

			output, err := s.uc.GetPocketDetail(ctx, tc.input)

			tc.assert(output, err)

			s.mockPocketRepository.AssertExpectations(s.T())
			s.mockUserRepository.AssertExpectations(s.T())
		})
	}
}
