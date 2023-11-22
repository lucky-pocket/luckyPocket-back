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

	var (
		marked bool
		out    *output.YutOutput
	)

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

		isNak := rand.Intn(20) == 0
		if isNak {
			out = &output.YutOutput{Output: "낙"}
			return nil
		}

		yutPieces := [3]bool{
			rand.Intn(2) == 1,
			rand.Intn(2) == 1,
			rand.Intn(2) == 1,
		}

		marked = rand.Intn(2) == 1

		coinsEarned, result := g.evaluateYutResult(marked, yutPieces)

		out = mapper.ToYutOutput(marked, yutPieces, coinsEarned, result)

		coins, err := g.UserRepository.CountCoinsByUserID(ctx, user.UserID)
		if err != nil {
			return errors.Wrap(err, "unexpected error")
		}

		err = g.UserRepository.UpdateCoin(ctx, user.UserID, max(coins+coinsEarned, 0))
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

		return nil
	},
		g.TicketRepository.(tx.Transactor).NewTx(),
		g.UserRepository.(tx.Transactor).NewTx(),
		g.GameLogRepository.(tx.Transactor).NewTx(),
	)

	return out, err
}

func (g *gameUseCase) evaluateYutResult(marked bool, yutPieces [3]bool) (coinsEarned int, result string) {
	var yutCount int
	for _, value := range yutPieces {
		if value {
			yutCount++
		}
	}

	if marked && yutCount == 0 {
		coinsEarned = constant.PrizeBackDo
		result = "빽도"
	} else {
		if marked {
			yutCount++
		}

		switch yutCount {
		case 1:
			coinsEarned = constant.PrizeDo
			result = "도"
		case 2:
			coinsEarned = constant.PrizeGae
			result = "개"
		case 3:
			coinsEarned = constant.PrizeGeol
			result = "걸"
		case 4:
			coinsEarned = constant.PrizeMo
			result = "모"
		default:
			coinsEarned = constant.PrizeYut
			result = "윷"
		}
	}
	return
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
