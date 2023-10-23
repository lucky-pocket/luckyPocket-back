package constant

import "time"

const (
	JwtSigningMethod = "HS256"
	JwtAccessTTL     = 5 * time.Minute
	JwtRefreshTTL    = 10 * 24 * time.Hour
)
