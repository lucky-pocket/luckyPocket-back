package router_test

import (
	"github.com/lucky-pocket/luckyPocket-back/internal/global/error/status"
	"net/http"
	"net/http/httptest"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
	"github.com/stretchr/testify/mock"
)

func (s *UserRouterTestSuite) TestUserDetail() {
	testcases := []struct {
		desc       string
		path       string
		on         func()
		statusCode int
	}{
		{
			desc: "success",
			path: "1",
			on: func() {
				s.mockUserUseCase.On("GetUserDetail", mock.Anything, mock.Anything).Return(&output.UserInfo{}, nil).Once()
			},
			statusCode: http.StatusOK,
		},
		{
			desc: "failed (not found)",
			path: "1",
			on: func() {
				s.mockUserUseCase.On("GetUserDetail", mock.Anything, mock.Anything).Return(nil, status.NewError(http.StatusNotFound, "not found")).Once()
			},
			statusCode: http.StatusNotFound,
		},
	}

	for _, tc := range testcases {
		s.Run(tc.desc, func() {
			tc.on()

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/users/1", nil)
			s.engine.ServeHTTP(w, req)

			s.Equal(tc.statusCode, w.Code, req)
			s.mockUserUseCase.AssertExpectations(s.T())
		})
	}
}
