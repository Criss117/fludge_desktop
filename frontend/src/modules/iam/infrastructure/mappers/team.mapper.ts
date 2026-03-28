import type { responses } from "@wails/go/models";
import type { Team, TeamMember } from "@iam/domain/entities/team.entity";

export function mapTeamMemberToDomain(
  member: responses.TeamMember,
): TeamMember {
  return {
    id: member.id,
    teamId: member.teamId,
    operatorId: member.operatorId,
    organizationId: member.organizationId,
    createdAt: member.createdAt,
    updatedAt: member.updatedAt,
  };
}

export function mapTeamToDomain(team: responses.Team): Team {
  return {
    id: team.id,
    name: team.name,
    organizationId: team.organizationId,
    permissions: team.permissions,
    description: team.description,
    createdAt: team.createdAt,
    updatedAt: team.updatedAt,
    deletedAt: team.deletedAt,
    Members: team.Members.map(mapTeamMemberToDomain),
  };
}
