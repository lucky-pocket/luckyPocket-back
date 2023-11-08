package router_test

import (
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
	"github.com/stretchr/testify/mock"
)

func (p *PocketRouterTestSuite) TestGetMyPocketDetail() {
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
				p.mockPocketUseCase.On("GetPocketDetail", mock.Anything, mock.Anything).Return(&output.PocketOutput{}, nil).Once()
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
			req, _ := http.NewRequest("GET", "/pockets/"+tc.param, nil)
			p.engine.ServeHTTP(w, req)

			log.Print("Response" + w.Body.String())

			p.Equal(tc.statusCode, w.Code, req)
			p.mockPocketUseCase.AssertExpectations(p.T())
		})
	}
}
