import type {
  RegisterRootOperatorCommand,
  SessionRepository,
  SignInCommand,
  SwitchOrganizationCommand,
} from "@iam/domain/ports/session.repository";
import type { Operator } from "@iam/domain/entities/operator.entity";
import type { AppSession } from "@iam/domain/entities/app-session.entity";
import type { Organization } from "@iam/domain/entities/organization.entity";

export class SessionService {
  private readonly sessionRepository: SessionRepository;

  constructor(sessionRepository: SessionRepository) {
    this.sessionRepository = sessionRepository;
  }

  public async signIn(cmd: SignInCommand): Promise<Operator> {
    return this.sessionRepository.signIn(cmd);
  }

  public async signOut(): Promise<void> {
    return this.sessionRepository.signOut();
  }

  public async registerRootOperator(
    cmd: RegisterRootOperatorCommand,
  ): Promise<Operator> {
    return this.sessionRepository.registerRootOperator(cmd);
  }

  public async getAppSession(): Promise<AppSession> {
    return this.sessionRepository.getAppSession();
  }

  public async switchOrganization(
    cmd: SwitchOrganizationCommand,
  ): Promise<Organization> {
    return this.sessionRepository.switchOrganization(cmd);
  }
}
