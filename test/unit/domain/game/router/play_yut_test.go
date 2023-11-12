package router_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
	"github.com/stretchr/testify/mock"
)

func (s *GameRouterTestSuite) TestPlayYut() {
	testcases := []struct {
		desc       string
		body       []byte
		on         func()
		statusCode int
	}{
		{
			desc: "success",
			body: []byte(`{"free": false}`),
			on: func() {
				s.mockGameUseCase.On("PlayYut", mock.Anything, mock.Anything).Return(&output.YutOutput{}, nil).Once()
			},
			statusCode: http.StatusOK,
		},
		{
			desc: "invalid json",
			body: []byte(`{"free": "hi?"}`),
			on: func() {
			},
			statusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testcases {
		s.Run(tc.desc, func() {
			tc.on()

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/games/yut", bytes.NewBuffer(tc.body))
			s.engine.ServeHTTP(w, req)

			s.Equal(tc.statusCode, w.Code, req)

			s.mockGameUseCase.AssertExpectations(s.T())
		})
	}
}
