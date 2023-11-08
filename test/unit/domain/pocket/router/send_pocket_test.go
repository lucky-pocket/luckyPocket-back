package router_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/mock"
)

func (p *PocketRouterTestSuite) TestSendPocket() {
	testcases := []struct {
		desc       string
		body       []byte
		on         func()
		statusCode int
	}{
		{
			desc: "success",
			body: []byte(`
			{
				"receiverID": 1,
  				"coins": 10,
  				"message": "HI?",
	  			"isPublic": true	
            }
			`),
			on: func() {
				p.mockPocketUseCase.On("SendPocket", mock.Anything, mock.Anything).Return(nil).Once()
			},
			statusCode: http.StatusOK,
		},
		{
			desc: "failed (BAD-REQUEST)",
			body: []byte(`
			{
				"receiverID": 1,
  				"coins": 10,
  				"message": "HI?",
	  			"isPublic": fe	
            }
			`),
			on: func() {
			},
			statusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testcases {
		p.Run(tc.desc, func() {
			tc.on()

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/pockets", bytes.NewBuffer(tc.body))
			p.engine.ServeHTTP(w, req)

			p.Equal(tc.statusCode, w.Code, req)
			p.mockPocketUseCase.AssertExpectations(p.T())
		})
	}
}
