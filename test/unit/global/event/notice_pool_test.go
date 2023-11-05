package event_test

import (
	"context"
	"testing"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/event/dispatcher"
	"github.com/lucky-pocket/luckyPocket-back/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNoticePoolDumper(t *testing.T) {
	mockNoticePool := mocks.NewNoticePool(t)

	d := dispatcher.NewNoticePoolDumper(&dispatcher.NoticePoolDumperDeps{
		NoticePool: mockNoticePool,
	})

	mockNoticePool.On("Put", mock.Anything, mock.Anything).Return(nil)

	err := d.Dispatch(context.Background(), "", &domain.Notice{})
	if !assert.NoError(t, err) {
		return
	}

	err = d.Dispatch(context.Background(), "", nil)
	if !assert.ErrorContains(t, err, "assertion") {
		return
	}
}
