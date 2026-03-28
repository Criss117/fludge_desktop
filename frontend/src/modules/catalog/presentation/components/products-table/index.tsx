import { createContext, use } from "react";
import {
  flexRender,
  getCoreRowModel,
  useReactTable,
  type Table as TSTable,
} from "@tanstack/react-table";

import { productsTableColumns } from "./columns";

import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@shared/components/ui/table";
import type { Product } from "@catalog/domain/entities/product.entity";

interface Context {
  table: TSTable<Product>;
  columnsLength: number;
}

interface RootProps {
  products: Product[];
  children: React.ReactNode;
}

const ProductsContext = createContext<Context | null>(null);

function useProductsTable() {
  const context = use(ProductsContext);

  if (!context)
    throw new Error(
      "useProductsTableContext must be used within a ProductsTableProvider",
    );

  return context;
}

function Root({ products, children }: RootProps) {
  const table = useReactTable({
    columns: productsTableColumns,
    data: products,
    getRowId: (row) => row.id,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <ProductsContext.Provider
      value={{ table, columnsLength: productsTableColumns.length }}
    >
      {children}
    </ProductsContext.Provider>
  );
}

function Content({ emptyMessage = "No hay productos" }) {
  const { table, columnsLength } = useProductsTable();

  return (
    <div className="overflow-hidden rounded-xl border">
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
          {table.getRowModel().rows?.length ? (
            table.getRowModel().rows.map((row) => (
              <TableRow
                key={row.id}
                data-state={row.getIsSelected() && "selected"}
              >
                {row.getVisibleCells().map((cell) => (
                  <TableCell key={cell.id} className="text-center h-20">
                    {flexRender(cell.column.columnDef.cell, cell.getContext())}
                  </TableCell>
                ))}
              </TableRow>
            ))
          ) : (
            <TableRow>
              <TableCell colSpan={columnsLength} className="h-20 text-center">
                {emptyMessage}
              </TableCell>
            </TableRow>
          )}
        </TableBody>
      </Table>
    </div>
  );
}

export const ProductsTable = {
  useProductsTable,
  Root,
  Content,
};
