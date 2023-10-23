package auth

import (
	"context"
	"errors"

	"github.com/lucky-pocket/luckyPocket-back/internal/domain/data/constant"
)

// key is a key of context's value map.
type key struct{}

// Info is a payload part of user's jwt token.
type Info struct {
	UserID uint64
	Role   constant.Role
}

// InjectInfo injects user info into context.
func Inject(ctx context.Context, info Info) context.Context {
	return context.WithValue(ctx, key{}, &info)
}

// ExtractInfo extracts user info from context. If not found, it returns error.
func Extract(ctx context.Context) (*Info, error) {
	info, ok := ctx.Value(key{}).(*Info)
	if !ok {
		return nil, errors.New("auth info not found")
	}
	return info, nil
}

// MustExtract does the same thing as Extract. But it panics if error != nil
func MustExtract(ctx context.Context) *Info {
	info, err := Extract(ctx)
	if err != nil {
		panic(err)
	}
	return info
}
