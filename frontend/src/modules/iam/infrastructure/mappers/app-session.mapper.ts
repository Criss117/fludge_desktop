import type { appstate } from "@wails/go/models";

import { mapOperatorToDomain } from "@iam/infrastructure/mappers/operator.mapper";
import { mapMemberToDomain } from "@iam/infrastructure/mappers/member.mapper";
import { mapTeamToDomain } from "@iam/infrastructure/mappers/team.mapper";
import type { AppSession } from "@iam/domain/entities/app-session.entity";
import { mapOrganizationToDomain } from "./organization.mapper";

export function mapActiveOperatorToDomain(
  values: appstate.SessionStateResponse["activeOperator"],
): AppSession["activeOperator"] {
  if (!values?.operator) return null;

  const operator = mapOperatorToDomain(values.operator);
  const member = values.member ? mapMemberToDomain(values.member) : null;
  const teams = values.teams ? values.teams.map(mapTeamToDomain) : null;

  return {
    ...operator,
    member,
    teams,
  };
}

export function mapAppSessionToDomain(
  values: appstate.SessionStateResponse,
): AppSession {
  return {
    activeOperator: values.activeOperator
      ? mapActiveOperatorToDomain(values.activeOperator)
      : null,
    activeOrganization: values.activeOrganization
      ? mapOrganizationToDomain(values.activeOrganization)
      : null,
  };
}
