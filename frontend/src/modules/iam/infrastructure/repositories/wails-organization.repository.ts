import {
  FindManyOrganizationsByRootOperator,
  FindOneOrganization,
  RegisterOrganization,
  UpdateOrganization,
} from "@wails/go/handlers/IamOrganizationHandler";

import { mapOrganizationToDomain } from "@iam/infrastructure/mappers/organization.mapper";
import type {
  FindOneOrganizationCommand,
  OrganizationRepository,
  RegisterOrganizationCommand,
  UpdateOrganizationCommand,
} from "@iam/domain/ports/organization.repository";
import type { Organization } from "@iam/domain/entities/organization.entity";

export class WailsOrganizationRepository implements OrganizationRepository {
  public async findManyOrganizationsByRootOperator(): Promise<Organization[]> {
    const orgs = await FindManyOrganizationsByRootOperator();

    return orgs.map(mapOrganizationToDomain);
  }

  public async findOneOrganization(
    cmd: FindOneOrganizationCommand,
  ): Promise<Organization> {
    const org = await FindOneOrganization(cmd);

    return mapOrganizationToDomain(org);
  }

  public async registerOrganization(
    cmd: RegisterOrganizationCommand,
  ): Promise<Organization> {
    const org = await RegisterOrganization(cmd);

    return mapOrganizationToDomain(org);
  }

  public async updateOrganization(
    cmd: UpdateOrganizationCommand,
  ): Promise<Organization> {
    const org = await UpdateOrganization(cmd);

    return mapOrganizationToDomain(org);
  }
}

export const wailsOrganizationRepository = new WailsOrganizationRepository();
