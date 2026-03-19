import { createFileRoute, redirect } from "@tanstack/react-router";
import { SignInScreen } from "@iam/presentation/screen/signin.screen";
import { appStateQueryOptions } from "@/integrations/iam";

export const Route = createFileRoute("/")({
  component: Index,
  beforeLoad: async ({ context }) => {
    const appState =
      await context.queryClient.ensureQueryData(appStateQueryOptions);

    const activeOperator = appState?.activeOperator;

    if (!activeOperator) return;

    const activeOrganization = appState?.activeOrganization;

    if (activeOrganization)
      throw redirect({
        to: "/dashboard/$orgid",
        params: { orgid: activeOrganization.id },
      });

    if (activeOperator.isRoot) {
      if (activeOperator.isMemberIn.length === 0)
        throw redirect({
          to: "/register-organization",
        });

      throw redirect({
        to: "/select-organization",
      });
    }

    // TODO: Here the operator is not root, redirect to unique organization in isMemberIn
    throw redirect({
      to: "/select-organization",
    });
  },
});

function Index() {
  return <SignInScreen />;
}
