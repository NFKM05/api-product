-- name: CreateProduct :one
INSERT INTO products(product_name, quantity, price, description)
VALUES($1, $2, $3, $4)
RETURNING *;

-- name: UpdateProduct :one
UPDATE products
SET product_name=$2, quantity=$3, price=$4, description=$5
WHERE id = $1
RETURNING *;

-- name: ProductById :one
SELECT * FROM products WHERE id=$1;

-- name: ListAllProducts :many
SELECT * FROM products ORDER BY product_name ASC;