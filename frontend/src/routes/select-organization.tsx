import { createFileRoute, redirect } from "@tanstack/react-router";

import { appStateQueryOptions } from "@/integrations/iam";

import { SelectOrganizationScreen } from "@iam/presentation/screens/select-organization.screen";
import { organizationQueryOptions } from "@iam/application/hooks/use-organization-queries";

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

    const organizations = await context.queryClient.ensureQueryData(
      organizationQueryOptions.findManyOrganizationsByRootOperator,
    );

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
