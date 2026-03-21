import { Suspense } from "react";

import { FiltersProvider } from "@shared/store/filters.store";

import { ProductsHeaderSection } from "@/modules/catalog/presentation/sections/products-header.section";
import { ProductsTableSection } from "@/modules/catalog/presentation/sections/products-table.section";
import { ProductsFiltersSection } from "../sections/products-filter.section";

export function ProductsScreen() {
  return (
    <div className="px-5 mt-4 mb-8 space-y-8">
      <ProductsHeaderSection totalProducts={0} />
      <div className="space-y-4">
        <FiltersProvider>
          <ProductsFiltersSection totalProducts={0} />
          <Suspense>
            {/* <UpdateProductRoot> */}
            <ProductsTableSection />
            {/* <UpdateProductForm />
            </UpdateProductRoot> */}
          </Suspense>
          <ProductsFiltersSection totalProducts={0} />
        </FiltersProvider>
      </div>
    </div>
  );
}
