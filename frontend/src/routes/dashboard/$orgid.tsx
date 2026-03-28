import { createFileRoute, Outlet, redirect } from "@tanstack/react-router";

import { appStateQueryOptions } from "@/integrations/iam";

import { SidebarInset, SidebarProvider } from "@shared/components/ui/sidebar";
import { AppSidebar } from "@shared/components/app-sidebar";

import { sessionService } from "@iam/application/container";

export const Route = createFileRoute("/dashboard/$orgid")({
  component: RouteComponent,
  beforeLoad: async ({ context, params }) => {
    const appState =
      await context.queryClient.ensureQueryData(appStateQueryOptions);

    if (!appState?.activeOperator)
      throw redirect({
        to: "/",
      });

    const activeOrganization = appState.activeOrganization;

    if (activeOrganization?.id === params.orgid) return;

    await sessionService.switchOrganization({
      organizationId: params.orgid,
    });

    await context.queryClient.refetchQueries(appStateQueryOptions);
  },
});

function RouteComponent() {
  const { orgid } = Route.useParams();

  return (
    <SidebarProvider>
      <AppSidebar orgId={orgid} />
      <SidebarInset>
        <main>
          <Outlet />
        </main>
      </SidebarInset>
    </SidebarProvider>
  );
}
