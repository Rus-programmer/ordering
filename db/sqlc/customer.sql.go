// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: customer.sql

package db

import (
	"context"
)

const createCustomer = `-- name: CreateCustomer :one
INSERT INTO customers (
  username,
  hashed_password,
  role
) VALUES (
  $1, $2, $3
) RETURNING id, username, role, hashed_password, created_at, updated_at
`

type CreateCustomerParams struct {
	Username       string   `json:"username"`
	HashedPassword string   `json:"hashed_password"`
	Role           UserRole `json:"role"`
}

func (q *Queries) CreateCustomer(ctx context.Context, arg CreateCustomerParams) (Customer, error) {
	row := q.db.QueryRow(ctx, createCustomer, arg.Username, arg.HashedPassword, arg.Role)
	var i Customer
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Role,
		&i.HashedPassword,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteCustomer = `-- name: DeleteCustomer :exec
DELETE FROM customers
WHERE id = $1
`

func (q *Queries) DeleteCustomer(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteCustomer, id)
	return err
}

const getCustomerByUsername = `-- name: GetCustomerByUsername :one
SELECT id, username, role, hashed_password, created_at, updated_at FROM customers
WHERE username = $1
`

func (q *Queries) GetCustomerByUsername(ctx context.Context, username string) (Customer, error) {
	row := q.db.QueryRow(ctx, getCustomerByUsername, username)
	var i Customer
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Role,
		&i.HashedPassword,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listCustomers = `-- name: ListCustomers :many
SELECT id, username, role, hashed_password, created_at, updated_at FROM customers
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListCustomersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListCustomers(ctx context.Context, arg ListCustomersParams) ([]Customer, error) {
	rows, err := q.db.Query(ctx, listCustomers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Customer{}
	for rows.Next() {
		var i Customer
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Role,
			&i.HashedPassword,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
