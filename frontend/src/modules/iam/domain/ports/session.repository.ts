import type { Operator } from "@iam/domain/entities/operator.entity";
import type { AppSession } from "@iam/domain/entities/app-session.entity";
import type { Organization } from "@iam/domain/entities/organization.entity";

export type RegisterRootOperatorCommand = {
  name: string;
  email: string;
  pin: string;
  username: string;
};

export type SignInCommand = {
  username: string;
  pin: string;
};

export type SwitchOrganizationCommand = {
  organizationId: string;
};

export interface SessionRepository {
  registerRootOperator(cmd: RegisterRootOperatorCommand): Promise<Operator>;
  signIn(cmd: SignInCommand): Promise<Operator>;
  signOut(): Promise<void>;
  getAppSession(): Promise<AppSession>;
  switchOrganization(cmd: SwitchOrganizationCommand): Promise<Organization>;
}
