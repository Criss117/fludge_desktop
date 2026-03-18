import { createContext, use } from "react";
import { SignIn, GetAppState } from "@wails/go/iam/IamHandler";
import {
  queryOptions,
  useMutation,
  useQueryClient,
  useSuspenseQuery,
  type UseMutationResult,
} from "@tanstack/react-query";
import type { responses } from "@wails/go/models";
import type { SignInSchema } from "@auth/application/validators/operator-form.validators";

type AppState = Awaited<ReturnType<typeof GetAppState>>;

interface Context {
  appState: AppState;
  signIn: UseMutationResult<
    responses.ResponseAppState,
    Error,
    {
      pin: string;
      username: string;
    },
    unknown
  >;
}

export const appStateQueryOptions = queryOptions({
  queryKey: ["auth", "app-state"],
  queryFn: GetAppState,
  refetchOnWindowFocus: true,
});

const AuthContext = createContext<Context | null>(null);

export function useAuth() {
  const context = use(AuthContext);

  if (!context) throw new Error("useAuth must be used within a AuthProvider");

  return context;
}

export function AuthProvider({ children }: { children: React.ReactNode }) {
  const queryClient = useQueryClient();
  const authQuery = useSuspenseQuery(appStateQueryOptions);

  const signIn = useMutation({
    mutationKey: ["auth", "signin"],
    mutationFn: async (data: SignInSchema) => {
      const appState = await SignIn(data);

      queryClient.setQueryData(appStateQueryOptions.queryKey, appState);

      return appState;
    },
  });

  return (
    <AuthContext.Provider value={{ appState: authQuery.data, signIn }}>
      {children}
    </AuthContext.Provider>
  );
}
