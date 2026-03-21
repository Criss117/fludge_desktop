import { Suspense } from "react";

import { FiltersProvider } from "@shared/store/filters.store";

import { ProductsHeaderSection } from "@catalog/presentation/sections/product-header.section";
import { ProductsTableSection } from "@catalog/presentation/sections/product-table.section";

export function ProductsScreen() {
  return (
    <div className="px-5 mt-4 mb-8 space-y-8">
      <ProductsHeaderSection totalProducts={0} />
      <div className="space-y-4">
        <FiltersProvider>
          <Suspense>
            {/* <ProductsFiltersSection totalProducts={totalProducts} /> */}
          </Suspense>
          <Suspense>
            {/* <UpdateProductRoot> */}
            <ProductsTableSection />
            {/* <UpdateProductForm />
            </UpdateProductRoot> */}
          </Suspense>
          <Suspense>
            {/* <ProductsFiltersSection totalProducts={totalProducts} /> */}
          </Suspense>
        </FiltersProvider>
      </div>
    </div>
  );
}
