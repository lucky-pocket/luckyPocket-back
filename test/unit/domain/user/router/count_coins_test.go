package router_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
	"github.com/stretchr/testify/mock"
)

func (s *UserRouterTestSuite) TestCountCoins() {
	testcases := []struct {
		desc       string
		on         func()
		statusCode int
	}{
		{
			desc: "success",
			on: func() {
				s.mockUserUseCase.On("CountCoins", mock.Anything).Return(&output.CoinOutput{}, nil).Once()
			},
			statusCode: http.StatusOK,
		},
	}

	for _, tc := range testcases {
		s.Run(tc.desc, func() {
			tc.on()

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/users/me/coins", nil)
			s.engine.ServeHTTP(w, req)

			s.Equal(tc.statusCode, w.Code, req)
			s.mockUserUseCase.AssertExpectations(s.T())
		})
	}
}
