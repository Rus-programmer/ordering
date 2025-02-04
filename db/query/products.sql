-- name: CreateProduct :one
INSERT INTO products (
  name,
  price,
  quantity
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetProduct :one
SELECT * FROM products
WHERE id = $1;

-- name: ListProducts :many
SELECT * FROM products
ORDER BY name
LIMIT $1
OFFSET $2;

-- name: UpdateProduct :one
UPDATE products
SET
    name = COALESCE(sqlc.narg(name), name),
    quantity = COALESCE(sqlc.narg(quantity), quantity),
    price = COALESCE(sqlc.narg(price), price)
WHERE
    id = sqlc.arg(id)
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1;