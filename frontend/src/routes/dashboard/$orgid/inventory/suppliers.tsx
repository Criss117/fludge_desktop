import { DashBoardHeader } from "@/modules/shared/components/dashboard-header";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/dashboard/$orgid/inventory/suppliers")({
  component: RouteComponent,
});

function RouteComponent() {
  const { orgid } = Route.useParams();

  return (
    <>
      <DashBoardHeader.Content orgid={orgid} currentPath="Suppliers">
        <DashBoardHeader.Suppliers />
      </DashBoardHeader.Content>
      <div>Hello "/dashboard/$orgid/inventory/suppliers"!</div>
    </>
  );
}
