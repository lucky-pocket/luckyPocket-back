package repository

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/ent"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/ent/pocket"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/ent/user"
	"github.com/lucky-pocket/luckyPocket-back/internal/infra/data/ent/mapper"
)

func (r *pocketRepository) FindByID(ctx context.Context, pocketID uint64) (*domain.Pocket, error) {
	pocket, err := r.getClient(ctx).Pocket.Query().
		Where(pocket.ID(pocketID)).
		First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	return mapper.ToPocketDomain(pocket), nil
}

func (r *pocketRepository) FindListByUserID(ctx context.Context, userID uint64, offset, limit int) ([]*domain.Pocket, error) {
	entities, err := r.getClient(ctx).Pocket.Query().
		Where(
			pocket.HasReceiverWith(user.ID(userID)),
		).
		Order(pocket.ByCreatedAt(sql.OrderAsc())).
		Offset(offset).
		Limit(limit).
		All(ctx)

	if err != nil {
		return nil, err
	}

	pockets := make([]*domain.Pocket, 0, len(entities))
	for _, entity := range entities {
		pockets = append(pockets, mapper.ToPocketDomain(entity))
	}
	return pockets, nil
}

func (r *pocketRepository) FillSenderNameOnRevealed(ctx context.Context, pockets []*domain.Pocket, receiverID, userID uint64) error {
	senderPockets, err := r.getClient(ctx).Pocket.Query().
		Select(pocket.FieldID, pocket.SenderColumn).
		Where(
			pocket.And(
				pocket.HasRevealersWith(user.ID(userID)),
				pocket.ReceiverID(receiverID),
			),
		).
		WithSender().
		All(ctx)
	if err != nil {
		return err
	}

	senderMap := make(map[uint64]*ent.User, len(pockets))
	for _, pocket := range senderPockets {
		senderMap[pocket.ID] = pocket.Edges.Sender
	}

	for _, pocket := range pockets {
		if sender, ok := senderMap[pocket.PocketID]; ok {
			pocket.Sender = mapper.ToUserDomain(sender)
		}
	}

	return nil
}
