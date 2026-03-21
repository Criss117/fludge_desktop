import {
  createColumnHelper,
  flexRender,
  getCoreRowModel,
  useReactTable,
} from "@tanstack/react-table";

import { Skeleton } from "@shared/components/ui/skeleton";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@shared/components/ui/table";

import type { ProductWithCategory } from "@catalog/application/hooks/use-products-queries";
import { WithChevronButtonSkeleton } from "./columns-header";

const columnHelper = createColumnHelper<ProductWithCategory>();

const columns = [
  columnHelper.accessor((p) => p.sku, {
    id: "sku",
    header: "SKU",
  }),
  columnHelper.accessor((p) => p.name, {
    id: "name",
    header: "Nombre",
  }),
  columnHelper.accessor((p) => p.description, {
    id: "description",
    header: "Descripción",
    cell: ({ getValue }) => {
      return <span>{getValue() || "-"}</span>;
    },
  }),
  columnHelper.accessor((p) => p.category, {
    id: "category",
    header: "Categoria",
    cell: ({ getValue }) => {
      return <span>{getValue()?.name || "-"}</span>;
    },
  }),
  columnHelper.accessor((p) => p.stock, {
    id: "stock",
    header: () => <WithChevronButtonSkeleton label="Stock" />,
  }),
  columnHelper.accessor((p) => p.costPrice, {
    id: "costPrice",
    header: () => <WithChevronButtonSkeleton label="Precio de Costo" />,
  }),
  columnHelper.accessor((p) => p.salePrice, {
    id: "salePrice",
    header: () => <WithChevronButtonSkeleton label="Precio de Venta" />,
  }),
  columnHelper.accessor((p) => p.wholesalePrice, {
    id: "wholesalePrice",
    header: () => <WithChevronButtonSkeleton label="Precio al por mayor" />,
  }),
  columnHelper.display({
    id: "actions",
  }),
];

export function ProductsTableSkeleton() {
  const table = useReactTable({
    data: [],
    columns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <div className="overflow-hidden rounded-md border">
      <Table>
        <TableHeader className="bg-muted sticky top-0">
          {table.getHeaderGroups().map((headerGroup) => (
            <TableRow key={headerGroup.id}>
              {headerGroup.headers.map((header) => {
                return (
                  <TableHead
                    key={header.id}
                    colSpan={header.colSpan}
                    className="text-center"
                  >
                    {header.isPlaceholder
                      ? null
                      : flexRender(
                          header.column.columnDef.header,
                          header.getContext(),
                        )}
                  </TableHead>
                );
              })}
            </TableRow>
          ))}
        </TableHeader>
        <TableBody className="**:data-[slot=table-cell]:first:w-8">
          {Array.from({ length: 10 }).map((_, index) => (
            <TableRow key={index}>
              <TableCell colSpan={columns.length} className="h-24 text-center">
                <Skeleton className="h-full w-full" />
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </div>
  );
}
