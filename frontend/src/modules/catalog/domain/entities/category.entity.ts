export interface Category {
  id: string;
  name: string;
  description?: string;
  organizationId: string;
  createdAt: number;
  updatedAt: number;
  deletedAt?: number;
}
