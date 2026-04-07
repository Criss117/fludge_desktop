import type { Product } from "@catalog/domain/entities/product.entity";

export type CreateProductCommand = {
  name: string;
  sku: string;
  description?: string;
  wholesalePrice: number;
  salePrice: number;
  costPrice: number;
  stock: number;
  minStock: number;
  categoryId?: string;
  supplierId?: string;
};

export type UpdateProductCommand = {
  id: string;
  name: string;
  sku: string;
  description?: string;
  wholesalePrice: number;
  salePrice: number;
  costPrice: number;
  stock: number;
  minStock: number;
  categoryId?: string;
  supplierId?: string;
};

export interface ProductRepository {
  findAllProducts(): Promise<Product[]>;
  createProduct(cmd: CreateProductCommand): Promise<Product>;
  updateProduct(cmd: UpdateProductCommand): Promise<Product>;
  deleteProduct(id: string): Promise<void>;
}
