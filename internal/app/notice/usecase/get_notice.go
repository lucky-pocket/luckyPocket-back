package usecase

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output/mapper"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth"
	"github.com/pkg/errors"
)

func (n *noticeUseCase) GetNotice(ctx context.Context) (*output.NoticeListOutput, error) {
	userInfo := auth.MustExtract(ctx)

	notices, err := n.NoticeRepository.FindAllByUserID(ctx, userInfo.UserID)
	if err != nil {
		return nil, errors.Wrap(err, "unexpected db error")
	}

	return mapper.ToNoticeListOutPut(notices), nil
}
