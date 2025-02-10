package metrics

import (
	"context"
	db "ordering/db/sqlc"
	"ordering/dto"
)

type Metric interface {
	GetMetrics(ctx context.Context) (dto.MetricsResponse, error)
}

type metric struct {
	store db.Store
}

func NewMetric(store db.Store) Metric {
	return &metric{
		store: store,
	}
}
