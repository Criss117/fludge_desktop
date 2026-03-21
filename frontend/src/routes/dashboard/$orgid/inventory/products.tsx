import { Suspense } from "react";
import { createFileRoute } from "@tanstack/react-router";

import { DashBoardHeader } from "@shared/components/dashboard-header";

import {
  ProductsScreen,
  ProductsScreenSkeleton,
} from "@catalog/presentation/screens/products.screen";

export const Route = createFileRoute("/dashboard/$orgid/inventory/products")({
  component: RouteComponent,
});

function RouteComponent() {
  const { orgid } = Route.useParams();

  return (
    <>
      <DashBoardHeader.Content orgid={orgid} currentPath="Products">
        <DashBoardHeader.Products />
      </DashBoardHeader.Content>
      <Suspense fallback={<ProductsScreenSkeleton />}>
        <ProductsScreen />
      </Suspense>
    </>
  );
}
