import { CategoryService } from "@catalog/application/services/category.service";
import { ProductService } from "@catalog/application/services/product.service";

import { wailsCategoryRepository } from "@catalog/infrastructure/repositories/wails-category.repository";
import { wailsProductRepository } from "@catalog/infrastructure/repositories/wails-product.repository";

export const categoryService = new CategoryService(wailsCategoryRepository);

export const productService = new ProductService(wailsProductRepository);
