package usecase

import (
	"context"
	"net/http"
	"time"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/input"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth/jwt"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/error/status"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
)

func (l *AuthUseCaseTestSuite) TestRefresh() {
	userInfo := auth.Info{
		UserID: 1,
		Role:   constant.RoleMember,
	}

	testcases := []struct {
		desc   string
		input  *input.RefreshInput
		on     func()
		assert func(output *output.TokenOutput, err error)
	}{
		{
			desc:  "success",
			input: &input.RefreshInput{RefreshToken: "secret"},
			on: func() {
				l.mockJwtParser.On("Parse", mock.Anything).Return(&jwt.Token{}, nil).Once()
				l.mockBlackListRepository.On("Exists", mock.Anything, mock.Anything).Return(false, nil).Once()
				l.mockBlackListRepository.On("Save", mock.Anything, mock.Anything).Return(nil).Once()
				l.mockJwtIssuer.On("IssueAccess", mock.Anything).Return("AccessToken", time.Time{}).Once()
				l.mockJwtIssuer.On("IssueRefresh", mock.Anything).Return("RefreshToken", time.Time{}).Once()
			},
			assert: func(output *output.TokenOutput, err error) {
				if l.Nil(err) {
					l.Equal(output.Access.Token, "AccessToken")
					l.Equal(output.Refresh.Token, "RefreshToken")
				}
			},
		},

		{
			desc:  "token is invalid",
			input: &input.RefreshInput{RefreshToken: "secret"},
			on: func() {
				l.mockJwtParser.On("Parse", mock.Anything).Return(nil, errors.New("token is not valid")).Once()
			},
			assert: func(output *output.TokenOutput, err error) {
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
				l.mockJwtParser.On("Parse", mock.Anything).Return(&jwt.Token{}, nil).Once()
				l.mockBlackListRepository.On("Exists", mock.Anything, mock.Anything).Return(true, nil).Once()
			},
			assert: func(output *output.TokenOutput, err error) {
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

			token, err := l.uc.RefreshToken(ctx, tc.input)

			tc.assert(token, err)

			l.mockBlackListRepository.AssertExpectations(l.T())
			l.mockUserRepository.AssertExpectations(l.T())
		})
	}
}
