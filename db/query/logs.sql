-- name: CreateLog :exec
INSERT INTO logs (
    method,
    path,
    status_code,
    elapsed_time,
    time
) VALUES ($1, $2, $3, $4, $5);