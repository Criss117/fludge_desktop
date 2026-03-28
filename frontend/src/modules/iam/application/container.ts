import { wailsSessionRepository } from "@iam/infrastructure/repositories/wails-session.repository";
import { SessionService } from "@iam/application/services/session.service";

export const sessionService = new SessionService(wailsSessionRepository);
