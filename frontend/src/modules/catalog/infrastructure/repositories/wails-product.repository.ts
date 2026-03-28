import {
  CreateProduct,
  FindAllProducts,
  UpdateProduct,
} from "@wails/go/handlers/CatalogProductHandler";

import { mapProductToDomain } from "@catalog/infrastructure/mappers/product.mappers";
import type {
  CreateProductCommand,
  ProductRepository,
  UpdateProductCommand,
} from "@catalog/domain/ports/product.repository";
import type { Product } from "@catalog/domain/entities/product.entity";

export class WailsProductRepository implements ProductRepository {
  public async findAllProducts(): Promise<Product[]> {
    const products = await FindAllProducts();

    return products.map(mapProductToDomain);
  }

  public async createProduct(cmd: CreateProductCommand): Promise<Product> {
    const product = await CreateProduct(cmd);

    return mapProductToDomain(product);
  }

  public async updateProduct(cmd: UpdateProductCommand): Promise<Product> {
    const product = await UpdateProduct(cmd);

    return mapProductToDomain(product);
  }
}

export const wailsProductRepository = new WailsProductRepository();
