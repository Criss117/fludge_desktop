import { useId, useState } from "react";
import { PlusIcon } from "lucide-react";
import { toast } from "sonner";

import { FieldGroup } from "@shared/components/ui/field";
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@shared/components/ui/dialog";
import { Button } from "@shared/components/ui/button";

import { useCategoryForm } from "@catalog/presentation/components/category-form";
import { useMutateCategories } from "@catalog/application/hooks/use-mutate-categories";
import {
  createCategorySchema,
  type CreateCategorySchema,
} from "@catalog/application/validators/category.validators";

const defaultValues: CreateCategorySchema = {
  name: "",
  description: "",
};

export function CreateCategory() {
  const [isOpen, setIsOpen] = useState(false);
  const formId = `create-category-form-${useId()}`;
  const { create } = useMutateCategories();

  const form = useCategoryForm({
    defaultValues,
    validators: {
      onChange: createCategorySchema,
    },
    onSubmit: ({ value, formApi }) => {
      const toastLoadingId = toast.loading("Creando categoria...", {
        position: "top-center",
      });

      create.mutate(value, {
        onSettled: () => toast.dismiss(toastLoadingId),
        onError: (error) =>
          toast.error(error.message, {
            position: "top-center",
          }),
        onSuccess: () => {
          toast.success("Categoria creada exitosamente!", {
            position: "top-center",
          });
          formApi.reset();
          setIsOpen(false);
        },
      });
    },
  });

  return (
    <Dialog
      open={isOpen}
      onOpenChange={(v) => {
        if (!v) form.reset();
        setIsOpen(v);
      }}
    >
      <DialogTrigger render={(props) => <Button {...props} />}>
        <PlusIcon />
        Nueva Categoria
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Nueva Categoria</DialogTitle>
          <DialogDescription>
            Llena los campos para crear una nueva categoria
          </DialogDescription>
        </DialogHeader>
        <form
          noValidate
          id={formId}
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
              name="description"
              children={(field) => <field.DescriptionField />}
            />
          </FieldGroup>
        </form>
        <DialogFooter>
          <Button type="submit" form={formId} disabled={create.isPending}>
            Crear Categoria
          </Button>
          <DialogClose
            render={(props) => <Button {...props} variant="outline" />}
          >
            Cancelar
          </DialogClose>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}
