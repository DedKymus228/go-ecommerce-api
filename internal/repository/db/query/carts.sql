-- name: GetCartByUserID :one
SELECT * FROM carts
WHERE user_id = $1 LIMIT 1;

-- name: CreateCart :one
INSERT INTO carts (user_id)
VALUES ($1)
RETURNING *;

-- name: AddCartItem :one
INSERT INTO cart_items (cart_id, product_id, quantity)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetCartItems :many
SELECT
    p.id,
    p.name,
    p.price,
    ci.quantity
FROM cart_items ci
JOIN products p ON ci.product_id = p.id
WHERE ci.cart_id = $1;

-- name: DeleteCartItem :exec
DELETE FROM cart_items
WHERE cart_id = $1 AND product_id = $2;

-- name: ClearCart :exec
DELETE FROM cart_items
WHERE cart_id = $1;

-- name: GetCartItemByProduct :one
SELECT * FROM cart_items
WHERE cart_id = $1 AND product_id = $2;

-- name: UpdateCartItemQuantity :exec
UPDATE cart_items
SET quantity = $3
WHERE cart_id = $1 AND product_id = $2;
