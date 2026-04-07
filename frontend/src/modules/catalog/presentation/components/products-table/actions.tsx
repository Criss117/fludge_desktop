import { MoreVerticalIcon, PencilIcon, TrashIcon } from "lucide-react";

import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@shared/components/ui/dropdown-menu";
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from "@shared/components/ui/alert-dialog";

import { Button } from "@shared/components/ui/button";

import { useUpdateProduct } from "@catalog/presentation/components/update-product";
import { useMutateProducts } from "@catalog/application/hooks/use-mutate-products";
import type { Product } from "@catalog/domain/entities/product.entity";

interface Props {
  product: Product;
}

function DeleteProduct({ product }: Props) {
  const { remove } = useMutateProducts();

  const handleDelete = () => {
    remove.mutate(product.id);
  };

  return (
    <AlertDialog>
      <AlertDialogTrigger
        render={(props) => (
          <DropdownMenuItem
            variant="destructive"
            className="gap-2"
            closeOnClick={false}
            {...props}
          />
        )}
      >
        <TrashIcon className="h-4 w-4" />
        Eliminar Producto
      </AlertDialogTrigger>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>Eliminar {product.name}</AlertDialogTitle>
          <AlertDialogDescription>
            Esta acción no se puede deshacer.
          </AlertDialogDescription>
        </AlertDialogHeader>
        <AlertDialogFooter>
          <AlertDialogAction
            variant="destructive"
            onClick={handleDelete}
            disabled={remove.isPending}
          >
            Eliminar
          </AlertDialogAction>
          <AlertDialogCancel
            render={(props) => <Button variant="outline" {...props} />}
          >
            Cancelar
          </AlertDialogCancel>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  );
}

export function ProductsTableActions({ product }: Props) {
  const { selectProduct } = useUpdateProduct();

  return (
    <DropdownMenu>
      <DropdownMenuTrigger
        render={(props) => <Button {...props} variant="ghost" size="icon" />}
      >
        <MoreVerticalIcon />
        <span className="sr-only">Acciones para {product.name}</span>
      </DropdownMenuTrigger>
      <DropdownMenuContent align="end" className="w-48">
        <DropdownMenuGroup>
          <DropdownMenuItem
            className="gap-2"
            onClick={() => selectProduct(product)}
          >
            <PencilIcon className="h-4 w-4" />
            Editar Producto
          </DropdownMenuItem>
        </DropdownMenuGroup>
        <DropdownMenuSeparator />
        <DropdownMenuGroup>
          <DropdownMenuItem
            variant="destructive"
            className="gap-2"
            render={() => <DeleteProduct product={product} />}
          />
        </DropdownMenuGroup>
      </DropdownMenuContent>
    </DropdownMenu>
  );
}
