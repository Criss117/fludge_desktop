import { createFileRoute, redirect } from "@tanstack/react-router";

import { appStateQueryOptions } from "@/integrations/iam";

import { RegisterOrganizationScreen } from "@iam/presentation/screens/register-organization.screen";
import { organizationQueryOptions } from "@iam/application/hooks/use-organization-queries";

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

    const organizations = await context.queryClient.ensureQueryData(
      organizationQueryOptions.findManyOrganizationsByRootOperator,
    );

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
