import { useMemo } from "react";
import { createFormHook, createFormHookContexts } from "@tanstack/react-form";

import {
  Field,
  FieldContent,
  FieldDescription,
  FieldError,
  FieldLabel,
} from "@shared/components/ui/field";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectLabel,
  SelectTrigger,
  SelectValue,
} from "@shared/components/ui/select";
import { Input } from "@shared/components/ui/input";
import { Textarea } from "@shared/components/ui/textarea";
import { inputNumberHelper } from "@shared/lib/utils";
import type { Category } from "@catalog/domain/entities/category.entity";

interface CategoryFieldProps {
  categories: Category[];
}

export const { fieldContext, formContext, useFieldContext } =
  createFormHookContexts();

const { useAppForm } = createFormHook({
  fieldComponents: {
    NameField,
    SkuField,
    DescriptionField,
    CostPriceField,
    WholesalePriceField,
    SalePriceField,
    MinStockField,
    StockField,
    CategoryField,
  },
  formComponents: {},
  fieldContext,
  formContext,
});

function NameField() {
  const field = useFieldContext<string>();
  const id = field.name + "-product-name";

  const isInvalid = field.state.meta.isTouched && !field.state.meta.isValid;

  return (
    <Field className="gap-2" data-invalid={isInvalid}>
      <FieldLabel htmlFor={id}>
        Nombre Del Producto
        <span className="text-destructive">*</span>
      </FieldLabel>
      <Input
        id={id}
        type="text"
        placeholder="Ej: Monitor Gamer 27 pulgadas"
        name={field.name}
        value={field.state.value}
        onBlur={field.handleBlur}
        onChange={(e) => field.handleChange(e.target.value)}
        aria-invalid={isInvalid}
        required
      />
      {isInvalid && <FieldError errors={field.state.meta.errors} />}
    </Field>
  );
}

function SkuField() {
  const field = useFieldContext<string>();
  const id = field.name + "-product-sku";

  const isInvalid = field.state.meta.isTouched && !field.state.meta.isValid;

  return (
    <Field className="gap-2" data-invalid={isInvalid}>
      <FieldLabel htmlFor={id}>
        SKU / Código de barras
        <span className="text-destructive">*</span>
      </FieldLabel>
      <Input
        id={id}
        type="text"
        placeholder="Ej: AUTO-GEN-001"
        name={field.name}
        value={field.state.value}
        onBlur={field.handleBlur}
        onChange={(e) => field.handleChange(e.target.value)}
        aria-invalid={isInvalid}
        required
      />
      {isInvalid && <FieldError errors={field.state.meta.errors} />}
    </Field>
  );
}

function DescriptionField() {
  const field = useFieldContext<string>();
  const id = field.name + "-product-description";

  const isInvalid = field.state.meta.isTouched && !field.state.meta.isValid;

  return (
    <Field className="gap-2" data-invalid={isInvalid}>
      <FieldLabel htmlFor={id}>Descripcion del Producto</FieldLabel>
      <Textarea
        id={id}
        name={field.name}
        value={field.state.value}
        onBlur={field.handleBlur}
        onChange={(e) => field.handleChange(e.target.value)}
        aria-invalid={isInvalid}
        className="resize-none"
        placeholder="Descripción del Producto"
      />
      {isInvalid && <FieldError errors={field.state.meta.errors} />}
    </Field>
  );
}

function WholesalePriceField() {
  const field = useFieldContext<string>();
  const id = field.name + "-product-wholesale-price";

  const isInvalid = field.state.meta.isTouched && !field.state.meta.isValid;

  return (
    <Field className="gap-2" data-invalid={isInvalid}>
      <FieldLabel htmlFor={id}>
        Precio al por mayor
        <span className="text-destructive">*</span>
      </FieldLabel>
      <Input
        id={id}
        type="number"
        placeholder="Ej: 200"
        name={field.name}
        value={inputNumberHelper.value(
          field.state.value,
          field.state.meta.isTouched,
        )}
        onBlur={field.handleBlur}
        onChange={(e) =>
          field.handleChange(inputNumberHelper.onChange(e.target.value))
        }
        aria-invalid={isInvalid}
        required
      />
      {isInvalid && <FieldError errors={field.state.meta.errors} />}
    </Field>
  );
}

