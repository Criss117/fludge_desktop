import type { Organization } from "@iam/domain/entities/organization.entity";

export type FindOneOrganizationCommand = {
  organizationId: string;
};

export type RegisterOrganizationCommand = {
  name: string;
  legalName: string;
  address: string;
  logo?: string;
  contactPhone?: string;
  contactEmail?: string;
};

export type UpdateOrganizationCommand = {
  id: string;
  name: string;
  legalName: string;
  address: string;
  logo?: string;
  contactPhone?: string;
  contactEmail?: string;
};

export interface OrganizationRepository {
  findManyOrganizationsByRootOperator(): Promise<Organization[]>;
  findOneOrganization(cmd: FindOneOrganizationCommand): Promise<Organization>;
  registerOrganization(cmd: RegisterOrganizationCommand): Promise<Organization>;
  updateOrganization(cmd: UpdateOrganizationCommand): Promise<Organization>;
}
