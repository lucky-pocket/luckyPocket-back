package usecase

import (
	"context"
	"net/http"

	"github.com/pkg/errors"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/input"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/error/status"
	"github.com/stretchr/testify/mock"
)

func (l *AuthUseCaseTestSuite) TestLogout() {
	userInfo := auth.Info{
		UserID: 1,
		Role:   constant.RoleMember,
	}

	testcases := []struct {
		desc   string
		input  *input.RefreshInput
		on     func()
		assert func(err error)
	}{
		{
			desc:  "success",
			input: &input.RefreshInput{RefreshToken: "secret"},
			on: func() {
				l.mockBlackListRepository.On("Exists", mock.Anything, mock.Anything).Return(false, nil).Once()
				l.mockJwtParser.On("Parse", mock.Anything).Return(generateTestToken(), nil).Once()
				l.mockBlackListRepository.On("Save", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
			},
			assert: func(err error) {
				l.Nil(err)
			},
		},

		{
			desc:  "token is invalid",
			input: &input.RefreshInput{RefreshToken: "secret"},
			on: func() {
				l.mockJwtParser.On("Parse", mock.Anything).Return(nil, errors.New("token is invalid")).Once()
			},
			assert: func(err error) {
				e, ok := err.(*status.Err)
				if l.True(ok) {
					l.Equal(http.StatusUnauthorized, e.Code)
				}
			},
		},

		{
			desc:  "token expired",
			input: &input.RefreshInput{RefreshToken: "secret"},
			on: func() {
				l.mockJwtParser.On("Parse", mock.Anything).Return(generateTestToken(), nil).Once()
				l.mockBlackListRepository.On("Exists", mock.Anything, mock.Anything).Return(true, nil).Once()
			},
			assert: func(err error) {
				e, ok := err.(*status.Err)
				if l.True(ok) {
					l.Equal(http.StatusUnauthorized, e.Code)
				}
			},
		},
	}

	ctx := auth.Inject(context.Background(), userInfo)
	for _, tc := range testcases {
		l.Run(tc.desc, func() {
			tc.on()

			err := l.uc.Logout(ctx, tc.input)

			tc.assert(err)

			l.mockBlackListRepository.AssertExpectations(l.T())
			l.mockUserRepository.AssertExpectations(l.T())
		})
	}
}
