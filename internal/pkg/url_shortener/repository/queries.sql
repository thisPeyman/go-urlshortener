-- name: CreateShortURL :exec
INSERT INTO urls (short_url, long_url) VALUES ($1, $2);

-- name: GetLongURL :one
SELECT long_url FROM urls WHERE short_url = $1;

-- name: DeleteShortURL :exec
DELETE FROM urls WHERE short_url = $1;