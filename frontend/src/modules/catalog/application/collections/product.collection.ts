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

import { productService } from "@catalog/application/container";
import type { Product } from "@catalog/domain/entities/product.entity";

type ProductCollection = Collection<
  Product,
  string | number,
  QueryCollectionUtils<Product, string | number, Product, unknown>,
  never,
  Product
> &
  NonSingleResult;

const collectionsCache = new Map<string, ProductCollection>();

export function productCollectionBuilder(orgId: string) {
  if (!collectionsCache.has(orgId)) {
    const newProductCollection = createCollection(
      queryCollectionOptions<Product>({
        queryKey: ["organization", orgId, "products"],
        queryFn: () => {
          return productService.findAllProducts();
        },
        getKey: (p) => p.id,
        queryClient,

        onInsert: async ({ transaction, collection }) => {
          const values = transaction.mutations[0].modified;

          const cretedProduct = await productService.createProduct({
            costPrice: values.costPrice,
            name: values.name,
            description: values.description,
            minStock: values.minStock,
            salePrice: values.salePrice,
            sku: values.sku,
            stock: values.stock,
            wholesalePrice: values.wholesalePrice,
            categoryId: values.categoryId,
          });

          collection.utils.writeInsert(cretedProduct);

          return { refetch: false };
        },

        onUpdate: async ({ transaction, collection }) => {
          const mutations = transaction.mutations;

          const changes = mutations[0].changes;
          const original = mutations[0].original;

          const productToUpdate = {
            ...original,
            ...changes,
            description: changes.description ?? original.description,
            minStock: changes.minStock ?? original.minStock,
          };

          const updatedProduct =
            await productService.updateProduct(productToUpdate);

          collection.utils.writeUpdate(updatedProduct);
          return { refetch: false };
        },

        // onDelete: async ({ transaction, collection }) => {
        //   const mutations = transaction.mutations;

        //   const productIds = mutations.map((m) => ({ id: m.original.id }));

        //   await orpc.inventory.products.delete.call(productIds);

        //   collection.utils.writeDelete(productIds.map((m) => m.id));
        //   return { refetch: false };
        // },
      }),
    );

    collectionsCache.set(orgId, newProductCollection);
  }

  return collectionsCache.get(orgId)!;
}
