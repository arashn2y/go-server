-- name: GetAllRoles :many
SELECT * FROM "role";

-- name: GetRoleByID :one
SELECT * FROM "role" WHERE id = $1;

-- name: GetRoleByName :one
SELECT * FROM "role" WHERE name = $1;

-- name: CreateRole :exec
INSERT INTO "role" (name) VALUES ($1);

-- name: UpdateRole :execrows
UPDATE "role" SET name = $2 WHERE id = $1;

-- name: DeleteRole :execrows
DELETE FROM "role" WHERE id = $1;

-- name: AssignRoleToUser :exec
INSERT INTO user_role (user_id, role_id) VALUES ($1, $2);

-- name: RemoveRoleFromUser :exec
DELETE FROM user_role WHERE user_id = $1 AND role_id = $2;