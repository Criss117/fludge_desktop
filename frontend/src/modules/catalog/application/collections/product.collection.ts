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
  FindAllProducts,
  CreateProduct,
} from "@wails/go/catalog/CatalogHandler";

export type Product = Awaited<ReturnType<typeof FindAllProducts>>[number] & {
  metadata?: {
    isPending?: boolean;
  };
};

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
        queryFn: () => FindAllProducts(),
        getKey: (p) => p.id,
        queryClient,

        onInsert: async ({ transaction, collection }) => {
          const values = transaction.mutations[0].modified;

          const cretedProduct = await CreateProduct({
            costPrice: values.costPrice,
            name: values.name,
            description: values.description,
            minStock: values.minStock,
            salePrice: values.salePrice,
            sku: values.sku,
            stock: values.stock,
            wholesalePrice: values.wholesalePrice,
          });

          console.log("Created product:", cretedProduct);

          collection.utils.writeInsert(cretedProduct);

          return { refetch: false };
        },

        // onUpdate: async ({ transaction, collection }) => {
        //   const mutations = transaction.mutations;

        //   const toUpdate = mutations[0].changes;
        //   const productId = mutations[0].original.id;

        //   const updatedProduct = await orpc.inventory.products.update.call({
        //     id: productId,
        //     costPrice: toUpdate.costPrice,
        //     name: toUpdate.name,
        //     description: toUpdate.description ?? undefined,
        //     minStock: toUpdate.minStock,
        //     salePrice: toUpdate.salePrice,
        //     sku: toUpdate.sku,
        //     stock: toUpdate.stock,
        //     wholesalePrice: toUpdate.wholesalePrice,
        //   });

        //   collection.utils.writeUpdate(updatedProduct);
        //   return { refetch: false };
        // },

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
