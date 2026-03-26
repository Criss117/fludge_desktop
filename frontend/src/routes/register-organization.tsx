import { createFileRoute, redirect } from "@tanstack/react-router";

import { appStateQueryOptions } from "@/integrations/iam";

import { RegisterOrganizationScreen } from "@organizations/presentation/screens/register-organization.screen";
import { FindManyOrganizationsByRootOperator } from "@wails/go/iam/IamHandler";

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

    const organizations = await FindManyOrganizationsByRootOperator();

    return {
      organizations,
    };
  },
  loader: async ({ context }) => {
    return context.organizations;
  },
});

function RouteComponent() {
  const organizations = Route.useLoaderData();

  return <RegisterOrganizationScreen organizations={organizations} />;
}
