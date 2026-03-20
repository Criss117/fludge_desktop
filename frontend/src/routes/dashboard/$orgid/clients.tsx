import { DashBoardHeader } from "@/modules/shared/components/dashboard-header";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/dashboard/$orgid/clients")({
  component: RouteComponent,
});

function RouteComponent() {
  const { orgid } = Route.useParams();
  return (
    <>
      <DashBoardHeader.Content orgid={orgid} currentPath="Clients">
        <DashBoardHeader.Clients />
      </DashBoardHeader.Content>
      <div>Hello "/dashboard/$orgid/clients"!</div>
    </>
  );
}
