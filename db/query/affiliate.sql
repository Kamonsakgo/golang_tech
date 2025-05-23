-- name: Createaffiliate :one
INSERT INTO affiliate ( name, master_affiliate,balance)
VALUES ( $1, $2 , $3)
RETURNING *;

-- name: Getaffiliate :one
SELECT * FROM affiliate
WHERE id = $1 LIMIT 1;
-- name: GetaffiliateByname :one
SELECT * FROM affiliate
WHERE name = $1 LIMIT 1;
-- name: Listaffiliate :many
SELECT * FROM affiliate
ORDER BY id;

-- name: GetAffiliateChain :many
WITH RECURSIVE affiliate_chain AS (
  SELECT a.id, a.master_affiliate, a.balance
  FROM affiliate a
  WHERE a.id = $1

  UNION ALL

  SELECT a2.id, a2.master_affiliate, a2.balance
  FROM affiliate a2
  INNER JOIN affiliate_chain ac ON a2.id = ac.master_affiliate
)
SELECT ac.id, ac.master_affiliate, ac.balance
FROM affiliate_chain ac;


-- name: AddBalance_affiliate :exec
UPDATE affiliate
SET balance = balance + $2
WHERE id = $1 ;