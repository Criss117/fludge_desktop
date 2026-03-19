import { createFileRoute, Outlet, redirect } from "@tanstack/react-router";

import { appStateQueryOptions } from "@/integrations/iam";

import { SwitchOrganization } from "@wails/go/iam/IamHandler";

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

    console.log(newAppState);

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
  const appState = Route.useLoaderData();

  return (
    <div>
      <pre>
        <code>{JSON.stringify(appState, null, 2)}</code>
      </pre>
      <Outlet />
    </div>
  );
}
