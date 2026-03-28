import { createContext, use } from "react";
import {
  queryOptions,
  useMutation,
  useQueryClient,
  useSuspenseQuery,
  type UseMutationResult,
} from "@tanstack/react-query";

import { sessionService } from "@iam/application/container";
import type {
  SignInSchema,
  SignUpSchema,
} from "@iam/application/validators/operator-form.validators";
import type { AppSession } from "@iam/domain/entities/app-session.entity";
import type { Operator } from "@iam/domain/entities/operator.entity";
import type { Organization } from "@iam/domain/entities/organization.entity";

interface Context {
  appState: AppSession;
  signIn: UseMutationResult<Operator, Error, SignInSchema, unknown>;
  signUp: UseMutationResult<Operator, Error, SignUpSchema, unknown>;
  signOut: UseMutationResult<void, Error, void, unknown>;
  switchOrganization: UseMutationResult<Organization, Error, string, unknown>;
}

export const appStateQueryOptions = queryOptions({
  queryKey: ["auth", "app-state"],
  queryFn: () => sessionService.getAppSession(),
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
      const res = await sessionService.signIn(data);

      await authQuery.refetch();

      return res;
    },
  });

  const signUp = useMutation({
    mutationKey: ["auth", "signup"],
    mutationFn: async (data: SignUpSchema) => {
      const res = await sessionService.registerRootOperator(data);

      await authQuery.refetch();

      return res;
    },
  });

  const signOut = useMutation({
    mutationKey: ["auth", "signout"],
    mutationFn: async () => {
      await sessionService.signOut();

      queryClient.invalidateQueries(appStateQueryOptions);
    },
  });

  const switchOrganization = useMutation({
    mutationKey: ["auth", "switch-organization"],
    mutationFn: async (organizationId: string) => {
      const newAppState = await sessionService.switchOrganization({
        organizationId,
      });

      await authQuery.refetch();

      return newAppState;
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
