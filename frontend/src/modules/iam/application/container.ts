import { wailsSessionRepository } from "@iam/infrastructure/repositories/wails-session.repository";
import { SessionService } from "@iam/application/services/session.service";
import { OrganizationService } from "@iam/application/services/organization.service";
import { wailsOrganizationRepository } from "@iam/infrastructure/repositories/wails-organization.repository";

export const sessionService = new SessionService(wailsSessionRepository);

export const organizationService = new OrganizationService(
  wailsOrganizationRepository,
);
