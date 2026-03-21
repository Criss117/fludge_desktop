import { useMutation } from "@tanstack/react-query";

import { useProductCollection } from "@catalog/application/hooks/use-product-collection";
import type { CreateProductSchema } from "@catalog/application/validators/product.validators";

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

  return { create };
}
