import { Suspense } from "react";
import { createFileRoute } from "@tanstack/react-router";
import {
  ProductsScreen,
  ProductsScreenSkeleton,
} from "@catalog/presentation/screens/products.screen";
import { DashBoardHeader } from "@shared/components/dashboard-header";

export const Route = createFileRoute("/dashboard/$orgid/inventory/")({
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
