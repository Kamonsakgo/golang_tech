-- name: ListProduct :many
SELECT * FROM product
ORDER BY id;
-- name: GetProduct :one
SELECT * FROM product
WHERE id = $1 LIMIT 1;
-- name: CreateProduct :one
INSERT INTO product (
  name,
  quantity,
  price
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteProduct :exec
UPDATE product
SET quantity = quantity - $2
WHERE id = $1 AND quantity >= $2;

-- name: GetProduct_quantity :one
SELECT quantity FROM product
WHERE id = $1 LIMIT 1;