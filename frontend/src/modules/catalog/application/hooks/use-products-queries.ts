import {
  count,
  eq,
  ilike,
  or,
  Query,
  useLiveSuspenseQuery,
} from "@tanstack/react-db";

import { useProductCollection } from "./use-product-collection";
import { useCategoryCollection } from "./use-category-collection";

interface Filters {
  limit?: number;
  page?: number;
  name?: string;
  sku?: string;
  orderBy?: {
    stock?: "asc" | "desc" | null;
    costPrice?: "asc" | "desc" | null;
    salePrice?: "asc" | "desc" | null;
    wholesalePrice?: "asc" | "desc" | null;
  };
}

export function useCountProductsQuery() {
  const productsCollection = useProductCollection();

  const { data } = useLiveSuspenseQuery((q) =>
    q
      .from({ product: productsCollection })
      .select(({ product }) => ({
        total: count(product.id),
      }))
      .findOne(),
  );

  return data?.total || 0;
}

export function useFindManyProducts(filters: Filters) {
  const productsCollection = useProductCollection();
  const categoryCollection = useCategoryCollection();

  const limit = filters.limit || 50;
  const offset = (filters.page || 0) * limit;
  const name = filters.name || "";
  const sku = filters.sku || "";

  const orderByStock = filters.orderBy?.stock || null;
  const orderByCostPrice = filters.orderBy?.costPrice || null;
  const orderBySalePrice = filters.orderBy?.salePrice || null;
  const orderByWholesalePrice = filters.orderBy?.wholesalePrice || null;

  const anyOrderBy =
    !!orderByStock ||
    !!orderByCostPrice ||
    !!orderBySalePrice ||
    !!orderByWholesalePrice;

  const { data } = useLiveSuspenseQuery(
    (q) => {
      let query = new Query()
        .from({ product: productsCollection })
        .select(({ product }) => ({
          id: product.id,
          name: product.name,
          sku: product.sku,
          description: product.description,
          wholesalePrice: product.wholesalePrice,
          salePrice: product.salePrice,
          costPrice: product.costPrice,
          organizationId: product.organizationId,
          createdAt: product.createdAt,
          updatedAt: product.updatedAt,
          stock: product.stock,
          minStock: product.minStock,
          category: q
            .from({ c: categoryCollection })
            .select(({ c }) => ({
              id: c.id,
              name: c.name,
              description: c.description,
            }))
            .where(({ c }) => eq(c.id, product.categoryId))
            .findOne(),
        }));
      if (orderByCostPrice) {
        query = query.orderBy(
          ({ product }) => product.costPrice,
          orderByCostPrice,
        );
      }

      if (orderByStock) {
        query = query.orderBy(({ product }) => product.stock, orderByStock);
      }

      if (orderBySalePrice) {
        query = query.orderBy(
          ({ product }) => product.salePrice,
          orderBySalePrice,
        );
      }

      if (orderByWholesalePrice) {
        query = query.orderBy(
          ({ product }) => product.wholesalePrice,
          orderByWholesalePrice,
        );
      }

      if (!anyOrderBy)
        query = query.orderBy(({ product }) => product.createdAt, "desc");

      return query
        .where(({ product }) =>
          or(ilike(product.name, `%${name}%`), ilike(product.sku, `%${sku}%`)),
        )
        .limit(limit)
        .offset(offset);
    },
    [
      limit,
      offset,
      name,
      sku,
      orderByStock,
      orderByCostPrice,
      orderBySalePrice,
      orderByWholesalePrice,
    ],
  );

  return data;
}

export type ProductWithCategory = ReturnType<typeof useFindManyProducts>[0];
