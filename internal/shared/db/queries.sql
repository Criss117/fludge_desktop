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