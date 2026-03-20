import { z } from "zod";
import { Link } from "@tanstack/react-router";

import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@shared/components/ui/card";
import { FieldGroup } from "@shared/components/ui/field";
import { Button } from "@shared/components/ui/button";

import { useAuthForm } from "@iam/presentation/components/auth-form";

const createOperatorSchema = z.object({
  name: z
    .string("El nombre es requerido")
    .min(3, "El nombre es muy corto")
    .max(20, "El nombre es muy largo"),
  email: z.string().email("El email es inválido"),
});

export function SignUpScreen() {
  const form = useAuthForm({
    defaultValues: {
      name: "",
      email: "",
    },
    onSubmit: () => {},
    validators: {
      onSubmit: createOperatorSchema,
    },
  });

  return (
    <main className="flex min-h-screen items-center justify-center p-4">
      <Card className="w-full max-w-xl">
        <CardHeader className="flex flex-col items-center">
          <CardTitle className="text-3xl font-bold tracking-tight">
            Crear cuenta
          </CardTitle>
          <CardDescription className="mt-1">
            Ingresa tus datos para registrarte
          </CardDescription>
        </CardHeader>
        <CardContent>
          <form
            className="flex flex-col gap-5"
            onSubmit={(e) => {
              e.preventDefault();
              form.handleSubmit();
            }}
          >
            <FieldGroup>
              <form.AppField
                name="name"
                children={(field) => <field.NameField />}
              />

              <form.AppField
                name="email"
                children={(field) => <field.NameField />}
              />
            </FieldGroup>

            {/* <Fieldset className="flex flex-row flex-1 ">
              <TextField
                isRequired
                name="password"
                type="password"
                className="flex-1"
                validate={(value) => {
                  if (value.length < 6) {
                    return "La contraseña debe tener al menos 6 caracteres";
                  }
                  return null;
                }}
              >
                <Label>Contraseña</Label>
                <Input placeholder="Crea una contraseña" />
                <FieldError />
              </TextField>

              <TextField
                isRequired
                name="confirmPassword"
                type="password"
                className="flex-1"
              >
                <Label>Confirmar contraseña</Label>
                <Input placeholder="Confirma tu contraseña" />
                <FieldError />
              </TextField>
            </Fieldset> */}

            <Button className="w-full text-base font-semibold" type="submit">
              Crear cuenta
            </Button>
          </form>
        </CardContent>
        <CardFooter>
          <p className="text-sm ">
            ¿Ya tienes una cuenta?{" "}
            <Button
              variant="link"
              render={(props) => (
                <Link {...props} to="/">
                  Iniciar sesión
                </Link>
              )}
            />
          </p>
        </CardFooter>
      </Card>
    </main>
  );
}
