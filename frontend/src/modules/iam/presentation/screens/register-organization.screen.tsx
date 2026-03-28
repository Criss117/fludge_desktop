import { useState } from "react";
import { useRouter } from "@tanstack/react-router";
import { Plus } from "lucide-react";

import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@shared/components/ui/card";
import { Logo } from "@shared/components/logo";
import { FieldError, FieldGroup } from "@shared/components/ui/field";
import { Button } from "@shared/components/ui/button";
import { LinkButton } from "@shared/components/link-button";

import {
  createOrganizationSchema,
  type CreateOrganizationSchema,
} from "@iam/application/validators/organizations.validators";
import { useRegisterOrganizationForm } from "@iam/presentation/components/register-organization-form";
import { useMutateOrganizations } from "@iam/application/hooks/use-mutate-organizations";
import type { Organization } from "@iam/domain/entities/organization.entity";

interface Props {
  organizations: Organization[];
}

const defaultValues: CreateOrganizationSchema = {
  name: "Tienda Andres",
  legalName: "tienda andres",
  address: "una direccion",
  contactPhone: "1234567890",
  contactEmail: "cristian@gmail.com",
};

export function RegisterOrganizationScreen({ organizations }: Props) {
  const [rootError, setRootError] = useState<{ message: string } | null>(null);

  const { create } = useMutateOrganizations();
  const router = useRouter();

  const form = useRegisterOrganizationForm({
    validators: {
      onChange: createOrganizationSchema,
    },
    defaultValues,
    onSubmit: async ({ value }) => {
      create.mutate(value, {
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
        onSuccess: (activeOrganization) => {
          router.navigate({
            to: "/dashboard/$orgid",
            params: { orgid: activeOrganization?.slug },
          });
        },
      });
    },
  });

  return (
    <main className="mx-auto w-full max-w-2xl min-h-dvh flex items-center justify-center">
      <Card className="w-full">
        <CardHeader className="flex items-center flex-col">
          <div className="flex items-center gap-x-2">
            <Logo size={60} />
            <CardTitle className="text-4xl font-semibold text-primary">
              Fludge
            </CardTitle>
          </div>
          <CardDescription>
            Ingresa los detalles para comenzar a gestionar tu negocio
          </CardDescription>
        </CardHeader>
        <CardContent className="space-y-2">
          {rootError && <FieldError errors={[rootError]} />}

          <form
            noValidate
            className="space-y-5"
            onSubmit={(e) => {
              e.preventDefault();
              form.handleSubmit();
            }}
          >
            <form.AppField
              name="name"
              children={(field) => <field.OrganizationName />}
            />
            <form.AppField
              name="name"
              children={(field) => <field.OrganizationSlug />}
            />
            <FieldGroup className="flex md:flex-row items-start">
              <form.AppField
                name="address"
                children={(field) => <field.OrganizationAddress />}
              />
              <form.AppField
                name="legalName"
                children={(field) => <field.OrganizationLegalName />}
              />
            </FieldGroup>

            <FieldGroup className="flex md:flex-row items-start">
              <form.AppField
                name="contactEmail"
                children={(field) => <field.OrganizationContactEmail />}
              />
              <form.AppField
                name="contactPhone"
                children={(field) => <field.OrganizationContactPhone />}
              />
            </FieldGroup>

            <Button type="submit" className="w-full">
              Registrar
            </Button>
          </form>
        </CardContent>
        {organizations.length > 0 && (
          <CardFooter className="flex flex-col gap-y-5">
            <LinkButton
              variant="outline"
              className="gap-2 w-full"
              to="/select-organization"
            >
              <Plus className="size-4" />
              Seleccionar una organización
            </LinkButton>
          </CardFooter>
        )}
      </Card>
    </main>
  );
}
