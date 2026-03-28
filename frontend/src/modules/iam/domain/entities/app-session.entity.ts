import type { Member } from "./member.entity";
import type { Operator } from "./operator.entity";
import type { Organization } from "./organization.entity";
import type { Team } from "./team.entity";

type ActiveOperator = Operator & {
  member: Member | null;
  teams: Team[] | null;
};

export interface AppSession {
  activeOrganization: Organization | null;
  activeOperator: ActiveOperator | null;
}
