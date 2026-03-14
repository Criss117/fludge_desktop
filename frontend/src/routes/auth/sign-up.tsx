import { createFileRoute } from "@tanstack/react-router";
import { SignUpScreen } from "@/modules/auth/presentation/screen/signup.screen";

export const Route = createFileRoute("/auth/sign-up")({
  component: RouteComponent,
});

function RouteComponent() {
  return <SignUpScreen />;
}
