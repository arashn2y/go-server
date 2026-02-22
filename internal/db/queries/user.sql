-- name: GetAllUsers :many
SELECT * FROM "user";

-- name: GetUserByID :one
SELECT * FROM "user" WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM "user" WHERE email = $1;

-- name: CreateUser :one
INSERT INTO "user" (name, email, password, is_active, role_id) VALUES ($1, $2, $3, $4, $5)
RETURNING id;

-- name: UpdateUser :execrows
UPDATE "user" SET name = $2, email = $3, password = $4, is_active = $5, updated_at = NOW() WHERE id = $1;

-- name: DeleteUser :execrows
DELETE FROM "user" WHERE id = $1;