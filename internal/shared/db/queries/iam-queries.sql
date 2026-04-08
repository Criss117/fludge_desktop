--------------------------------------------------------------------------------
-- AppState Queries
--------------------------------------------------------------------------------
-- name: FindAppState :one
SELECT *
FROM app_state
WHERE id = 'local';

-- name: UpdateAppState :exec
UPDATE app_state
SET active_organization_id = ?,
  active_operator_id = ?,
  updated_at = ?
WHERE id = 'local';
-------------------------------------------------------------------------------
-- Operator Queries
-------------------------------------------------------------------------------
-- name: FindOneOperatorByEmail :one
SELECT *
FROM operator
WHERE email = ?
  AND deleted_at IS NULL;

-- name: FindOneOperatorByUsername :one
SELECT *
FROM operator
WHERE username = ?
  AND deleted_at IS NULL;

-- name: FindOneOperatorById :one
SELECT *
FROM operator
WHERE id = ?
  AND deleted_at IS NULL;

-- name: CreateOperator :exec
INSERT INTO operator (
    id,
    name,
    username,
    email,
    pin,
    operator_type,
    created_at,
    updated_at
  )
VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- name: UpdateOperator :exec
UPDATE operator
SET name = ?,
  username = ?,
  email = ?,
  pin = ?,
  operator_type = ?,
  updated_at = ?
WHERE id = ?
  AND deleted_at IS NULL;

-- name: DeleteOperator :exec
DELETE FROM operator
WHERE id = ?;

-- name: SoftDeleteOperator :exec
UPDATE operator
SET deleted_at = ?
WHERE id = ?;

--------------------------------------------------------------------------------
-- Organization Queries
--------------------------------------------------------------------------------
-- name: ExistsOrganization :one
SELECT COUNT(id) as total
FROM organization
WHERE (
    lower(name) = lower(sqlc.arg(name))
    OR lower(legal_name) = lower(sqlc.arg(legal_name))
    OR lower(slug) = lower(sqlc.arg(slug))
  )
  AND deleted_at IS NULL;

-- name: FindOneOrganization :one
SELECT *
FROM organization
WHERE id = ?
  AND deleted_at IS NULL;

-- name: FindManyOrganizationsByRootOperator :many
SELECT organization.*
FROM member
  INNER JOIN organization ON member.organization_id = organization.id
WHERE operator_id = ?
  AND role = 'ROOT';

-- name: CreateOrganization :exec
INSERT INTO organization (
    id,
    name,
    slug,
    logo,
    metadata,
    legal_name,
    address,
    contact_phone,
    contact_email,
    created_at,
    updated_at
  )
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: UpdateOrganization :exec
UPDATE organization
SET name = ?,
  slug = ?,
  logo = ?,
  metadata = ?,
  legal_name = ?,
  address = ?,
  contact_phone = ?,
  contact_email = ?,
  updated_at = ?
WHERE id = ?
  AND deleted_at IS NULL;

--------------------------------------------------------------------------------
-- Team Queries
--------------------------------------------------------------------------------
-- name: FindAllTeams :many
SELECT *
FROM team
WHERE organization_id = ?
  AND deleted_at IS NULL;

-- name: ExistsTeam :one
SELECT COUNT(id) as total
FROM team
WHERE lower(name) = lower(sqlc.arg(name))
  OR (
    organization_id = sqlc.arg(organization_id)
    AND deleted_at IS NULL
  );

-- name: CreateTeam :exec
INSERT INTO team (
    id,
    name,
    organization_id,
    permissions,
    description,
    created_at,
    updated_at
  )
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: DeleteTeam :exec
DELETE FROM team
WHERE id = ? 
  AND organization_id = ?;

-- name: UpdateTeam :exec
UPDATE team
SET name = ?,
  permissions = ?,
  description = ?,
  updated_at = ?
WHERE id = ?
  AND organization_id = ?
  AND deleted_at IS NULL;

-- name: FindAllTeamsMembersByTeam :many
SELECT *
FROM team_member
WHERE team_id = ?
  AND deleted_at IS NULL;

-- name: AddTeamMember :exec
INSERT INTO team_member (id, team_id, operator_id, created_at, updated_at)
VALUES (?, ?, ?, ?, ?);

-- name: DeleteTeamMembers :exec
DELETE FROM team_member
WHERE team_id = ?;

--------------------------------------------------------------------------------
-- Member Queries
--------------------------------------------------------------------------------
-- name: FindAllMembers :many
SELECT *
FROM member
WHERE organization_id = ?
  AND deleted_at IS NULL;
-- name: DeleteMember :exec
UPDATE member
SET deleted_at = ?
WHERE id = ?
  AND organization_id = ?;
-- name: CreateMember :exec
INSERT INTO member (
    id,
    organization_id,
    operator_id,
    role,
    created_at,
    updated_at
  )
VALUES (?, ?, ?, ?, ?, ?);