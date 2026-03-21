import { useFilters } from "@shared/store/filters.store";

import { ProductsTable } from "@catalog/presentation/components/products-table";
import { useFindManyProducts } from "@catalog/application/hooks/use-products-queries";

export function ProductsTableSection() {
  const { filters } = useFilters();

  const products = useFindManyProducts({
    limit: filters.limit,
    name: filters.query,
    sku: filters.query,
    orderBy: {
      costPrice: filters.orderBy.get("costPrice"),
      salePrice: filters.orderBy.get("salePrice"),
      stock: filters.orderBy.get("stock"),
      wholesalePrice: filters.orderBy.get("wholesalePrice"),
    },
    page: filters.page,
  });

  return (
    <ProductsTable.Root products={products}>
      <ProductsTable.Content />
    </ProductsTable.Root>
  );
}
