package repository

import (
	"context"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/tx"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/ent"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/ent/notice"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/ent/user"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/mapper"
	ent_tx "github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/tx"
)

type noticeRepository struct{ client *ent.Client }

func NewNoticeRepository(client *ent.Client) domain.NoticeRepository {
	return &noticeRepository{client}
}

func (r *noticeRepository) NewTx() tx.Tx { return ent_tx.New(r.client) }

func (r *noticeRepository) getClient(ctx context.Context) *ent.Client {
	tx, err := ent_tx.FromContext(ctx)
	if err != nil {
		return r.client
	}
	return tx.Client()
}

func (r *noticeRepository) CreateBulk(ctx context.Context, notices []*domain.Notice) error {
	builders := make([]*ent.NoticeCreate, 0, len(notices))

	for _, notice := range notices {
		builders = append(builders,
			r.client.Notice.Create().
				SetPocketID(notice.PocketID).
				SetUserID(notice.UserID).
				SetType(notice.Type).
				SetChecked(false),
		)
	}

	return r.getClient(ctx).Notice.CreateBulk(builders...).Exec(ctx)
}

func (r *noticeRepository) FindAllByUserID(ctx context.Context, userID uint64) ([]*domain.Notice, error) {
	entities, err := r.getClient(ctx).Notice.Query().
		Where(notice.HasUserWith(user.ID(userID))).
		All(ctx)

	if err != nil {
		return nil, err
	}

	notices := make([]*domain.Notice, 0, len(entities))
	for _, entity := range entities {
		notices = append(notices, mapper.ToNoticeDomain(entity))
	}

	return notices, nil
}

func (r *noticeRepository) FindByID(ctx context.Context, noticeID uint64) (*domain.Notice, error) {
	notice, err := r.getClient(ctx).Notice.Get(ctx, noticeID)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	return mapper.ToNoticeDomain(notice), err
}

func (r *noticeRepository) ExistsByUserID(ctx context.Context, userID uint64) (bool, error) {
	return r.getClient(ctx).Notice.Query().
		Where(notice.HasUserWith(user.ID(userID))).
		Exist(ctx)
}
