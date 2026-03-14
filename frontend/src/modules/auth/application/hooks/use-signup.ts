import { useMutation } from "@tanstack/react-query";

import { CreateOperator } from "@wails/go/auth/AuthHandler";

import type { SignUpSchema } from "@auth/application/validators/operator-form.validators";

export function useSignUp() {
  return useMutation({
    mutationKey: ["auth", "signup"],
    mutationFn: async (data: SignUpSchema) => {
      const operator = await CreateOperator(data);

      return operator;
    },
  });
}
