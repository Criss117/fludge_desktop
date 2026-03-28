import { Suspense } from "react";

import { FiltersProvider } from "@shared/store/filters.store";

import {
  ProductsHeaderSection,
  ProductsHeaderSectionSkeleton,
} from "@catalog/presentation/sections/products-header.section";
import { ProductsTableSection } from "@catalog/presentation/sections/products-table.section";
import {
  ProductsFiltersSection,
  ProductsFiltersSectionSkeleton,
} from "@/modules/catalog/presentation/sections/products-filters.section";
import {
  UpdateProductForm,
  UpdateProductRoot,
} from "@catalog/presentation/components/update-product";
import { useCountProductsQuery } from "@catalog/application/hooks/use-products-queries";
import { ProductsTableSkeleton } from "@catalog/presentation/components/products-table/skeleton";

export function ProductsScreen() {
  const totalProducts = useCountProductsQuery();

  return (
    <div className="px-5 mt-4 mb-8 space-y-8">
      <ProductsHeaderSection totalProducts={totalProducts} />
      <div className="space-y-4">
        <FiltersProvider>
          <ProductsFiltersSection totalProducts={totalProducts} />
          <UpdateProductRoot>
            <Suspense fallback={<ProductsTableSkeleton />}>
              <ProductsTableSection />
            </Suspense>
            <UpdateProductForm />
          </UpdateProductRoot>
          <ProductsFiltersSection totalProducts={totalProducts} />
        </FiltersProvider>
      </div>
    </div>
  );
}

export function ProductsScreenSkeleton() {
  return (
    <div className="px-5 mt-4 mb-8 space-y-8">
      <ProductsHeaderSectionSkeleton />
      <div className="space-y-4">
        <ProductsFiltersSectionSkeleton />
        <ProductsTableSkeleton />
        <ProductsFiltersSectionSkeleton />
      </div>
    </div>
  );
}
