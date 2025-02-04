-- name: CreateCustomer :one
INSERT INTO customers (
  username,
  hashed_password,
  role
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetCustomerByUsername :one
SELECT * FROM customers
WHERE username = $1;

-- name: ListCustomers :many
SELECT * FROM customers
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteCustomer :exec
DELETE FROM customers
WHERE id = $1;
