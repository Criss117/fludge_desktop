import { Suspense } from "react";

import { FiltersProvider } from "@shared/store/filters.store";

import { ProductsHeaderSection } from "@catalog/presentation/sections/products-header.section";
import { ProductsTableSection } from "@catalog/presentation/sections/products-table.section";
import { ProductsFiltersSection } from "@/modules/catalog/presentation/sections/products-filters.section";
import {
  UpdateProductForm,
  UpdateProductRoot,
} from "@catalog/presentation/components/update-product";
import { useCountProductsQuery } from "@catalog/application/hooks/use-products-queries";

export function ProductsScreen() {
  const totalProducts = useCountProductsQuery();

  return (
    <div className="px-5 mt-4 mb-8 space-y-8">
      <ProductsHeaderSection totalProducts={totalProducts} />
      <div className="space-y-4">
        <FiltersProvider>
          <ProductsFiltersSection totalProducts={totalProducts} />
          <Suspense>
            <UpdateProductRoot>
              <ProductsTableSection />
              <UpdateProductForm />
            </UpdateProductRoot>
          </Suspense>
          <ProductsFiltersSection totalProducts={totalProducts} />
        </FiltersProvider>
      </div>
    </div>
  );
}
