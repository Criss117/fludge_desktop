import {
  CreateCategory,
  DeleteManyCategories,
  FindAllCategories,
  UpdateCategory,
} from "@wails/go/handlers/CatalogCategoryHandler";

import type {
  CategoryRepository,
  CreateCategoryCommand,
  DelelteManyCategoriesCommand,
  UpdateCategoryCommand,
} from "@catalog/domain/ports/category.repository";
import type { Category } from "@catalog/domain/entities/category.entity";
import { mapCategoryToDomain } from "@catalog/infrastructure/mappers/category.mappers";

export class WailsCategoryRepository implements CategoryRepository {
  public async craeteCategory(cmd: CreateCategoryCommand): Promise<Category> {
    const res = await CreateCategory(cmd);

    return mapCategoryToDomain(res);
  }

  public async deleteManyCategories(
    cmd: DelelteManyCategoriesCommand,
  ): Promise<void> {
    await DeleteManyCategories(cmd);
  }

  public async updateCategory(cmd: UpdateCategoryCommand): Promise<Category> {
    const res = await UpdateCategory(cmd);

    return mapCategoryToDomain(res);
  }

  public async findAllCategories(): Promise<Category[]> {
    const res = await FindAllCategories();

    return res.map(mapCategoryToDomain);
  }
}

export const wailsCategoryRepository = new WailsCategoryRepository();
