package usecase_test

import (
	"context"
	"net/http"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output/mapper"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/error/status"
	"github.com/stretchr/testify/mock"
)

func (s *UserUseCaseTestSuite) TestGetMyDetail() {
	userInfo := auth.Info{
		UserID: 1,
		Role:   constant.RoleMember,
	}

	user := domain.User{
		UserID: userInfo.UserID,
		Name:   "hi",
		Role:   userInfo.Role,
	}

	testcases := []struct {
		desc   string
		on     func()
		assert func(output *output.MyDetailOutput, err error)
	}{
		{
			desc: "success",
			on: func() {
				s.mockUserRepository.On("FindByID", mock.Anything, mock.Anything).Return(&user, nil).Once()
				s.mockNoticeRepository.On("ExistsByUserID", mock.Anything, mock.Anything).Return(true, nil).Once()
			},
			assert: func(output *output.MyDetailOutput, err error) {
				if s.Nil(err) {
					s.Equal(mapper.ToMyDetailOutput(user, true), output)
				}
			},
		},
		{
			desc: "user not found",
			on: func() {
				s.mockUserRepository.On("FindByID", mock.Anything, mock.Anything).Return(nil, nil).Once()
			},
			assert: func(output *output.MyDetailOutput, err error) {
				e, ok := err.(*status.Err)
				if s.True(ok) {
					s.Equal(http.StatusNotFound, e.Code)
				}
			},
		},
	}

	ctx := auth.Inject(context.Background(), userInfo)
	for _, tc := range testcases {
		s.Run(tc.desc, func() {
			tc.on()

			output, err := s.uc.GetMyDetail(ctx)

			tc.assert(output, err)

			s.mockUserRepository.AssertExpectations(s.T())
			s.mockNoticeRepository.AssertExpectations(s.T())
		})
	}
}
