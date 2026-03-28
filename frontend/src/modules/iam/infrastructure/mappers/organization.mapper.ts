import type { responses } from "@wails/go/models";

import { mapMemberToDomain } from "@iam/infrastructure/mappers/member.mapper";
import { mapTeamToDomain } from "@iam/infrastructure/mappers/team.mapper";
import type { Organization } from "@iam/domain/entities/organization.entity";

export function mapOrganizationToDomain(
  organization: responses.Organization,
): Organization {
  const members = organization.Members
    ? organization.Members.map(mapMemberToDomain)
    : [];
  const teams = organization.Teams
    ? organization.Teams.map(mapTeamToDomain)
    : [];

  return {
    id: organization.id,
    name: organization.name,
    slug: organization.slug,
    logo: organization.logo,
    metadata: organization.metadata,
    legalName: organization.legalName,
    address: organization.address,
    contactPhone: organization.contactPhone,
    contactEmail: organization.contactEmail,
    createdAt: organization.createdAt,
    updatedAt: organization.updatedAt,
    deletedAt: organization.deletedAt,
    Members: members,
    Teams: teams,
  };
}
