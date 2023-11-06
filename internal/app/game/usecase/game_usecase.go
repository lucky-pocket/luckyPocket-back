package usecase

import (
	"context"
	"math/rand"
	"net/http"
	"time"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/input"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output"
	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/output/mapper"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/auth"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/error/status"
	"github.com/lucky-pocket/luckyPocket-back/internal/global/tx"
	"github.com/pkg/errors"
)

type Deps struct {
	UserRepository    domain.UserRepository
	TicketRepository  domain.TicketRepository
	GameLogRepository domain.GameLogRepository
	TxManager         tx.Manager
}

type gameUseCase struct {
	*Deps
}

func NewGameUseCase(deps *Deps) domain.GameUseCase {
	return &gameUseCase{deps}
}

func (g *gameUseCase) PlayYut(ctx context.Context, input *input.FreeInput) (*output.YutOutput, error) {
	user := auth.MustExtract(ctx)

	var output *output.YutOutput
	err := g.TxManager.WithTx(ctx, func(ctx context.Context) error {
		if input.Free {
			count, err := g.TicketRepository.CountByUserID(ctx, user.UserID)
			if err != nil {
				return errors.Wrap(err, "unexpected error")
			}

			if count >= constant.TicketLimit {
				return status.NewError(http.StatusForbidden, "ticket limit exceeded")
			}

			err = g.TicketRepository.Increase(ctx, user.UserID)
			if err != nil {
				return errors.Wrap(err, "unexpected error")
			}
		} else {
			coins, err := g.UserRepository.CountCoinsByUserID(ctx, user.UserID)
			if err != nil {
				return errors.Wrap(err, "unexpected error")
			}

			if coins < constant.CostPlayYut {
				return status.NewError(http.StatusForbidden, "insufficient coins")
			}

			err = g.UserRepository.UpdateCoin(ctx, user.UserID, coins-constant.CostPlayYut)
			if err != nil {
				return errors.Wrap(err, "unexpected error")
			}
		}

		marked := rand.Intn(2) == 1
		yutPieces := [3]bool{
			rand.Intn(2) == 1,
			rand.Intn(2) == 1,
			rand.Intn(2) == 1,
		}

		coinsEarned := g.evaluateYutResult(marked, yutPieces)

		coins, err := g.UserRepository.CountCoinsByUserID(ctx, user.UserID)
		if err != nil {
			return errors.Wrap(err, "unexpected error")
		}

		err = g.UserRepository.UpdateCoin(ctx, user.UserID, coins+coinsEarned)
		if err != nil {
			return errors.Wrap(err, "unexpected error")
		}

		err = g.GameLogRepository.Create(ctx, domain.GameLog{
			User:      &domain.User{UserID: user.UserID},
			TimeStamp: time.Now(),
			GameType:  "Yut",
		})
		if err != nil {
			return errors.Wrap(err, "unexpected error")
		}

		output = mapper.ToYutOutput(marked, yutPieces, coinsEarned)
		return nil
	},
		g.TicketRepository.(tx.Transactor).NewTx(),
		g.UserRepository.(tx.Transactor).NewTx(),
		g.GameLogRepository.(tx.Transactor).NewTx(),
	)

	return output, err
}

func (g *gameUseCase) evaluateYutResult(marked bool, yutPieces [3]bool) (coinsEarned int) {
	yutCount := 0
	for _, value := range yutPieces {
		if value {
			yutCount++
		}
	}

	if marked && yutCount == 0 {
		coinsEarned = constant.PrizeBackDo
	} else {
		if marked {
			yutCount++
		}

		switch yutCount {
		case 1:
			coinsEarned = constant.PrizeDo
		case 2:
			coinsEarned = constant.PrizeGae
		case 3:
			coinsEarned = constant.PrizeGeol
		case 4:
			coinsEarned = constant.PrizeMo
		default:
			coinsEarned = constant.PrizeYut
		}
	}

	return coinsEarned
}

func (g *gameUseCase) GetTicketInfo(ctx context.Context) (*output.TicketOutput, error) {
	user := auth.MustExtract(ctx)

	count, err := g.TicketRepository.CountByUserID(ctx, user.UserID)
	if err != nil {
		return nil, errors.Wrap(err, "unexpected error")
	}

	refillAt, err := g.TicketRepository.GetRefillAt(ctx, user.UserID)
	if err != nil {
		return nil, errors.Wrap(err, "unexpected error")
	}

	ticket := constant.TicketLimit - count

	return mapper.ToFreeTicketOutput(ticket, refillAt), nil
}
