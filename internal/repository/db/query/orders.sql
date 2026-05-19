-- name: CreateOrder :one
INSERT INTO orders (
	user_id,
	status_id,
	total_amount,
	shipping_address
	) VALUES (
		$1, $2,$3,$4
        )
RETURNING *;

-- name: GetOrderByID :one
SELECT * FROM orders
WHERE id = $1 LIMIT 1;

-- name: ListUserOrders :many
SELECT * FROM orders
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: CreateOrderItem :one
INSERT INTO order_items (
    order_id,
    product_id,
    quantity,
    price_at_purchase
)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateOrderStatus :exec
UPDATE orders
SET status_id = $2 ,updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: GetOrderStatusByName :one
SELECT id, name FROM order_statuses
WHERE name = $1 LIMIT 1;
