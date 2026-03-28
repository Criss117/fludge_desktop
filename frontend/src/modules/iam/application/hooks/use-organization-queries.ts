import { queryOptions, useSuspenseQuery } from "@tanstack/react-query";
import { organizationService } from "@iam/application/container";

export const organizationQueryOptions = {
  findManyOrganizationsByRootOperator: queryOptions({
    queryKey: ["organizations"],
    queryFn: () => organizationService.findManyOrganizationsByRootOperator(),
  }),

  findOneOrganization: (orgId: string) =>
    queryOptions({
      queryKey: ["organizations", orgId],
      queryFn: () =>
        organizationService.findOneOrganization({
          organizationId: orgId,
        }),
    }),
};

export function useFindManyOrganizationsByRootOperator() {
  return useSuspenseQuery(
    organizationQueryOptions.findManyOrganizationsByRootOperator,
  );
}

export function useFindOneOrganization(orgId: string) {
  return useSuspenseQuery(organizationQueryOptions.findOneOrganization(orgId));
}
