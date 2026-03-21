import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@shared/components/ui/card";
import { CreateProduct } from "@catalog/presentation/components/create-product";
import { Button } from "@/modules/shared/components/ui/button";
import { PlusIcon } from "lucide-react";

interface Props {
  totalProducts: number;
}

export function ProductsHeaderSection({ totalProducts }: Props) {
  return (
    <Card className="flex justify-between flex-row">
      <CardHeader className="flex-1">
        <CardTitle className="text-2xl">Gestion de Productos</CardTitle>
        <CardDescription>
          Administra el stock, precio y categorias de tus productos
        </CardDescription>
        <CardDescription>
          ({totalProducts}) productos registrados
        </CardDescription>
      </CardHeader>
      <CardContent className="flex flex-col gap-y-2">
        <CreateProduct />
        {/* <RegisterProduct />
        <CategoriesList /> */}
      </CardContent>
    </Card>
  );
}

export function ProductsHeaderSectionSkeleton() {
  return (
    <Card className="flex justify-between flex-row">
      <CardHeader className="flex-1">
        <CardTitle className="text-2xl">Gestion de Productos</CardTitle>
        <CardDescription>
          Administra el stock, precio y categorias de tus productos
        </CardDescription>
        <CardDescription>({0}) productos registrados</CardDescription>
      </CardHeader>
      <CardContent className="flex flex-col gap-y-2">
        <Button disabled>
          <PlusIcon />
          Registrar un producto
        </Button>
      </CardContent>
    </Card>
  );
}
