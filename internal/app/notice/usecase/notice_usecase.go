package usecase

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output/mapper"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth"
	"github.com/pkg/errors"
)

type Deps struct {
	NoticeRepository domain.NoticeRepository
	UserRepository   domain.UserRepository
}

type noticeUseCase struct{ *Deps }

func NewNoticeUseCase(deps *Deps) domain.NoticeUseCase {
	return &noticeUseCase{deps}
}

func (n *noticeUseCase) GetNotice(ctx context.Context) (*output.NoticeListOutput, error) {
	userInfo := auth.MustExtract(ctx)

	user, err := n.UserRepository.FindByID(ctx, userInfo.UserID)

	if err != nil {
		return nil, errors.Wrap(err, "unexpected db error")
	}

	if user == nil {
		return nil, errors.Wrap(err, "user not found")
	}

	notices, err := n.NoticeRepository.FindAllByUserID(ctx, user.UserID)

	if err != nil {
		return nil, errors.Wrap(err, "unexpected db error")
	}

	return mapper.ToNoticeListOutPut(*notices), nil
}
