package constant

import "time"

const (
	JwtSigningMethod = "HS256"
	// TODO: UNSAFE! change this into proper value on production.
	JwtAccessTTL  = 5 * time.Minute
	JwtRefreshTTL = 10 * 24 * time.Hour
)

const (
	CostRevealSender = 1
	CostSendPocket   = 1
	CostPlayYut      = 2

	PrizeDo     = 1
	PrizeGae    = 1
	PrizeGeol   = 2
	PrizeYut    = 3
	PrizeMo     = 4
	PrizeBackDo = -2
)

const (
	TicketLimit = 1
)
