import { createFileRoute, redirect } from "@tanstack/react-router";

import { appStateQueryOptions } from "@/integrations/iam";

import { SelectOrganizationScreen } from "@organizations/presentation/screens/select-organization.screen";
import { FindManyOrganizationsByRootOperator } from "@wails/go/handlers/IamOrganizationHandler";

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

    const organizations = await FindManyOrganizationsByRootOperator();

    if (organizations.length === 0) {
      throw redirect({
        to: "/register-organization",
      });
    }

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

  return <SelectOrganizationScreen organizations={organizations} />;

  // return null;
}
