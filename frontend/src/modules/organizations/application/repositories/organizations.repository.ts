import { queryOptions } from "@tanstack/react-query";
import { FindManyOrganizationsByOperatorId } from "@wails/go/iam/IamHandler";

export const findManyOrganizations = {
  byOperator: (operatorId: string) =>
    queryOptions({
      queryKey: [operatorId, "all-organizations"],
      queryFn: async () => {
        const response = await FindManyOrganizationsByOperatorId(operatorId);

        return response;
      },
    }),
};
