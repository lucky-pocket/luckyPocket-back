package tx_test

import (
	"context"
	"errors"
	"testing"

	"github.com/lucky-pocket/luckyPocket-back/internal/global/tx"
	"github.com/lucky-pocket/luckyPocket-back/test/mocks"
	"github.com/stretchr/testify/assert"
)

func TestWithTx(t *testing.T) {
	mockTx := mocks.NewTx(t)

	txManager := tx.NewTxManager()

	testcases := []struct {
		desc   string
		err    error
		on     func()
		assert func(t *testing.T, err error)
	}{
		{
			desc: "success",
			err:  nil,
			on: func() {
				mockTx.On("Transactor").Return("hi").Once()
				mockTx.On("Begin").Return("hi", nil).Once()
				mockTx.On("Commit").Return(nil).Once()
			},
			assert: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
		{
			desc: "rollback succes",
			err:  errors.New("haha you should rollback"),
			on: func() {
				mockTx.On("Transactor").Return("hi").Once()
				mockTx.On("Begin").Return("hi", nil).Once()
				mockTx.On("Rollback").Return(nil).Once()
			},
			assert: func(t *testing.T, err error) {
				assert.ErrorContains(t, err, "haha")
			},
		},
		{
			desc: "begin failed",
			err:  nil,
			on: func() {
				mockTx.On("Transactor").Return("hi").Once()
				mockTx.On("Begin").Return("hi", errors.New("haha new error")).Once()
			},
			assert: func(t *testing.T, err error) {
				assert.ErrorContains(t, err, "new")
			},
		},
		{
			desc: "commit failed",
			err:  nil,
			on: func() {
				mockTx.On("Transactor").Return("hi").Once()
				mockTx.On("Begin").Return("hi", nil).Once()
				mockTx.On("Commit").Return(errors.New("haha commit failed")).Once()
			},
			assert: func(t *testing.T, err error) {
				assert.ErrorContains(t, err, "commit")
			},
		},
		{
			desc: "rollback failed",
			err:  errors.New("haha you should rollback"),
			on: func() {
				mockTx.On("Transactor").Return("hi").Once()
				mockTx.On("Begin").Return("hi", nil).Once()
				mockTx.On("Rollback").Return(nil).Once()
			},
			assert: func(t *testing.T, err error) {
				assert.ErrorContains(t, err, "rollback")
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			tc.on()

			err := txManager.WithTx(context.Background(),
				func(ctx context.Context) error {
					return tc.err
				}, mockTx,
			)

			tc.assert(t, err)

			mockTx.AssertExpectations(t)
		})
	}
}
