package router_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
	"github.com/stretchr/testify/mock"
)

func (p *PocketRouterTestSuite) TestRevealSender() {
	senderInfo := output.UserInfo{
		UserID: 1,
		Name:   "mock",
	}

	testcases := []struct {
		desc       string
		param      string
		on         func()
		statusCode int
	}{
		{
			desc:  "success",
			param: "1",
			on: func() {
				p.mockPocketUseCase.On("RevealSender", mock.Anything, mock.Anything).Return(&senderInfo, nil).Once()
			},
			statusCode: http.StatusOK,
		},
		{
			desc:  "failed (bad-request)",
			param: "11111111111111111111111111111111111111111111111111111111111111111111",
			on: func() {
			},
			statusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testcases {
		p.Run(tc.desc, func() {
			tc.on()

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/users/me/pockets/"+tc.param+"/sender", nil)
			p.engine.ServeHTTP(w, req)

			p.Equal(tc.statusCode, w.Code, req)
			p.mockPocketUseCase.AssertExpectations(p.T())
		})
	}
}