function SalePriceField() {
  const field = useFieldContext<string>();
  const id = field.name + "-product-sale-price";

  const isInvalid = field.state.meta.isTouched && !field.state.meta.isValid;

  return (
    <Field className="gap-2" data-invalid={isInvalid}>
      <FieldLabel htmlFor={id}>
        Precio de venta
        <span className="text-destructive">*</span>
      </FieldLabel>
      <Input
        id={id}
        type="number"
        placeholder="Ej: 200"
        name={field.name}
        value={inputNumberHelper.value(
          field.state.value,
          field.state.meta.isTouched,
        )}
        onBlur={field.handleBlur}
        onChange={(e) =>
          field.handleChange(inputNumberHelper.onChange(e.target.value))
        }
        aria-invalid={isInvalid}
        required
      />
      {isInvalid && <FieldError errors={field.state.meta.errors} />}
    </Field>
  );
}

function CostPriceField() {
  const field = useFieldContext<string>();
  const id = field.name + "-product-cost-price";

  const isInvalid = field.state.meta.isTouched && !field.state.meta.isValid;

  return (
    <Field className="gap-2" data-invalid={isInvalid}>
      <FieldLabel htmlFor={id}>
        Precio de compra
        <span className="text-destructive">*</span>
      </FieldLabel>
      <Input
        id={id}
        type="number"
        placeholder="Ej: 200"
        name={field.name}
        value={inputNumberHelper.value(
          field.state.value,
          field.state.meta.isTouched,
        )}
        onBlur={field.handleBlur}
        onChange={(e) =>
          field.handleChange(inputNumberHelper.onChange(e.target.value))
        }
        aria-invalid={isInvalid}
        required
      />
      {isInvalid && <FieldError errors={field.state.meta.errors} />}
    </Field>
  );
}

function StockField() {
  const field = useFieldContext<string>();
  const id = field.name + "-product-stock";

  const isInvalid = field.state.meta.isTouched && !field.state.meta.isValid;

  return (
    <Field className="gap-2" data-invalid={isInvalid}>
      <FieldLabel htmlFor={id}>
        Stock
        <span className="text-destructive">*</span>
      </FieldLabel>
      <Input
        id={id}
        type="number"
        placeholder="Ej: 10"
        name={field.name}
        value={inputNumberHelper.value(
          field.state.value,
          field.state.meta.isTouched,
        )}
        onBlur={field.handleBlur}
        onChange={(e) =>
          field.handleChange(inputNumberHelper.onChange(e.target.value))
        }
        aria-invalid={isInvalid}
        required
      />
      {isInvalid && <FieldError errors={field.state.meta.errors} />}
    </Field>
  );
}

function MinStockField() {
  const field = useFieldContext<string>();
  const id = field.name + "-product-minstock-level";

  const isInvalid = field.state.meta.isTouched && !field.state.meta.isValid;

  return (
    <Field className="gap-2" data-invalid={isInvalid}>
      <FieldLabel htmlFor={id}>Stock Minimo</FieldLabel>
      <Input
        id={id}
        type="number"
        placeholder="Ej: 10"
        name={field.name}
        value={inputNumberHelper.value(
          field.state.value,
          field.state.meta.isTouched,
        )}
        onBlur={field.handleBlur}
        onChange={(e) =>
          field.handleChange(inputNumberHelper.onChange(e.target.value))
        }
        aria-invalid={isInvalid}
        required
      />
      <FieldDescription>
        No llenar o colocar 0 para no tener stock minimo y permitir stock
        negativo
      </FieldDescription>
      {isInvalid && <FieldError errors={field.state.meta.errors} />}
    </Field>
  );
}

function CategoryField({ categories }: CategoryFieldProps) {
  const field = useFieldContext<string>();
  const id = field.name + "-product-category";

  const isInvalid = field.state.meta.isTouched && !field.state.meta.isValid;

  const categoryItems = useMemo(
    () =>
      categories.map((category) => ({
        value: category.id,
        label: category.name,
      })),
    [categories],
  );

  return (
    <Field data-invalid={isInvalid}>
      <FieldContent>
        <FieldLabel htmlFor={id}>Categoría</FieldLabel>
      </FieldContent>
      {isInvalid && <FieldError errors={field.state.meta.errors} />}
      <Select
        items={categoryItems}
        name={field.name}
        value={field.state.value}
        onValueChange={(v) => {
          if (v) field.handleChange(v);
        }}
      >
        <SelectTrigger id={id} aria-invalid={isInvalid}>
          <SelectValue placeholder="Selecciona una categoría" />
        </SelectTrigger>
        <SelectContent>
          <SelectGroup>
            <SelectLabel>Categorias</SelectLabel>
            {categoryItems.map((category) => (
              <SelectItem value={category.value} key={category.value}>
                {category.label}
              </SelectItem>
            ))}
          </SelectGroup>
        </SelectContent>
      </Select>
    </Field>
  );
}

export const useProductForm = useAppForm;
