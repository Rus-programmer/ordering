package metrics

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	db "ordering/db/sqlc"
	"ordering/dto"
)

func (m *metric) GetMetrics(ctx context.Context) (dto.MetricsResponse, error) {
	totalResponse, err := m.store.GetTotalRequests(ctx)
	if err != nil {
		return dto.MetricsResponse{}, err
	}

	requestsByMethod, err := m.store.GetTotalRequestsByMethod(ctx)
	if err != nil {
		return dto.MetricsResponse{}, err
	}

	requestsByPath, err := m.store.GetTotalRequestsByPath(ctx)
	if err != nil {
		return dto.MetricsResponse{}, err
	}

	requestsByStatus, err := m.store.GetTotalRequestsByStatusCode(ctx)
	if err != nil {
		return dto.MetricsResponse{}, err
	}

	errorRate, err := m.store.GetErrorRate(ctx)
	if err != nil {
		return dto.MetricsResponse{}, err
	}
	rate, _ := numericToFloat64(errorRate)

	return dto.MetricsResponse{
		TotalRequests: totalResponse,
		RequestsByMethod: sliceToMapGeneric(
			requestsByMethod,
			func(r db.GetTotalRequestsByMethodRow) string { return r.Method },
			func(r db.GetTotalRequestsByMethodRow) int64 { return r.Count },
		),
		RequestsByPath: sliceToMapGeneric(
			requestsByPath,
			func(r db.GetTotalRequestsByPathRow) string { return r.Path },
			func(r db.GetTotalRequestsByPathRow) int64 { return r.Count },
		),
		RequestsByStatus: sliceToMapGeneric(
			requestsByStatus,
			func(r db.GetTotalRequestsByStatusCodeRow) int32 { return r.StatusCode },
			func(r db.GetTotalRequestsByStatusCodeRow) int64 { return r.Count },
		),
		ErrorRate: rate,
	}, nil
}

func sliceToMapGeneric[T any, K comparable, V any](rows []T, keyFunc func(T) K, valueFunc func(T) V) map[K]V {
	result := make(map[K]V)
	for _, row := range rows {
		result[keyFunc(row)] = valueFunc(row)
	}
	return result
}

func numericToFloat64(num pgtype.Numeric) (float64, error) {
	if !num.Valid {
		return 0, fmt.Errorf("invalid numeric value")
	}

	f, _ := num.Float64Value()
	return f.Float64, nil
}
