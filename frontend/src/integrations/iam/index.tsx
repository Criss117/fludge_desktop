import { createContext, use } from "react";
import {
  SignIn,
  GetAppState,
  SignUp,
  SignOut,
  SwitchOrganization,
} from "@wails/go/iam/IamHandler";
import {
  queryOptions,
  useMutation,
  useQueryClient,
  useSuspenseQuery,
  type UseMutationResult,
} from "@tanstack/react-query";
import type { responses } from "@wails/go/models";
import type {
  SignInSchema,
  SignUpSchema,
} from "@iam/application/validators/operator-form.validators";

type AppState = Awaited<ReturnType<typeof GetAppState>>;

interface Context {
  appState: AppState;
  signIn: UseMutationResult<
    responses.ResponseAppState,
    Error,
    SignInSchema,
    unknown
  >;
  signUp: UseMutationResult<
    responses.ResponseAppState,
    Error,
    SignUpSchema,
    unknown
  >;
  signOut: UseMutationResult<void, Error, void, unknown>;
  switchOrganization: UseMutationResult<void, Error, string, unknown>;
}

export const appStateQueryOptions = queryOptions({
  queryKey: ["auth", "app-state"],
  queryFn: GetAppState,
  refetchOnWindowFocus: true,
});

const IamContext = createContext<Context | null>(null);

export function useAppState() {
  const { appState } = useIam();

  const activeOrganization = appState.activeOrganization;
  const activeOperator = appState.activeOperator;

  if (!activeOrganization || !activeOperator)
    throw new Error("useAppState must be used within a IamProvider");

  return {
    activeOrganization,
    activeOperator,
  };
}

export function useIam() {
  const context = use(IamContext);

  if (!context) throw new Error("useAuth must be used within a IamProvider");

  return context;
}

export function IamProvider({ children }: { children: React.ReactNode }) {
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

  const signUp = useMutation({
    mutationKey: ["auth", "signup"],
    mutationFn: async (data: SignUpSchema) => {
      const appState = await SignUp(data);

      queryClient.setQueryData(appStateQueryOptions.queryKey, appState);

      return appState;
    },
  });

  const signOut = useMutation({
    mutationKey: ["auth", "signout"],
    mutationFn: async () => {
      await SignOut();

      queryClient.invalidateQueries(appStateQueryOptions);
    },
  });

  const switchOrganization = useMutation({
    mutationKey: ["auth", "switch-organization"],
    mutationFn: async (organizationId: string) => {
      const newAppState = await SwitchOrganization({
        organizationId,
      });

      queryClient.setQueryData(appStateQueryOptions.queryKey, newAppState);
    },
  });

  return (
    <IamContext.Provider
      value={{
        appState: authQuery.data,
        signIn,
        signUp,
        signOut,
        switchOrganization,
      }}
    >
      {children}
    </IamContext.Provider>
  );
}
