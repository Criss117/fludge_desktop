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
import {
  FindAllCategories,
  CreateCategory,
  UpdateCategory,
} from "@wails/go/catalog/CatalogHandler";

export type Category = Awaited<ReturnType<typeof FindAllCategories>>[number] & {
  metadata?: {
    isPending?: boolean;
  };
};

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
        queryFn: () => FindAllCategories(),
        getKey: (p) => p.id,
        queryClient,

        onInsert: async ({ transaction, collection }) => {
          const values = transaction.mutations[0].modified;

          const cretedCategory = await CreateCategory({
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

          const updatedCategory = await UpdateCategory(categoryToUpdate);

          collection.utils.writeUpdate(updatedCategory);

          return { refetch: false };
        },
      }),
    );

    collectionsCache.set(orgId, newCategoryCollection);
  }

  return collectionsCache.get(orgId)!;
}
