import { useMutation } from "@tanstack/react-query";

import { SignUp } from "@wails/go/iam/IamHandler";

import type { SignUpSchema } from "@auth/application/validators/operator-form.validators";

export function useSignUp() {
  return useMutation({
    mutationKey: ["auth", "signup"],
    mutationFn: async (data: SignUpSchema) => {
      const operator = await SignUp(data);

      return operator;
    },
  });
}
