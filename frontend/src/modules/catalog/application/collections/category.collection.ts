import {
  createCollection,
  type Collection,
  type NonSingleResult,
} from "@tanstack/db";
import {
  queryCollectionOptions,
  type QueryCollectionUtils,
} from "@tanstack/query-db-collection";

import { queryClient } from "@/integrations/ts-query";

import type { Category } from "@catalog/domain/entities/category.entity";
import { categoryService } from "@catalog/application/container";

type CategoryCollection = Collection<
  Category,
  string | number,
  QueryCollectionUtils<Category, string | number, Category, unknown>,
  never,
  Category
> &
  NonSingleResult;

const collectionsCache = new Map<string, CategoryCollection>();

export function categoryCollectionBuilder(orgId: string) {
  if (!collectionsCache.has(orgId)) {
    const newCategoryCollection = createCollection(
      queryCollectionOptions<Category>({
        queryKey: ["organization", orgId, "categories"],
        queryFn: () => categoryService.findAllCategories(),
        getKey: (p) => p.id,
        queryClient,

        onInsert: async ({ transaction, collection }) => {
          const values = transaction.mutations[0].modified;

          const cretedCategory = await categoryService.createCategory({
            name: values.name,
            description: values.description,
          });

          collection.utils.writeInsert(cretedCategory);

          return { refetch: false };
        },

        onUpdate: async ({ transaction, collection }) => {
          const changes = transaction.mutations[0].changes;
          const original = transaction.mutations[0].original;

          const categoryToUpdate = {
            ...original,
            ...changes,
            description: changes.description ?? original.description,
          };

          const updatedCategory =
            await categoryService.updateCategory(categoryToUpdate);

          collection.utils.writeUpdate(updatedCategory);

          return { refetch: false };
        },
      }),
    );

    collectionsCache.set(orgId, newCategoryCollection);
  }

  return collectionsCache.get(orgId)!;
}
