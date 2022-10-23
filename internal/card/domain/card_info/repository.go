package card_info

import (
	"context"
	"time"
)

type Repository interface {
	GetCardInfo(ctx context.Context, hourTime time.Time) (*CardInfo, error)
}
