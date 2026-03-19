import { useState } from "react";
import { Link, useRouter } from "@tanstack/react-router";

import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@shared/components/ui/card";
import {
  FieldError,
  FieldGroup,
  FieldLegend,
  FieldSet,
} from "@shared/components/ui/field";
import { Button } from "@shared/components/ui/button";

import { useAuthForm } from "@iam/presentation/components/auth-form";
import { signUpSchema } from "@iam/application/validators/operator-form.validators";
import { useSignUp } from "@iam/application/hooks/use-signup";

export function SignUpScreen() {
  const signUp = useSignUp();
  const router = useRouter();
  const [rootError, setRootError] = useState<{ message: string } | null>(null);

  const form = useAuthForm({
    defaultValues: {
      name: "cristian",
      email: "cristian@fludge.dev",
      pin: "123456",
      username: "cristian",
    },
    onSubmit: ({ value, formApi }) => {
      signUp.mutate(value, {
        onSuccess: () => {
          setRootError(null);
          formApi.reset();
          router.navigate({
            to: "/select-organization",
          });
        },
        onError: (error) => {
          if (typeof error === "string") {
            setRootError({
              message: error,
            });
            return;
          }

          setRootError({
            message: "Error al crear la cuenta",
          });
        },
      });
    },
    validators: {
      onChange: signUpSchema,
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
            className="space-y-8"
            onSubmit={(e) => {
              e.preventDefault();
              form.handleSubmit();
            }}
          >
            {rootError && <FieldError errors={[rootError]} />}
            <FieldSet>
              <FieldLegend>Informacion de usuario</FieldLegend>
              <FieldGroup className="flex flex-row">
                <form.AppField
                  name="name"
                  children={(field) => <field.NameField />}
                />
                <form.AppField
                  name="email"
                  children={(field) => <field.EmailField />}
                />
              </FieldGroup>
            </FieldSet>

            <FieldSet>
              <FieldLegend>Credenciales de acceso</FieldLegend>
              <FieldGroup className="flex flex-row">
                <form.AppField
                  name="username"
                  children={(field) => <field.UsernameField />}
                />
                <form.AppField
                  name="pin"
                  children={(field) => <field.PinField />}
                />
              </FieldGroup>
            </FieldSet>

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
              nativeButton={false}
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
