import { IamProvider } from "./iam";
import { QueryProvider } from "./ts-query";
import { WailsProvider } from "./wails";

export function Integrations({ children }: { children: React.ReactNode }) {
  return (
    <WailsProvider>
      <QueryProvider>
        <IamProvider>{children}</IamProvider>
      </QueryProvider>
    </WailsProvider>
  );
}
