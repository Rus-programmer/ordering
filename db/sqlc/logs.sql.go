// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: logs.sql

package db

import (
	"context"
	"time"
)

const createLog = `-- name: CreateLog :exec
INSERT INTO logs (
    method,
    path,
    status_code,
    elapsed_time,
    time
) VALUES ($1, $2, $3, $4, $5)
`

type CreateLogParams struct {
	Method      string    `json:"method"`
	Path        string    `json:"path"`
	StatusCode  int32     `json:"status_code"`
	ElapsedTime string    `json:"elapsed_time"`
	Time        time.Time `json:"time"`
}

func (q *Queries) CreateLog(ctx context.Context, arg CreateLogParams) error {
	_, err := q.db.Exec(ctx, createLog,
		arg.Method,
		arg.Path,
		arg.StatusCode,
		arg.ElapsedTime,
		arg.Time,
	)
	return err
}
