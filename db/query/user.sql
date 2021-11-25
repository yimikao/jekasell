-- name: CreateUser :one
INSERT into users (
    name,
    email,
    password
) VALUES (
    $1, $2, $3
) RETURNING *;


-- name: GetUser :one
SELECT * FROM users 
WHERE id = $1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY name;

-- name: UpdateUser :one
UPDATE users
SET password = $2
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;