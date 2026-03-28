import type { Member } from "./member.entity";
import type { Team } from "./team.entity";

export interface Organization {
  id: string;
  name: string;
  slug: string;
  logo?: string;
  metadata: number[];
  legalName: string;
  address: string;
  contactPhone?: string;
  contactEmail?: string;
  createdAt: number;
  updatedAt: number;
  deletedAt?: number;
  Members: Member[];
  Teams: Team[];
}
