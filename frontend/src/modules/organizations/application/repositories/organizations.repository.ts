import { queryOptions } from "@tanstack/react-query";
import { FindManyOrganizationsByRootOperator } from "@wails/go/iam/IamHandler";

export const findManyOrganizations = {
  byOperator: (operatorId: string) =>
    queryOptions({
      queryKey: [operatorId, "all-organizations"],
      queryFn: async () => {
        const response = await FindManyOrganizationsByRootOperator();

        return response;
      },
    }),
};
