import type { Category } from "@catalog/domain/entities/category.entity";

export type CreateCategoryCommand = {
  name: string;
  description?: string;
};

export type DelelteManyCategoriesCommand = {
  ids: string[];
};

export type UpdateCategoryCommand = {
  id: string;
  name: string;
  description?: string;
};

export interface CategoryRepository {
  craeteCategory(cmd: CreateCategoryCommand): Promise<Category>;
  deleteManyCategories(cmd: DelelteManyCategoriesCommand): Promise<void>;
  updateCategory(cmd: UpdateCategoryCommand): Promise<Category>;
  findAllCategories(): Promise<Category[]>;
}
