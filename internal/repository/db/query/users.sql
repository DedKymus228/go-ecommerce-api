-- name: CreateUser :one
INSERT INTO users (
    email,
    password_hash,
    first_name,
    last_name
) VALUES (
             $1, $2, $3, $4
         ) RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: UpdateUserRole :exec
UPDATE users
SET role_id = $2
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
