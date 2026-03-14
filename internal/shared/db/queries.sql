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