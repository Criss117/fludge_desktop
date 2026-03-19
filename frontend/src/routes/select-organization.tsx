import { createFileRoute, redirect } from "@tanstack/react-router";

import { appStateQueryOptions } from "@/integrations/iam";

import { SelectOrganizationScreen } from "@organizations/presentation/screens/select-organization.screen";

export const Route = createFileRoute("/select-organization")({
  component: RouteComponent,
  beforeLoad: async ({ context }) => {
    const appState =
      await context.queryClient.ensureQueryData(appStateQueryOptions);

    const activeOperator = appState?.activeOperator;

    if (!activeOperator)
      throw redirect({
        to: "/",
      });

    if (activeOperator.isMemberIn.length === 0)
      throw redirect({
        to: "/register-organization",
      });

    return {
      activeOperator,
    };
  },
  loader: async ({ context }) => {
    return context.activeOperator;
  },
});

function RouteComponent() {
  const activeOperator = Route.useLoaderData();

  return <SelectOrganizationScreen organizations={activeOperator.isMemberIn} />;
}
