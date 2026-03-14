import { createFileRoute, redirect } from "@tanstack/react-router";
import { SignInScreen } from "@auth/presentation/screen/signin.screen";
import { appStateQueryOptions } from "@/integrations/auth";

export const Route = createFileRoute("/")({
  component: Index,
  beforeLoad: async ({ context }) => {
    const appState =
      await context.queryClient.ensureQueryData(appStateQueryOptions);

    if (appState.activeOperator) {
      throw redirect({
        to: "/select-organization",
      });
    }
  },
});

function Index() {
  return <SignInScreen />;
}
