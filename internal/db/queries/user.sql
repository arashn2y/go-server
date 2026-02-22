-- name: GetAllUsers :many
SELECT * FROM "user";

-- name: GetUserByID :one
SELECT * FROM "user" WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM "user" WHERE email = $1;

-- name: CreateUser :exec
INSERT INTO "user" (name, email, password) VALUES ($1, $2, $3);

-- name: UpdateUser :execrows
UPDATE "user" SET name = $2, email = $3, password = $4 WHERE id = $1;

-- name: DeleteUser :execrows
DELETE FROM "user" WHERE id = $1;