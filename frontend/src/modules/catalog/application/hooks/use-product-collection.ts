import { useIam } from "@/integrations/iam";

import { productCollectionBuilder } from "@catalog/application/collections/product.collection";

export function useProductCollection() {
  const { appState } = useIam();

  const activeOrganization = appState.activeOrganization;

  if (!activeOrganization) throw new Error("No active organization");

  return productCollectionBuilder(activeOrganization.id);
}
