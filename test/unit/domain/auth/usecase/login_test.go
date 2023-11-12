package usecase

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/input"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/error/status"
	"github.com/stretchr/testify/mock"
)

func (l *AuthUseCaseTestSuite) TestLogin() {
	userInfo := auth.Info{
		UserID: 1,
		Role:   constant.RoleMember,
	}

	testcases := []struct {
		desc   string
		input  *input.CodeInput
		on     func()
		assert func(output *output.TokenOutput, err error)
	}{
		{
			desc:  "success",
			input: &input.CodeInput{Code: "secret"},
			on: func() {
				l.mockGAuthClient.On("IssueToken", mock.Anything).Return("", "", nil).Once()
				l.mockGAuthClient.On("GetUserInfo", mock.Anything).Return(&domain.GAuthUser{}, nil).Once()
				l.mockUserRepository.On("ExistsByEmail", mock.Anything, mock.Anything).Return(false, nil).Once()
				l.mockUserRepository.On("Create", mock.Anything, mock.Anything).Return(&domain.User{UserID: 1, Role: constant.RoleMember}, nil).Once()
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
			desc:  "success (exists user)",
			input: &input.CodeInput{Code: "secret"},
			on: func() {
				l.mockGAuthClient.On("IssueToken", mock.Anything).Return("", "", nil).Once()
				l.mockGAuthClient.On("GetUserInfo", mock.Anything).Return(&domain.GAuthUser{}, nil).Once()
				l.mockUserRepository.On("ExistsByEmail", mock.Anything, mock.Anything).Return(true, nil).Once()
				l.mockUserRepository.On("FindByEmail", mock.Anything, mock.Anything).Return(&domain.User{}, nil).Once()
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
			desc:  "bad request",
			input: &input.CodeInput{},
			on: func() {
				l.mockGAuthClient.On("IssueToken", mock.Anything).Return("", "", errors.New("bad request")).Once()
			},
			assert: func(output *output.TokenOutput, err error) {
				e, ok := err.(*status.Err)
				if l.True(ok) {
					l.Equal(http.StatusBadRequest, e.Code)
				}
			},
		},
	}

	ctx := auth.Inject(context.Background(), userInfo)
	for _, tc := range testcases {
		l.Run(tc.desc, func() {
			tc.on()

			token, err := l.uc.Login(ctx, tc.input)

			tc.assert(token, err)

			l.mockBlackListRepository.AssertExpectations(l.T())
			l.mockUserRepository.AssertExpectations(l.T())
		})
	}
}
