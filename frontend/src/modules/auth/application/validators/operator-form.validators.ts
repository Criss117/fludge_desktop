import { z } from "zod";

export const signUpSchema = z.object({
  name: z
    .string("El nombre es requerido")
    .min(3, "El nombre es muy corto")
    .max(20, "El nombre es muy largo"),
  email: z.email("El correo electrónico es requerido"),
  pin: z
    .string("El pin es requerido")
    .min(6, "El pin debe tener 6 caracteres")
    .max(6, "El pin debe tener 6 caracteres"),
  username: z
    .string("El nombre de usuario es requerido")
    .min(3, "El nombre de usuario es muy corto")
    .max(10, "El nombre de usuario es muy largo"),
});

export const signInSchema = signUpSchema.pick({
  username: true,
  pin: true,
});

export type SignInSchema = z.infer<typeof signInSchema>;
export type SignUpSchema = z.infer<typeof signUpSchema>;
