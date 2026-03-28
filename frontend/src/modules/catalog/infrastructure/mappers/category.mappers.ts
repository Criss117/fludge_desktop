import type { responses } from "@wails/go/models";
import type { Category } from "@catalog/domain/entities/category.entity";

export function mapCategoryToDomain(category: responses.Category): Category {
  return {
    id: category.id,
    name: category.name,
    description: category.description,
    organizationId: category.organizationId,
    createdAt: category.createdAt,
    updatedAt: category.updatedAt,
    deletedAt: category.deletedAt,
  };
}
