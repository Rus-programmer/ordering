-- name: CreateOrder :one
INSERT INTO orders (customer_id, total_price)
VALUES ($1, $2)
RETURNING *;

-- name: ListOrders :many
SELECT id
FROM orders
WHERE (customer_id = $1 OR EXISTS (SELECT 1 FROM customers WHERE id = $1 AND role = 'admin'))
  AND (sqlc.narg(status)::order_status IS NULL OR status = sqlc.narg(status))
  AND (sqlc.narg(min_price)::bigint IS NULL OR total_price >= sqlc.narg(min_price))
  AND (sqlc.narg(max_price)::bigint IS NULL OR total_price <= sqlc.narg(max_price))
  AND is_deleted = false;

-- name: GetOrder :one
SELECT *
FROM orders o
WHERE o.id = $1
  and (customer_id = $2 OR EXISTS (SELECT 1 FROM customers WHERE id = $2 AND role = 'admin'))
  AND is_deleted = false;

-- name: UpdateOrder :one
UPDATE orders o
SET
    status = COALESCE(sqlc.narg(status), status),
    total_price = COALESCE(sqlc.narg(total_price), total_price)
WHERE o.id = $1
  AND (customer_id = $2 OR EXISTS (SELECT 1 FROM customers WHERE id = $2 AND role = 'admin'))
RETURNING *;

-- name: SoftDeleteOrder :one
UPDATE orders o
SET is_deleted = TRUE
WHERE o.id = $1
  AND (customer_id = $2 OR EXISTS (SELECT 1 FROM customers WHERE id = $2 AND role = 'admin'))
RETURNING *;

-- name: DeleteOrder :exec
DELETE
FROM orders
WHERE id = $1;
