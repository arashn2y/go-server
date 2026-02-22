-- name: GetAllPermissions :many
SELECT * FROM "permission";

-- name: GetPermissionByID :one
SELECT * FROM "permission" WHERE id = $1;

-- name: GetPermissionByName :one
SELECT * FROM "permission" WHERE name = $1;

-- name: CreatePermission :exec
INSERT INTO "permission" (name) VALUES ($1);

-- name: UpdatePermission :execrows
UPDATE "permission" SET name = $2 WHERE id = $1;

-- name: DeletePermission :execrows
DELETE FROM "permission" WHERE id = $1;

-- name: GetPermissionsByRoleID :many
SELECT p.* FROM "permission" p
JOIN role_permission rp ON p.id = rp.permission_id
WHERE rp.role_id = $1 AND rp.resource = $2;

-- name: GetPermissionsByRoleIDAndResource :many
SELECT p.* FROM "permission" p
JOIN role_permission rp ON p.id = rp.permission_id
WHERE rp.role_id = $1 AND rp.resource = $2;

-- name: AssignPermissionToRole :exec
INSERT INTO role_permission (role_id, resource, permission_id) VALUES ($1, $2, $3);

-- name: RevokePermissionFromRole :execrows
DELETE FROM role_permission WHERE role_id = $1 AND resource = $2 AND permission_id = $3;
