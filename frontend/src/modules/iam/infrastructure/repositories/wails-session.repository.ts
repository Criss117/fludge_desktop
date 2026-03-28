import {
  RegisterRootOperator,
  SignIn,
  SignOut,
  SwitchOrganization,
} from "@wails/go/handlers/IamSessionHandler";
import { GetAppSession } from "@wails/go/main/App";

import { mapOperatorToDomain } from "@iam/infrastructure/mappers/operator.mapper";
import { mapAppSessionToDomain } from "@iam/infrastructure/mappers/app-session.mapper";
import type { Operator } from "@iam/domain/entities/operator.entity";
import type { AppSession } from "@iam/domain/entities/app-session.entity";
import type { Organization } from "@iam/domain/entities/organization.entity";
import type {
  RegisterRootOperatorCommand,
  SessionRepository,
  SignInCommand,
  SwitchOrganizationCommand,
} from "@iam/domain/ports/session.repository";

export class WailsSessionRepository implements SessionRepository {
  public async registerRootOperator(
    cmd: RegisterRootOperatorCommand,
  ): Promise<Operator> {
    const res = await RegisterRootOperator(cmd);

    return mapOperatorToDomain(res);
  }

  public async signIn(cmd: SignInCommand): Promise<Operator> {
    const res = await SignIn(cmd);

    return mapOperatorToDomain(res);
  }

  public async signOut(): Promise<void> {
    return SignOut();
  }

  public async getAppSession(): Promise<AppSession> {
    const res = await GetAppSession();

    return mapAppSessionToDomain(res);
  }

  public async switchOrganization(
    cmd: SwitchOrganizationCommand,
  ): Promise<Organization> {
    return SwitchOrganization(cmd);
  }
}

export const wailsSessionRepository = new WailsSessionRepository();
