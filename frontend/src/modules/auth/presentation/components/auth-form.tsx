import { useId, useState } from "react";
import { createFormHook, createFormHookContexts } from "@tanstack/react-form";
import { Eye, EyeOff } from "lucide-react";

import {
  Field,
  FieldDescription,
  FieldError,
  FieldLabel,
} from "@shared/components/ui/field";
import { Input } from "@shared/components/ui/input";
import { Button } from "@shared/components/ui/button";

const { fieldContext, formContext, useFieldContext } = createFormHookContexts();

const { useAppForm } = createFormHook({
  fieldComponents: {
    NameField,
    EmailField,
    PinField,
    UsernameField,
  },
  formComponents: {},
  fieldContext,
  formContext,
});

function NameField() {
  const field = useFieldContext<string>();
  const id = `form-name-input-${useId()}`;

  const isInvalid = field.state.meta.isTouched && !field.state.meta.isValid;

  return (
    <Field data-invalid={isInvalid}>
      <FieldLabel htmlFor={id}>Nombre Completo</FieldLabel>
      <Input
        id={id}
        name={field.name}
        value={field.state.value}
        onBlur={field.handleBlur}
        onChange={(e) => field.handleChange(e.target.value)}
        aria-invalid={isInvalid}
        placeholder="shadcn"
        autoComplete="username"
      />
      {isInvalid && <FieldError errors={field.state.meta.errors} />}
    </Field>
  );
}

function EmailField() {
  const field = useFieldContext<string>();
  const id = `form-email-input-${useId()}`;

  const isInvalid = field.state.meta.isTouched && !field.state.meta.isValid;

  return (
    <Field data-invalid={isInvalid}>
      <FieldLabel htmlFor={id}>Correo Electronico</FieldLabel>
      <Input
        id={id}
        name={field.name}
        value={field.state.value}
        onBlur={field.handleBlur}
        onChange={(e) => field.handleChange(e.target.value)}
        aria-invalid={isInvalid}
        placeholder="ejemplo@fludge.dev"
        autoComplete="email"
      />
      {isInvalid && <FieldError errors={field.state.meta.errors} />}
    </Field>
  );
}

function UsernameField({ hideDescription = false }) {
  const field = useFieldContext<string>();
  const id = `form-username-input-${useId()}`;

  const isInvalid = field.state.meta.isTouched && !field.state.meta.isValid;

  return (
    <Field data-invalid={isInvalid}>
      <FieldLabel htmlFor={id}>Nombre de usuario</FieldLabel>
      <Input
        id={id}
        name={field.name}
        value={field.state.value}
        onBlur={field.handleBlur}
        onChange={(e) => field.handleChange(e.target.value)}
        aria-invalid={isInvalid}
        placeholder="jhondow"
        autoComplete="username"
      />
      {isInvalid ? (
        <FieldError errors={field.state.meta.errors} />
      ) : hideDescription ? null : (
        <FieldDescription>
          Ingresa tu nombre de usuario, con el podras acceder a tu cuenta
        </FieldDescription>
      )}
    </Field>
  );
}

function PinField({ hideDescription = false }) {
  const field = useFieldContext<string>();
  const id = `form-pin-input-${useId()}`;
  const [showPin, setShowPin] = useState(false);

  const isInvalid = field.state.meta.isTouched && !field.state.meta.isValid;

  return (
    <Field data-invalid={isInvalid}>
      <FieldLabel htmlFor={id}>Pin de acceso</FieldLabel>
      <div className="relative">
        <Input
          id={id}
          name={field.name}
          value={field.state.value}
          onBlur={field.handleBlur}
          onChange={(e) => field.handleChange(e.target.value)}
          aria-invalid={isInvalid}
          placeholder="123456"
          autoComplete="current-password"
          type={showPin ? "text" : "password"}
        />
        <Button
          variant="ghost"
          size="icon"
          className="absolute right-2"
          onClick={() => setShowPin((prev) => !prev)}
        >
          {showPin ? <EyeOff /> : <Eye />}
        </Button>
      </div>

      {isInvalid ? (
        <FieldError errors={field.state.meta.errors} />
      ) : hideDescription ? null : (
        <FieldDescription>
          Ingresa tu pin de acceso, debe tener 6 caracteres
        </FieldDescription>
      )}
    </Field>
  );
}

export const useAuthForm = useAppForm;
