-- name: GetCommission :one
SELECT * FROM commission
WHERE id = $1 LIMIT 1;
-- name: ListCommission :many
SELECT * FROM commission
ORDER BY id;


-- name: Createcommission :one
INSERT INTO commission (
  order_id,
  affiliate_id,
  amount
) VALUES (
  $1, $2, $3
)
RETURNING *;