import { createFileRoute, Outlet, redirect } from "@tanstack/react-router";

import { appStateQueryOptions } from "@/integrations/iam";

import { SwitchOrganization } from "@wails/go/iam/IamHandler";
import { SidebarInset, SidebarProvider } from "@shared/components/ui/sidebar";
import { AppSidebar } from "@shared/components/app-sidebar";

export const Route = createFileRoute("/dashboard/$orgid")({
  component: RouteComponent,
  beforeLoad: async ({ context, params }) => {
    const appState =
      await context.queryClient.ensureQueryData(appStateQueryOptions);

    if (!appState?.activeOperator)
      throw redirect({
        to: "/",
      });

    const newAppState = await SwitchOrganization({
      organizationId: params.orgid,
    });

    context.queryClient.setQueryData(
      appStateQueryOptions.queryKey,
      newAppState,
    );

    return {
      appState: newAppState,
    };
  },

  loader: async ({ context }) => {
    return context.appState;
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
