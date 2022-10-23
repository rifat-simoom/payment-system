package app

import (
	"github.com/rifat-simoom/payment-system/internal/card/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
}

type Queries struct {
	GetCardInfo query.GetCardInfo
}
