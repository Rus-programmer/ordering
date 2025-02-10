-- name: GetTotalRequests :one
SELECT COUNT(*) AS total_requests
FROM logs;

-- name: GetTotalRequestsByMethod :many
SELECT method, COUNT(*) AS count
FROM logs
GROUP BY method;

-- name: GetTotalRequestsByStatusCode :many
SELECT status_code, COUNT(*) AS count
FROM logs
GROUP BY status_code;

-- name: GetTotalRequestsByPath :many
SELECT path, COUNT(*) AS count
FROM logs
GROUP BY path
ORDER BY count DESC;

-- name: GetErrorRate :one
SELECT ROUND(100.0 * SUM(CASE WHEN status_code >= 400 THEN 1 ELSE 0 END) / COUNT(*), 2) AS error_rate
FROM logs;


