-- name: CreateUsers :one
INSERT INTO users (
  username,
  balance,
  affiliate_id
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;


-- name: AddBalance :one
UPDATE users
SET balance = balance + $2
WHERE id = $1
RETURNING *;


-- name: DeductBalance :one
UPDATE users
SET balance = balance - $2
WHERE id = $1
RETURNING *;

-- name: UpdateUserAffiliate :one
UPDATE users
SET affiliate_id = $2          
WHERE id = $1
RETURNING *;

-- name: CountUsers :one
SELECT COUNT(*) FROM users;

-- name: GetBalance :one
SELECT balance FROM users
WHERE id = $1 LIMIT 1;

