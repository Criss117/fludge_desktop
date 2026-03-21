import { useMutation } from "@tanstack/react-query";

import { useCategoryCollection } from "@catalog/application/hooks/use-category-collection";
import type {
  CreateCategorySchema,
  DeleteCategoriesSchema,
  UpdateCategorySchema,
} from "@catalog/application/validators/category.validators";
import { DeleteManyCategories } from "@wails/go/catalog/CatalogHandler";

export function useMutateCategories() {
  const categoryCollection = useCategoryCollection();

  const create = useMutation({
    mutationKey: ["categories", "create"],
    mutationFn: async (values: CreateCategorySchema) => {
      const tx = categoryCollection.insert({
        id: crypto.randomUUID(),
        name: values.name,
        organizationId: "default",
        description: values.description || undefined,
        createdAt: new Date().getMilliseconds(),
        updatedAt: new Date().getMilliseconds(),
        metadata: {
          isPending: true,
        },
      });

      await tx.isPersisted.promise;
    },
  });

  const update = useMutation({
    mutationKey: ["categories", "update"],
    mutationFn: async (values: UpdateCategorySchema) => {
      const tx = categoryCollection.update(values.id, (draft) => {
        draft.name = values.name || draft.name;
        draft.description = values.description || draft.description;
      });

      await tx.isPersisted.promise;
    },
  });

  const remove = useMutation({
    mutationKey: ["categories", "remove"],
    mutationFn: async (values: DeleteCategoriesSchema) => {
      await DeleteManyCategories(values);

      categoryCollection.utils.writeDelete(values.ids);
    },
  });

  return { create, update, remove };
}
