import { useMutation, useQueryClient } from "@tanstack/react-query";

import type { CreateOrganizationSchema } from "@iam/application/validators/organizations.validators";
import { appStateQueryOptions } from "@/integrations/iam";
import { organizationService } from "@iam/application/container";

export function useMutateOrganizations() {
  const queryClient = useQueryClient();

  const create = useMutation({
    mutationKey: ["create-organization"],
    mutationFn: async (values: CreateOrganizationSchema) => {
      const res = await organizationService.registerOrganization({
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
