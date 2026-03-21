import { useLiveQuery } from "@tanstack/react-db";

import { useAppState } from "@/integrations/iam";

import { ProductsTable } from "@catalog/presentation/components/products-table";
import { useProductCollection } from "@catalog/application/hooks/use-product-collection";

export function ProductsTableSection() {
  const { activeOrganization } = useAppState();
  const productCollection = useProductCollection();

  const { data } = useLiveQuery((q) => q.from({ productCollection }));

  return (
    <ProductsTable.Root products={data} orgId={activeOrganization.id}>
      <ProductsTable.Content />
    </ProductsTable.Root>
  );
}
