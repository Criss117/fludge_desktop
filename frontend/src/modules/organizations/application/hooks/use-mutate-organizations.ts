import { useMutation, useQueryClient } from "@tanstack/react-query";

import type { CreateOrganizationSchema } from "@organizations/application/validators/organizations.validators";
import { RegisterOrganization } from "@wails/go/iam/IamHandler";
import { appStateQueryOptions } from "@/integrations/iam";

export function useMutateOrganizations() {
  const queryClient = useQueryClient();

  const create = useMutation({
    mutationKey: ["create-organization"],
    mutationFn: async (values: CreateOrganizationSchema) => {
      const res = await RegisterOrganization({
        address: values.address,
        contactEmail: values.contactEmail,
        contactPhone: values.contactPhone,
        legalName: values.legalName,
        name: values.name,
      });

      return res;
    },
    onSuccess: () => {
      queryClient.invalidateQueries(appStateQueryOptions);
    },
  });

  return { create };
}
