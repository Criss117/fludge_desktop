import type {
  FindOneOrganizationCommand,
  OrganizationRepository,
  RegisterOrganizationCommand,
  UpdateOrganizationCommand,
} from "@iam/domain/ports/organization.repository";
import type { Organization } from "@iam/domain/entities/organization.entity";

export class OrganizationService {
  private readonly organizationRepository: OrganizationRepository;

  constructor(organizationRepository: OrganizationRepository) {
    this.organizationRepository = organizationRepository;
  }

  public async findManyOrganizationsByRootOperator(): Promise<Organization[]> {
    return this.organizationRepository.findManyOrganizationsByRootOperator();
  }

  public async findOneOrganization(
    cmd: FindOneOrganizationCommand,
  ): Promise<Organization> {
    return this.organizationRepository.findOneOrganization(cmd);
  }

  public async registerOrganization(
    cmd: RegisterOrganizationCommand,
  ): Promise<Organization> {
    return this.organizationRepository.registerOrganization(cmd);
  }

  public async updateOrganization(
    cmd: UpdateOrganizationCommand,
  ): Promise<Organization> {
    return this.organizationRepository.updateOrganization(cmd);
  }
}
