-------------------------------------------------------------------------------
-- Operators
-------------------------------------------------------------------------------

-- name: FindAllOperators :many
SELECT * FROM operator;

-- name: FindOneOperatorByEmail :many
SELECT * FROM operator 
WHERE email = ?
LIMIT 1;

-- name: FindOneOperatorByUsername :many
SELECT * FROM operator 
WHERE username = ?
LIMIT 1;

-- name: FindOneOperatorById :many
SELECT * FROM operator 
WHERE id = ?
LIMIT 1;

-- name: FindManyOperatorsByEmailOrUsername :many
SELECT * FROM operator 
WHERE email = ? OR username = ?;

-- name: CreateOperator :exec
INSERT INTO operator (id, name, username, email, pin, is_root, created_at, updated_at) 
VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-------------------------------------------------------------------------------
-- AppState
-------------------------------------------------------------------------------

-- name: FindAppState :one
SELECT * FROM app_state WHERE id = "local" LIMIT 1;

-- name: UpdateAppState :exec
UPDATE app_state SET active_organization_id = ?, active_operator_id = ?, updated_at = ? WHERE id = "local";

-------------------------------------------------------------------------------
-- Organization
-------------------------------------------------------------------------------

-- name: FindOneOrganizationById :many
SELECT * FROM organization WHERE id = ?;

-- name: FindManyOrganizationsByOperatorId :many
SELECT organization.* FROM member
INNER JOIN organization ON organization.id = member.organization_id
WHERE operator_id = ?;

-- name: FindManyOrganizationsBy :many
SELECT * FROM organization
WHERE organization.slug = ? OR organization.legal_name = ? OR organization.name = ?;

-- name: CreateOrganization :exec
INSERT INTO organization (id, name, slug, legal_name, address, logo, contact_phone, contact_email, created_at, updated_at) 
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-------------------------------------------------------------------------------
-- Member
-------------------------------------------------------------------------------

-- name: FindAllMembersByOrganizationId :many
SELECT * FROM member WHERE organization_id = ?;

-- name: FindOneMemberById :many
SELECT * FROM member m WHERE m.id = ?;

-- name: CreateMember :exec
INSERT INTO member (id, organization_id, operator_id, role, created_at, updated_at) 
VALUES (?, ?, ?, ?, ?, ?);

-------------------------------------------------------------------------------
-- Team
-------------------------------------------------------------------------------

-- name: FindAllTeamsByOrganizationId :many
SELECT * FROM team WHERE organization_id = ?;

-- name: FindAllTeamsMembersByTeamId :many
SELECT * FROM team_member WHERE team_id = ?;

-- name: FindAllTeamsByOperatorId :many
SELECT team.*
FROM team_member
INNER JOIN team ON team.id = team_member.team_id
WHERE team_member.operator_id = ?
GROUP BY team.id;

-------------------------------------------------------------------------------
-- Prodcut
-------------------------------------------------------------------------------

-- name: FindOneProductBySku :many
SELECT * FROM product WHERE sku = ? AND organization_id = ? LIMIT 1;

-- name: FindOneProductByName :many
SELECT * FROM product WHERE lower(name) = lower(?) AND organization_id = ? LIMIT 1;

-- name: FindAllProductsByOrganizationId :many
SELECT * FROM product WHERE organization_id = ?;

-- name: CreateProduct :exec
INSERT INTO product (id, sku, name, description, wholesale_price, sale_price, cost_price, stock, min_stock, category_id, organization_id, supplier_id, created_at, updated_at) 
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: UpdateProduct :exec
UPDATE product 
SET sku = ?, 
name = ?, 
description = ?, 
wholesale_price = ?, 
sale_price = ?, 
cost_price = ?, 
stock = ?, 
min_stock = ?, 
category_id = ?, 
supplier_id = ?, 
updated_at = ? 
WHERE id = ? AND organization_id = ?;

-- name: FindOneProductById :many
SELECT * FROM product WHERE id = ? AND organization_id = ? LIMIT 1;

-------------------------------------------------------------------------------
-- Category
-------------------------------------------------------------------------------

-- name: FindAllCategories :many
SELECT * FROM category WHERE organization_id = ?;

-- name: FindOneCategoryById :many
SELECT * FROM category WHERE id = ? AND organization_id = ? LIMIT 1;

-- name: FindOneCategoryByName :many
SELECT * FROM category WHERE lower(name) = lower(?) AND organization_id = ? LIMIT 1;

-- name: CreateCategory :exec
INSERT INTO category (id, name, description, organization_id, created_at, updated_at) 
VALUES (?, ?, ?, ?, ?, ?);

-- name: UpdateCategory :exec
UPDATE category 
SET name = ?, 
description = ?, 
updated_at = ? 
WHERE id = ? AND organization_id = ?;

-- name: SoftDeleteCategory :exec
UPDATE category 
SET deleted_at = ? 
WHERE id = ? AND organization_id = ?;

-- name: DeleteCategory :exec
DELETE FROM category 
WHERE id = ? AND organization_id = ?;

-- name: DeleteManyCategories :exec
DELETE FROM category 
WHERE id IN (sqlc.slice('ids')) AND organization_id = ?;