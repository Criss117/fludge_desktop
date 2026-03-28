--------------------------------------------------------------------------------
-- Category Queries
--------------------------------------------------------------------------------

-- name: CreateCategory :exec
INSERT INTO category (
  id, 
  name, 
  description, 
  organization_id, 
  created_at, 
  updated_at
) 
VALUES (?, ?, ?, ?, ?, ?);

-- name: UpdateCategory :exec
UPDATE category 
SET 
  name = ?, 
  description = ?, 
  updated_at = ?
WHERE id = sqlc.arg(category_id) 
AND organization_id = sqlc.arg(organization_id) 
AND deleted_at IS NULL;

-- name: DeleteCategory :exec
DELETE FROM category
WHERE id = sqlc.arg(category_id)
AND organization_id = sqlc.arg(organization_id);

-- name: DeleteCategories :exec
DELETE FROM category
WHERE id IN (sqlc.slice(category_ids))
AND organization_id = sqlc.arg(organization_id);

-- name: FindOneCategory :one
SELECT * FROM category
WHERE id = sqlc.arg(category_id) 
AND organization_id = sqlc.arg(organization_id) 
AND deleted_at IS NULL;

-- name: FindAllCategories :many
SELECT * FROM category
WHERE organization_id = ? AND deleted_at IS NULL;

-- name: FindOneCategoryByName :one
SELECT * FROM category
WHERE name = sqlc.arg(category_name) 
AND organization_id = sqlc.arg(organization_id) 
AND deleted_at IS NULL;

--------------------------------------------------------------------------------
-- Product Queries
--------------------------------------------------------------------------------

-- name: CreateProduct :exec
INSERT INTO product (
  id, 
  sku, 
  name, 
  description, 
  wholesale_price, 
  sale_price, 
  cost_price, 
  category_id, 
  supplier_id, 
  organization_id, 
  created_at, 
  updated_at
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: UpdateProduct :exec
UPDATE product 
SET 
  sku = ?, 
  name = ?, 
  description = ?, 
  wholesale_price = ?, 
  sale_price = ?, 
  cost_price = ?, 
  category_id = ?,
  supplier_id = ?,
  updated_at = ?
WHERE id = ? AND organization_id = ? AND deleted_at IS NULL;

-- name: DeleteProduct :exec
DELETE FROM product
WHERE id = sqlc.arg(product_id)
AND organization_id = sqlc.arg(organization_id);

-- name: FindOneProduct :one
SELECT * FROM product
WHERE product.id = sqlc.arg(product_id) 
AND product.organization_id = sqlc.arg(organization_id) 
AND product.deleted_at IS NULL;

-- name: FindAllProducts :many
SELECT product.*, inventory_item.stock, inventory_item.min_stock
FROM product
INNER JOIN inventory_item ON inventory_item.product_id = product.id
WHERE product.organization_id = ? AND deleted_at IS NULL;

-- name: ExistsProductByNameOrSku :many
SELECT name, sku FROM product
WHERE (
  lower(name) = lower(sqlc.arg(name))
  OR sku = sqlc.arg(sku)
)
AND organization_id = sqlc.arg(organization_id)
AND (
  sqlc.narg(product_id) IS NULL
  OR id != sqlc.narg(product_id)
)
AND deleted_at IS NULL;

--------------------------------------------------------------------------------
-- Inventory Items Queries
--------------------------------------------------------------------------------

-- name: CreateInventoryItem :exec
INSERT INTO inventory_item (
  product_id, 
  organization_id, 
  stock, 
  min_stock, 
  created_at, 
  updated_at
) VALUES (?, ?, ?, ?, ?, ?);

-- name: UpdateInventoryItem :exec
UPDATE inventory_item 
SET stock = ?, min_stock = ?, updated_at = ?
WHERE product_id = ? AND organization_id = ?;

-- name: FindOneInventoryItem :one
SELECT * FROM inventory_item
WHERE inventory_item.product_id = sqlc.arg(product_id) 
AND inventory_item.organization_id = sqlc.arg(organization_id);