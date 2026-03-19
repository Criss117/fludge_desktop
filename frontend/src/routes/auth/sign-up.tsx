import { createFileRoute, redirect } from "@tanstack/react-router";
import { SignUpScreen } from "@iam/presentation/screen/signup.screen";
import { appStateQueryOptions } from "@/integrations/iam";

export const Route = createFileRoute("/auth/sign-up")({
  component: RouteComponent,
  beforeLoad: async ({ context }) => {
    const appState =
      await context.queryClient.ensureQueryData(appStateQueryOptions);

    const activeOperator = appState?.activeOperator;

    if (!activeOperator) return;

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

function RouteComponent() {
  return <SignUpScreen />;
}
