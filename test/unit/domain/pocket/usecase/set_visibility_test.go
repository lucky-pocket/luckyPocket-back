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

func (s *PocketUseCaseTestSuite) TestSetVisibility() {
	userInfo := auth.Info{
		UserID: 1,
		Role:   constant.RoleMember,
	}

	myPocket := domain.Pocket{
		PocketID: 123,
		Receiver: &domain.User{
			UserID: userInfo.UserID,
		},
	}

	otherPocket := domain.Pocket{
		PocketID: 1531,
		Receiver: &domain.User{
			UserID: userInfo.UserID + 1,
		},
	}

	testcases := []struct {
		desc   string
		input  *input.VisibilityInput
		on     func()
		assert func(err error)
	}{
		{
			desc:  "success",
			input: &input.VisibilityInput{Visible: false},
			on: func() {
				s.mockPocketRepository.On("FindByID", mock.Anything, mock.Anything).Return(&myPocket, nil).Once()
				s.mockPocketRepository.On("UpdateVisibility", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
			},
			assert: func(err error) {
				s.Nil(err)
			},
		},
		{
			desc:  "not owner",
			input: &input.VisibilityInput{Visible: false},
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
	}

	ctx := auth.Inject(context.Background(), userInfo)
	for _, tc := range testcases {
		s.Run(tc.desc, func() {
			tc.on()

			err := s.uc.SetVisibility(ctx, tc.input)

			tc.assert(err)

			s.mockPocketRepository.AssertExpectations(s.T())
			s.mockUserRepository.AssertExpectations(s.T())
		})
	}
}
