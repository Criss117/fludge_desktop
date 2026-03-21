import { useMutation } from "@tanstack/react-query";

import { useProductCollection } from "@catalog/application/hooks/use-product-collection";
import type {
  CreateProductSchema,
  UpdateProductSchema,
} from "@catalog/application/validators/product.validators";

export function useMutateProducts() {
  const productCollection = useProductCollection();

  const create = useMutation({
    mutationKey: ["catalog", "create-product"],
    mutationFn: async (values: CreateProductSchema) => {
      const tx = productCollection.insert({
        ...values,
        id: crypto.randomUUID(),
        organizationId: "default",
        createdAt: new Date().getMilliseconds(),
        updatedAt: new Date().getMilliseconds(),
      });

      await tx.isPersisted.promise;
    },
  });

  const update = useMutation({
    mutationKey: ["catalog", "create-product"],
    mutationFn: async (values: UpdateProductSchema) => {
      const tx = productCollection.update(values.id, (draft) => {
        draft.name = values.name ?? draft.name;
        draft.sku = values.sku ?? draft.sku;
        draft.description = values.description;
        draft.costPrice = values.costPrice ?? draft.costPrice;
        draft.wholesalePrice = values.wholesalePrice ?? draft.wholesalePrice;
        draft.salePrice = values.salePrice ?? draft.salePrice;
        draft.stock = values.stock ?? draft.stock;
        draft.minStock = values.minStock ?? draft.minStock;
        draft.updatedAt = new Date().getMilliseconds();
      });

      await tx.isPersisted.promise;
    },
  });

  return { create, update };
}
