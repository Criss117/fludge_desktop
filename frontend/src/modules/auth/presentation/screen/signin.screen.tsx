import { useState } from "react";
import { Link, useRouter } from "@tanstack/react-router";

import { Button } from "@shared/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@shared/components/ui/card";
import { FieldError, FieldGroup } from "@shared/components/ui/field";

import { useAuthForm } from "@auth/presentation/components/auth-form";
import { signInSchema } from "@auth/application/validators/operator-form.validators";

import { useAuth } from "@/integrations/auth";

export function SignInScreen() {
  const router = useRouter();
  const [rootError, setRootError] = useState<{ message: string } | null>(null);
  const { signIn } = useAuth();

  const form = useAuthForm({
    defaultValues: {
      pin: "123456",
      username: "cristian",
    },
    onSubmit: ({ value, formApi }) => {
      signIn.mutate(value, {
        onSuccess: (data) => {
          console.log(data);

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
      onChange: signInSchema,
    },
  });

  return (
    <main className="flex min-h-screen items-center justify-center">
      <Card className="w-full max-w-md ">
        <CardHeader className="">
          <CardTitle className="">Bienvenido</CardTitle>
          <CardDescription className="mt-1">
            Ingresa tus credenciales para continuar
          </CardDescription>
        </CardHeader>
        <CardContent>
          <form
            className="space-y-5"
            onSubmit={(e) => {
              e.preventDefault();
              form.handleSubmit();
            }}
          >
            {rootError && <FieldError errors={[rootError]} />}

            <FieldGroup>
              <form.AppField
                name="username"
                children={(field) => <field.UsernameField hideDescription />}
              />
              <form.AppField
                name="pin"
                children={(field) => <field.PinField hideDescription />}
              />
            </FieldGroup>

            <div className="flex justify-end">
              <Link className="text-sm text-primary hover:underline" to="/">
                ¿Olvidaste tu contraseña?
              </Link>
            </div>

            <Button className="font-semibold w-full" type="submit">
              Iniciar sesión
            </Button>
          </form>
        </CardContent>
        <CardFooter>
          <p className="text-sm">
            ¿No tienes una cuenta?{" "}
            <Link
              className="font-medium text-primary hover:underline"
              to="/auth/sign-up"
            >
              Crear cuenta
            </Link>
          </p>
        </CardFooter>
      </Card>
    </main>
  );
}
