-- name: CreateUser :one
INSERT INTO users(name, email, password)
VALUES($1, $2, $3)
RETURNING *;

-- name: UpdateName :one
UPDATE users SET name=$2 WHERE id = $1
RETURNING *;

-- name: UpdateEmail :one
UPDATE users SET email=$2 WHERE id = $1
RETURNING *;

-- name: UpdateSenha :one
UPDATE users SET password=$2 WHERE id = $1
RETURNING *;

-- name: UserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: UserById :one
SELECT * FROM users WHERE id=$1;

-- name: ListAllUsers :many
SELECT * FROM users ORDER BY name ASC;

-- name: DesactiveUser :one
UPDATE users SET is_active=false, inactive_dt=NOW() WHERE id=$1
RETURNING *;