import { createFileRoute, redirect } from "@tanstack/react-router";

import { appStateQueryOptions } from "@/integrations/iam";

import { SignUpScreen } from "@iam/presentation/screens/signup.screen";
import { activeOperatorIsRoot } from "@iam/domain/entities/app-session.entity";
import { organizationQueryOptions } from "@iam/application/hooks/use-organization-queries";

export const Route = createFileRoute("/auth/sign-up")({
  component: RouteComponent,
  beforeLoad: async ({ context }) => {
    const appState =
      await context.queryClient.ensureQueryData(appStateQueryOptions);

    const activeOperator = appState?.activeOperator;

    if (!activeOperator) return;

    if (activeOperatorIsRoot(appState)) {
      const organizations = await context.queryClient.ensureQueryData(
        organizationQueryOptions.findManyOrganizationsByRootOperator,
      );

      if (organizations.length === 0)
        throw redirect({
          to: "/register-organization",
        });

      throw redirect({
        to: "/select-organization",
      });
    }

    // TODO: Here the operator is not root, redirect to unique organization
    throw redirect({
      to: "/select-organization",
    });
  },
});

function RouteComponent() {
  return <SignUpScreen />;
}
