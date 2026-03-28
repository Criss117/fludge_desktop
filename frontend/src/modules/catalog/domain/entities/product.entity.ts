export interface Product {
  id: string;
  sku: string;
  name: string;
  description?: string;
  wholesalePrice: number;
  salePrice: number;
  costPrice: number;
  categoryId?: string;
  organizationId: string;
  supplierId?: string;
  createdAt: number;
  updatedAt: number;
  deletedAt?: number;
  stock: number;
  minStock: number;
}
