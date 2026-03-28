export interface TeamMember {
  id: string;
  teamId: string;
  operatorId: string;
  organizationId: string;
  createdAt: number;
  updatedAt: number;
  deletedAt?: number;
}

export interface Team {
  id: string;
  name: string;
  organizationId: string;
  permissions: string[];
  description?: string;
  createdAt: number;
  updatedAt: number;
  deletedAt?: number;
  Members: TeamMember[];
}
