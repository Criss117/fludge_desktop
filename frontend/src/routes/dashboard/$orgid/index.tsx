import { DashBoardHeader } from "@/modules/shared/components/dashboard-header";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/dashboard/$orgid/")({
  component: RouteComponent,
});

function RouteComponent() {
  const { orgid } = Route.useParams();

  return (
    <>
      <DashBoardHeader.Content orgid={orgid} currentPath="Home">
        <DashBoardHeader.Home />
      </DashBoardHeader.Content>
      <div>Hello "/dashboard/$orgslug/"!</div>
    </>
  );
}
