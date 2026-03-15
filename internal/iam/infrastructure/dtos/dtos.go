package dtos

import "desktop/internal/shared/db"

type DBActiveTeamDto struct {
	db.Team
	Members []db.TeamMember
}

type DBActiveOrganizationDto struct {
	Organization *db.Organization
	Members      []db.Member
	Teams        []DBActiveTeamDto
}
