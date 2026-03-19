import { useMutation } from "@tanstack/react-query";
import type { CreateOrganizationSchema } from "@organizations/application/validators/organizations.validators";

export function useMutateOrganizations() {
  const create = useMutation({
    mutationKey: ["create-organization"],
    mutationFn: async (values: CreateOrganizationSchema) => {},
  });

  return { create };
}
