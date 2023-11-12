package usecase_test

import (
	"context"
	"net/http"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/input"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output/mapper"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/error/status"
	"github.com/stretchr/testify/mock"
)

func (s *UserUseCaseTestSuite) TestGetUserDetail() {
	user := domain.User{
		UserID: 1,
		Name:   "hi",
		Role:   constant.RoleMember,
	}

	testcases := []struct {
		desc   string
		input  *input.UserIDInput
		on     func()
		assert func(output *output.UserInfo, err error)
	}{
		{
			desc:  "success",
			input: &input.UserIDInput{},
			on: func() {
				s.mockUserRepository.On("FindByID", mock.Anything, mock.Anything).Return(&user, nil).Once()
			},
			assert: func(output *output.UserInfo, err error) {
				if s.Nil(err) {
					s.Equal(mapper.ToUserInfo(user), output)
				}
			},
		},
		{
			desc:  "user not found",
			input: &input.UserIDInput{},
			on: func() {
				s.mockUserRepository.On("FindByID", mock.Anything, mock.Anything).Return(nil, nil).Once()
			},
			assert: func(output *output.UserInfo, err error) {
				e, ok := err.(*status.Err)
				if s.True(ok) {
					s.Equal(http.StatusNotFound, e.Code)
				}
			},
		},
	}

	for _, tc := range testcases {
		s.Run(tc.desc, func() {
			tc.on()

			output, err := s.uc.GetUserDetail(context.Background(), tc.input)

			tc.assert(output, err)

			s.mockUserRepository.AssertExpectations(s.T())
			s.mockNoticeRepository.AssertExpectations(s.T())
		})
	}
}
