package router_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/error/status"
	"github.com/stretchr/testify/mock"
)

func (s *AuthRouterTestSuite) TestRefreshToken() {
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
				s.mockAuthUseCase.On("RefreshToken", mock.Anything, mock.Anything).Return(&output.TokenOutput{}, nil).Once()
			},
			statusCode: http.StatusOK,
		},
		{
			desc:  "failed (BAD-REQUEST)",
			token: "",
			on: func() {
				s.mockAuthUseCase.On("RefreshToken", mock.Anything, mock.Anything).Return(nil, status.NewError(http.StatusUnauthorized, "invalid token")).Once()
			},
			statusCode: http.StatusUnauthorized,
		},
	}

	for _, tc := range testcases {
		s.Run(tc.desc, func() {
			tc.on()

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/auth/refresh", nil)
			req.AddCookie(&http.Cookie{Name: "refreshToken", Value: tc.token})
			s.engine.ServeHTTP(w, req)

			s.Equal(tc.statusCode, w.Code, req)
			s.mockAuthUseCase.AssertExpectations(s.T())
		})
	}
}
