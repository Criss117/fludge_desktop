import { useAppState } from "@/integrations/iam";
import { categoryCollectionBuilder } from "../collections/category.collection";

export function useCategoryCollection() {
  const { activeOrganization } = useAppState();

  return categoryCollectionBuilder(activeOrganization.id);
}
