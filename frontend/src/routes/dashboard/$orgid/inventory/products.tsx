import { DashBoardHeader } from "@/modules/shared/components/dashboard-header";
import { createFileRoute } from "@tanstack/react-router";

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
      <div>Hello "/dashboard/$orgid/inventory/products"!</div>
    </>
  );
}
