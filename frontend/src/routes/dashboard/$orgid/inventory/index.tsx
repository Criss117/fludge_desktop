import { DashBoardHeader } from "@/modules/shared/components/dashboard-header";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/dashboard/$orgid/inventory/")({
  component: RouteComponent,
});

function RouteComponent() {
  const { orgid } = Route.useParams();
  return (
    <>
      <DashBoardHeader.Content orgid={orgid} currentPath="Inventory">
        <DashBoardHeader.Inventory />
      </DashBoardHeader.Content>
      <div>Hello "/dashboard/$orgid/inventory/"!</div>
    </>
  );
}
