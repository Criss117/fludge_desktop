import type { responses } from "@wails/go/models";
import type { Product } from "@catalog/domain/entities/product.entity";

export function mapProductToDomain(product: responses.Product): Product {
  return {
    id: product.id,
    sku: product.sku,
    name: product.name,
    description: product.description,
    wholesalePrice: product.wholesalePrice,
    salePrice: product.salePrice,
    costPrice: product.costPrice,
    categoryId: product.categoryId,
    organizationId: product.organizationId,
    supplierId: product.supplierId,
    createdAt: product.createdAt,
    updatedAt: product.updatedAt,
    deletedAt: product.deletedAt,
    stock: product.stock,
    minStock: product.minStock,
  };
}
