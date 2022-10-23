package query

import (
	"context"
	"github.com/rifat-simoom/payment-system/internal/card/domain/card_info"
	"github.com/rifat-simoom/payment-system/internal/common/decorator"
	"github.com/sirupsen/logrus"
)

type GetCardInfo struct {
	CardNo string
}

type GetCardInfoHandler decorator.QueryHandler[GetCardInfo, bool]

type getCardInfoHandler struct {
	cardInfoRepo card_info.Repository
}

func NewGetCardInfoHandler(
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) GetCardInfoHandler {
	if hourRepo == nil {
		panic("nil hourRepo")
	}

	return decorator.ApplyQueryDecorators[GetCardInfo, bool](
		getCardInfoHandler{cardInfoRepo: cardInfoRepo},
		logger,
		metricsClient,
	)
}

func (h getCardInfoHandler) Handle(ctx context.Context, query GetCardInfo) (bool, error) {
	cardInfo, err := h.hourRepo.GetCardInfo(ctx, query.Hour)
	if err != nil {
		return false, err
	}

	return true, nil
}
