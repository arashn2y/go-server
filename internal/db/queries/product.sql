-- name: GetAllProducts :many
SELECT * FROM product;

-- name: GetProductByID :one
SELECT * FROM product WHERE id = $1;

-- name: CreateProduct :exec
INSERT INTO product (name, description, price_in_cents) VALUES ($1, $2, $3);

-- name: UpdateProduct :exec
UPDATE product SET name = $2, description = $3, price_in_cents = $4 WHERE id = $1;

-- name: DeleteProduct :execrows
DELETE FROM product WHERE id = $1;