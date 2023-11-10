package router_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/lucky-pocket/luckyPocket-back/internal/global/error/status"
	"github.com/stretchr/testify/mock"
)

func (s *AuthRouterTestSuite) TestLogout() {
	testcases := []struct {
		desc       string
		token      string
		on         func()
		statusCode int
	}{
		{
			desc:  "success",
			token: "refreshToken",
			on: func() {
				s.mockAuthUseCase.On("Logout", mock.Anything, mock.Anything).Return(nil).Once()
			},
			statusCode: http.StatusOK,
		},
		{
			desc:  "failed (BAD-REQUEST)",
			token: "",
			on: func() {
				s.mockAuthUseCase.On("Logout", mock.Anything, mock.Anything).Return(status.NewError(http.StatusUnauthorized, "invalid token")).Once()
			},
			statusCode: http.StatusUnauthorized,
		},
	}

	for _, tc := range testcases {
		s.Run(tc.desc, func() {
			tc.on()

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/auth/logout", nil)
			req.AddCookie(&http.Cookie{Name: "refreshToken", Value: tc.token})
			s.engine.ServeHTTP(w, req)

			s.Equal(tc.statusCode, w.Code, req)
			s.mockAuthUseCase.AssertExpectations(s.T())
		})
	}
}
