import { AuthProvider } from "./auth";
import { QueryProvider } from "./ts-query";
import { WailsProvider } from "./wails";

export function Integrations({ children }: { children: React.ReactNode }) {
  return (
    <WailsProvider>
      <QueryProvider>
        <AuthProvider>{children}</AuthProvider>
      </QueryProvider>
    </WailsProvider>
  );
}
