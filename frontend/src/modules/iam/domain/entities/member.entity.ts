export interface Member {
  id: string;
  organizationId: string;
  operatorId: string;
  role: string;
  createdAt: number;
  updatedAt: number;
  deletedAt?: number;
}
