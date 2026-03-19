import { createFileRoute, redirect } from "@tanstack/react-router";

import { appStateQueryOptions } from "@/integrations/iam";

import { RegisterOrganizationScreen } from "@organizations/presentation/screens/register-organization.screen";

export const Route = createFileRoute("/register-organization")({
  component: RouteComponent,
  beforeLoad: async ({ context }) => {
    const appState =
      await context.queryClient.ensureQueryData(appStateQueryOptions);

    const activeOperator = appState?.activeOperator;

    if (!activeOperator)
      throw redirect({
        to: "/",
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

  return (
    <RegisterOrganizationScreen organizations={activeOperator.isMemberIn} />
  );
}
